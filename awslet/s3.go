package awslet

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hellgate75/aws-architect/helpers"
	"github.com/hellgate75/aws-architect/log"
	"os"
	"time"
	"strings"
	"errors"
	"net/http"
	"bytes"
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
				Status:    aws.String("Enabled"),
				MFADelete: aws.String("Disabled"),
			},
		})
	}
	if corsFile != "" {
		if _, err = os.Stat(corsFile); err == nil {
			cors, err := helpers.LoadCORs(corsFile)
			if err == nil {
				var corsList []*s3.CORSRule = make([]*s3.CORSRule, 0)
				for i := 0; i < len(cors.CORSConfiguration.CORSRules); i++ {
					ruleConf := cors.CORSConfiguration.CORSRules[i]
					if ruleConf.MaxAgeSeconds > 0 {
						corsList = append(corsList, &s3.CORSRule{
							MaxAgeSeconds:  aws.Int64(int64(ruleConf.MaxAgeSeconds)),
							AllowedOrigins: aws.StringSlice(ruleConf.AllowedOrigins),
							AllowedHeaders: aws.StringSlice(ruleConf.AllowedHeaders),
							AllowedMethods: aws.StringSlice(ruleConf.AllowedMethods),
							ExposeHeaders:  aws.StringSlice(ruleConf.ExposeHeaders),
						})
					} else {
						corsList = append(corsList, &s3.CORSRule{
							AllowedOrigins: aws.StringSlice(ruleConf.AllowedOrigins),
							AllowedHeaders: aws.StringSlice(ruleConf.AllowedHeaders),
							AllowedMethods: aws.StringSlice(ruleConf.AllowedMethods),
							ExposeHeaders:  aws.StringSlice(ruleConf.ExposeHeaders),
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
	return err == nil, err
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
			Quiet:   aws.Bool(Quite),
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
	return err == nil, err
}

func BucketStatus(service *s3.S3, bucketName string) (bool, error) {
	output, err := service.GetBucketLocation(&s3.GetBucketLocationInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return false, err
	}
	logger.Log(fmt.Sprintf("Bucket Name: %s\nBucket Location Constraint: %s", bucketName, output.String()))
	return true, err
}

func ListBuckets(service *s3.S3) ([]*s3.Bucket, error) {
	var output *s3.ListBucketsOutput
	var err error
	output, err = service.ListBuckets(&s3.ListBucketsInput{
	})
	if err != nil {
		return make([]*s3.Bucket, 0), err
	}
	return output.Buckets, err
}

func BucketUpload(service *s3.S3, bucketName string, acl string, fileName string, keyName string, storageClass string, multipart bool) (bool, error) {
	if storageClass == "" || ! strings.Contains(STORAGE_TYPE_STRING, strings.ToUpper(storageClass+" ")) {
		return false, errors.New("Storage type '"+storageClass+"' is not valid.")
	}
	if acl == "" || ! strings.Contains(ACL_TYPE_STRING, strings.ToLower(acl+" ")) {
		return false, errors.New("Acl '"+acl+"' is not valid.")
	}
	_, err := os.Stat(fileName)
	if err != nil {
		return false, err
	}

	file, err := os.Open(fileName)
	if err != nil {
		return false, err
	}

	defer file.Close()

	stat, _ := file.Stat()
	size := stat.Size()
	buffer := make([]byte, size)

	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	output, err := service.PutObject(&s3.PutObjectInput {
		Bucket: aws.String(bucketName),
		ACL: aws.String(strings.ToLower(acl)),
		Key: aws.String(keyName),
		Body: fileBytes,
		StorageClass: aws.String(strings.ToUpper(storageClass)),
		ContentLength: aws.Int64(size),
		ContentType: aws.String(fileType),
	})
	if err != nil {
		return false, err
	}
	logger.Log(fmt.Sprintf("Bucket Name: %s\nUploaded File Location : %s", bucketName, output.String()))
	return true, err
}
