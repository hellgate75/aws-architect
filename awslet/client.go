package awslet

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

const DEFAULT_AWS_REGION = "us-east-1"

func CreateSession() *session.Session {
	return session.Must(session.NewSession())
}

func createRegionConfigWithCredentials(region string, aws_access_key_id string, aws_secret_access_key string, aws_access_token string) (*aws.Config, error) {
	credentialVault := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, aws_access_token)

	_, err := credentialVault.Get()
	if err != nil {
		return  nil, err
	}

	return aws.NewConfig().WithRegion(region).WithCredentials(credentialVault), nil

}

func CreateS3ServiceWithCredentials(session *session.Session, region string, aws_access_key_id string, aws_secret_access_key string, aws_access_token string) (*s3.S3, error) {
	cfg, err := createRegionConfigWithCredentials(region, aws_access_key_id, aws_secret_access_key, aws_access_token)
	if err != nil {
		return  nil, err
	}
	return s3.New(session, cfg), err
}

func CreateS3Service(session *session.Session, region string) *s3.S3 {
	var endpoint_region string = "-" + region
	if region == "us-east-1" {
		endpoint_region = ""
	}
	var endpoint string = "s3" + endpoint_region + ".amazonaws.com"
	return s3.New(session, &aws.Config{Region: aws.String(region), Endpoint: aws.String(endpoint)})
}

func CreateS3ServiceAssumeRole(session *session.Session, region string, role string) *s3.S3 {
	var endpoint_region string = "-" + region
	if region == "us-east-1" {
		endpoint_region = ""
	}
	var endpoint string = "s3" + endpoint_region + ".amazonaws.com"
	return s3.New(session, &aws.Config{Region: aws.String(region), Endpoint: aws.String(endpoint), Credentials: stscreds.NewCredentials(session, role)})
}

func CreateCfnServiceWithCredentials(session *session.Session, region string, aws_access_key_id string, aws_secret_access_key string, aws_access_token string) (*cloudformation.CloudFormation, error) {
	cfg, err := createRegionConfigWithCredentials(region, aws_access_key_id, aws_secret_access_key, aws_access_token)
	if err != nil {
		return  nil, err
	}
	return cloudformation.New(session, cfg), err
}


func CreateCfnService(session *session.Session, region string) *cloudformation.CloudFormation {
	return cloudformation.New(session, &aws.Config{Region: aws.String(region), Endpoint: aws.String("cloudformation." + region + ".amazonaws.com")})
}

func CreateCfnServiceAssumeRole(session *session.Session, region string, role string) *cloudformation.CloudFormation {
	return cloudformation.New(session, &aws.Config{Region: aws.String(region), Endpoint: aws.String("cloudformation." + region + ".amazonaws.com"), Credentials: stscreds.NewCredentials(session, role)})
}
