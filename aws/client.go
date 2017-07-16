package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

const DEFAULT_AWS_REGION="us-east-1"

func CreateSession() (*session.Session) {
	return  session.Must(session.NewSession())
}

func CreateS3Service(session *session.Session, region string) (*s3.S3) {
	var endpoint_region string  = "-"+region
	if region == "us-east-1" {
		endpoint_region = ""
	}
	var endpoint string  = "s3" + endpoint_region + ".amazonaws.com"
	return  s3.New(session, &aws.Config{Region: &region, Endpoint: &endpoint})
}

func CreateS3ServiceAssumeRole(session *session.Session, region string, role string) (*s3.S3) {
	creds := stscreds.NewCredentials(session, role)
	return  s3.New(session, &aws.Config{Region: &region, Credentials: creds})
}
