package s3_bucket

import (
	"github.com/hellgate75/aws-architect/awslet"
	"fmt"
	"flag"
	"github.com/hellgate75/aws-architect/command"
	"github.com/hellgate75/aws-architect/abstract"
	"strconv"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3DeleteBucket struct {
}

func (p *S3DeleteBucket) Execute(action *abstract.ActionImpl, arguments []interface{}, logChannel chan string) bool {
	var bucketName string = fmt.Sprintf("%v", arguments[0])
	var awsRegion string = fmt.Sprintf("%v", arguments[1])
	var iamRole string = fmt.Sprintf("%v", arguments[5])
	s3Recursive, _ := strconv.ParseBool(fmt.Sprintf("%v", arguments[2]))
	defer func() {
		// recover from panic caused by writing to a closed channel
		if r := recover(); r != nil {
		}
	}()
	action.InProgress = true
	session := awslet.CreateSession()
	var awsService *s3.S3
	if iamRole == "" {
		awsService = awslet.CreateS3Service(session, awsRegion)
	} else {
		awsService = awslet.CreateS3ServiceAssumeRole(session, awsRegion, iamRole)

	}
	logChannel <- fmt.Sprintf("Deleting bucket : %s", bucketName)
	logChannel <- fmt.Sprintf("Bucket Region : %s", awsRegion)
	var removed bool
	var err error
	if s3Recursive {
		removed, err = awslet.DeleteBucketRecursive(awsService, bucketName)
	} else {
		removed, err = awslet.DeleteBucket(awsService, bucketName)
	}
	if err == nil {
		logChannel <- fmt.Sprintf("Bucket '%s' removed : %t", bucketName, removed)
		action.Success = true
		action.InProgress = false
		action.Message = fmt.Sprintf("Deletion of S3 Bucket %s in Region %s successful!!", bucketName, awsRegion)
		return true
	}
	logChannel <- err.Error()
	action.Message = fmt.Sprintf("Error in deleting S3 Bucket %s : %s", bucketName, err.Error())
	action.Success = false
	action.InProgress = false
	return false
}

type S3DeleteBucketParser struct {
	BucketName string
	Region     string
	Recursive  bool
	UseRole    string
}

func (p *S3DeleteBucketParser) Validate() bool {
	flag.StringVar(&p.BucketName, "bucket", "", "Amazon Web Services S3 Bucket Name")
	flag.StringVar(&p.Region, "region", command.DEFAULT_AWS_REGION, "Amazon Web Service reference Region (default : "+command.DEFAULT_AWS_REGION+")")
	flag.BoolVar(&p.Recursive, "recursive", false, "Delete recursively keys from S3 Bucket")
	flag.StringVar(&p.UseRole, "use-role", "", "Amazon Web Services IAM Role for action (default : \"\")")
	flag.Parse()
	return p.BucketName != ""
}

func (p *S3DeleteBucketParser) Parse() []interface{} {
	var arguments []interface{} = make([]interface{}, 0)
	arguments = append(arguments, p.BucketName)
	arguments = append(arguments, p.Region)
	arguments = append(arguments, p.Recursive)
	arguments = append(arguments, p.UseRole)
	return arguments
}

