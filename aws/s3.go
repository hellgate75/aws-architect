package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"time"
	"os"
	"github.com/hellgate75/aws-architect/helpers"
	"github.com/hellgate75/aws-architect/log"
	"fmt"
)

var logger log.Logger = log.Logger{}

func CreateBucket(service *s3.S3, bucketName string, region string, acl string, versioning bool, corsFile string) (*string, error) {
	var output *s3.CreateBucketOutput
	var err error
	if region != DEFAULT_AWS_REGION {
		output, err = service.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
			ACL:    &acl,
			CreateBucketConfiguration: &s3.CreateBucketConfiguration{
				LocationConstraint: aws.String(region),
			},
		})
	} else {
		output, err = service.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
			ACL:    &acl,
		})
	}
	if err != nil {
		return nil, err
	}
	err = service.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if versioning {
		time.Sleep(time.Millisecond * 500)
		_, err = service.PutBucketVersioning(&s3.PutBucketVersioningInput{
			Bucket: aws.String(bucketName),
			VersioningConfiguration: &s3.VersioningConfiguration{
				Status: aws.String("Enabled"),
				MFADelete: aws.String("Disabled"),
			},
		})
	}
	if corsFile != "" {
		if _,err = os.Stat(corsFile); err==nil {
			cors, err := helpers.LoadCORs(corsFile)
			if err == nil {
				var corsList []*s3.CORSRule = make([]*s3.CORSRule, 0)
				for i := 0; i < len(cors.CORSConfiguration.CORSRules); i++ {
					ruleConf := cors.CORSConfiguration.CORSRules[i]
					if ruleConf.MaxAgeSeconds > 0 {
						corsList = append(corsList, &s3.CORSRule{
							MaxAgeSeconds: aws.Int64(int64(ruleConf.MaxAgeSeconds)),
							AllowedOrigins: aws.StringSlice(ruleConf.AllowedOrigins),
							AllowedHeaders: aws.StringSlice(ruleConf.AllowedHeaders),
							AllowedMethods: aws.StringSlice(ruleConf.AllowedMethods),
							ExposeHeaders: aws.StringSlice(ruleConf.ExposeHeaders),
						})
					} else {
						corsList = append(corsList, &s3.CORSRule{
							AllowedOrigins: aws.StringSlice(ruleConf.AllowedOrigins),
							AllowedHeaders: aws.StringSlice(ruleConf.AllowedHeaders),
							AllowedMethods: aws.StringSlice(ruleConf.AllowedMethods),
							ExposeHeaders: aws.StringSlice(ruleConf.ExposeHeaders),
						})
					}
				}
				println(fmt.Sprintf("%v", corsList))
				time.Sleep(time.Millisecond * 500)
				_, err = service.PutBucketCors(&s3.PutBucketCorsInput{
					Bucket: aws.String(bucketName),
					CORSConfiguration: &s3.CORSConfiguration{
						CORSRules: corsList,
					},
				})
			} else {
				logger.WarningE(err)
			}
		} else {
			logger.WarningE(err)
		}
	}
	return output.Location, err
}

func DeleteBucket(service *s3.S3, bucketName string) (bool, error) {
	_, err := service.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return false, err
	}
	err = service.WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	return true, err
}

func DeleteBucketRecursive(service *s3.S3, bucketName string) (bool, error) {
	objectOutput, objerr := service.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
	})
	if objerr != nil {
		return false, objerr
	}
	var objects []*s3.Object = objectOutput.Contents
	var removeList []*s3.ObjectIdentifier = make([]*s3.ObjectIdentifier, 0)
	for i := 0; i < len(objects); i++ {
		obj := objects[i]
		removeList = append(removeList, &s3.ObjectIdentifier{
			Key: obj.Key,
		})
	}
	var Quite bool = true
	service.DeleteObjects(&s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &s3.Delete{
			Quiet:   &Quite,
			Objects: removeList,
		},
	})
	_, err := service.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return false, err
	}
	err = service.WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	return true, err
}
