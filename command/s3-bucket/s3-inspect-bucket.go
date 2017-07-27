package s3_bucket

import (
	"github.com/hellgate75/aws-architect/awslet"
	"fmt"
	"flag"
	"github.com/hellgate75/aws-architect/command"
	"github.com/hellgate75/aws-architect/abstract"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3BucketStatus struct {
}

func (p *S3BucketStatus) Execute(action *abstract.ActionImpl, arguments []interface{}, logChannel chan string) bool {
	var bucketName string = fmt.Sprintf("%v", arguments[0])
	var awsRegion string = fmt.Sprintf("%v", arguments[1])
	defer func() {
		// recover from panic caused by writing to a closed channel
		if r := recover(); r != nil {
		}
	}()
	action.InProgress = true
	session := awslet.CreateSession()
	var awsService *s3.S3 = awslet.CreateS3Service(session, awsRegion)
	logChannel <- fmt.Sprintf("Status for bucket : %s", bucketName)
	logChannel <- fmt.Sprintf("Bucket Region : %s", awsRegion)
	var exists bool
	var err error
	exists, err = awslet.BucketStatus(awsService, bucketName)
	if err == nil {
		logChannel <- fmt.Sprintf("Bucket '%s' exists : %t", bucketName, exists)
		action.Message = fmt.Sprintf("Status of S3 Bucket %s in Region %s : %t!!", bucketName, awsRegion, exists)
		action.Success = exists
		action.InProgress = false
		return true
	}
	logChannel <- err.Error()
	action.Message = fmt.Sprintf("Error in recoverying S3 Bucket %s : %s", bucketName, err.Error())
	action.Success = false
	action.InProgress = false
	return false
}

type S3BucketStatusParser struct {
	BucketName string
	Region     string
}

func (p *S3BucketStatusParser) Validate() bool {
	flag.StringVar(&p.BucketName, "bucket", "", "Amazon Web Services S3 Bucket Name")
	flag.StringVar(&p.Region, "region", command.DEFAULT_AWS_REGION, "Amazon Web Service reference Region (default : "+command.DEFAULT_AWS_REGION+")")
	flag.Parse()
	return p.BucketName != ""
}

func (p *S3BucketStatusParser) Parse() []interface{} {
	var arguments []interface{} = make([]interface{}, 0)
	arguments = append(arguments, p.BucketName)
	arguments = append(arguments, p.Region)
	return arguments
}
