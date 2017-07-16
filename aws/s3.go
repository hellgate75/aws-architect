package aws

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

func CreateBucket(service *s3.S3, bucketName string, acl string) (*string, error) {
	output, err := service.CreateBucket(s3.CreateBucketInput{
		Bucket: bucketName,
		ACL: acl,
	})
	if err != nil {
		return  nil, err
	}
	err = service.WaitUntilBucketExists(s3.HeadBucketInput{
		Bucket: bucketName,
	})
	return &output.Location, err
}

func DeleteBucket(service *s3.S3, bucketName string) (bool, error) {
	_, err := service.DeleteBucket(s3.DeleteBucketInput{
		Bucket: bucketName,
	})
	if err != nil {
		return  false, err
	}
	err = service.WaitUntilBucketNotExists(s3.HeadBucketInput{
		Bucket: bucketName,
	})
	return true, err
}

func DeleteBucketrecursive(service *s3.S3, bucketName string) (bool, error) {
	objectOutput, objerr := service.ListObjects(s3.ListObjectsInput{
		Bucket: bucketName,
	})
	if objerr != nil {
		return false, objerr
	}
	var objects []*s3.Object = objectOutput.Contents
	var removeList []*s3.ObjectIdentifier = make([]*s3.ObjectIdentifier, 0)
	for i:=0; i<len(objects); i++ {
		obj := objects[i]
		removeList=append(removeList,s3.ObjectIdentifier{
			Key: obj.Key,
		})
	}
	service.DeleteObjects(s3.DeleteObjectsInput{
		Bucket: bucketName,
		Delete: s3.Delete{
			Quiet: true,
			Objects: removeList,
		},
	})
	_, err := service.DeleteBucket(s3.DeleteBucketInput{
		Bucket: bucketName,
	})
	if err != nil {
		return  false, err
	}
	err = service.WaitUntilBucketNotExists(s3.HeadBucketInput{
		Bucket: bucketName,
	})
	return true, err
}


