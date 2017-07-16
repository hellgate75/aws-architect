package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/endpoints"
)

func CreateSession() (*session.Session) {
	return  session.Must(session.NewSession())
}

func

func CreateS3Client(session *session.Session, region endpoints.Region, role string) (*s3.S3) {
	return  s3.New(session, &aws.Config{Region: region.ID()})
}

func CreateS3ClientAssumeRole(session *session.Session, region endpoints.Region, role string) (*s3.S3) {
	creds := stscreds.NewCredentials(session, role)
	return  s3.New(session, &aws.Config{Region: region.ID(), Credentials: creds})
}
