package s3_bucket

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hellgate75/aws-architect/abstract"
	"github.com/hellgate75/aws-architect/awslet"
	"strings"
	"github.com/hellgate75/aws-architect/command"
)

type S3CreateBucket struct {
}

func (p *S3CreateBucket) Execute(action *abstract.ActionImpl, arguments []interface{}, logChannel chan string) bool {
	var bucketName string = fmt.Sprintf("%v", arguments[0])
	var awsRegion string = fmt.Sprintf("%v", arguments[1])
	var awsAcl string = fmt.Sprintf("%v", arguments[2])
	var s3Versioning string = fmt.Sprintf("%v", arguments[3])
	var s3Cors string = fmt.Sprintf("%v", arguments[4])
	var iamRole string = fmt.Sprintf("%v", arguments[5])
	//var s3WebSite string = fmt.Sprintf("%v", arguments[5])
	var accessId string = fmt.Sprintf("%v", arguments[6])
	var accessKey string = fmt.Sprintf("%v", arguments[7])
	var accessToken string = fmt.Sprintf("%v", arguments[8])

	defer func() {
		// recover from panic caused by writing to a closed channel
		if r := recover(); r != nil {
		}
	}()
	action.InProgress = true
	session := awslet.CreateSession()
	var awsService *s3.S3
	var credErr error
	if accessId != "" && accessKey != "" {
		awsService, credErr = awslet.CreateS3ServiceWithCredentials(session, awslet.DEFAULT_AWS_REGION, accessId, accessKey, accessToken)
		if credErr != nil {
			logChannel <- credErr.Error()
			action.Message = fmt.Sprintf("Credential Error for ID :  %s connectiong to S3 Service : %s", accessId, credErr.Error())
			action.Success = false
			action.InProgress = false
			return false
		}
	} else if iamRole == "" {
		awsService = awslet.CreateS3Service(session, awsRegion)
	} else {
		awsService = awslet.CreateS3ServiceAssumeRole(session, awsRegion, iamRole)

	}
	logChannel <- fmt.Sprintf("Creating bucket : %s", bucketName)
	logChannel <- fmt.Sprintf("Bucket Region : %s", awsRegion)
	logChannel <- fmt.Sprintf("Bucket ACL : %s", awsAcl)
	logChannel <- fmt.Sprintf("Bucket Versioning : %s", s3Versioning)
	logChannel <- fmt.Sprintf("Bucket Cors file : %s", s3Cors)
	bucketLocation, err := awslet.CreateBucket(awsService, bucketName, awsRegion, awsAcl, strings.ToLower(s3Versioning) == "enabled", s3Cors)
	if err == nil {
		logChannel <- fmt.Sprintf("Bucket '%s' created at location : %s", bucketName, *bucketLocation)
		action.Success = true
		action.InProgress = false
		action.Message = fmt.Sprintf("Creation S3 Bucket %s (location: %s) in Region %s successful!!", bucketName, *bucketLocation, awsRegion)
		return true
	}
	logChannel <- err.Error()
	action.Message = fmt.Sprintf("Error in creating S3 Bucket %s : %s", bucketName, err.Error())
	action.Success = false
	action.InProgress = false
	return false
}

type S3CreateBucketParser struct {
	BucketName 		string
	Region     		string
	ACL        		string
	Versioning 		string
	WebSite    		string
	CORs       		string
	UseRole    		string
	AccessId			string
	AccessKey			string
	AccessToken		string
}

func (p *S3CreateBucketParser) Validate() bool {
	flag.StringVar(&p.BucketName, "bucket", "", "Amazon Web Services S3 Bucket Name")
	flag.StringVar(&p.Region, "region", command.DEFAULT_AWS_REGION, "Amazon Web Service reference Region (default : "+command.DEFAULT_AWS_REGION+")")
	flag.StringVar(&p.ACL, "acl", command.DEFAULT_S3_ACL, "Amazon Web ServiceS3 Bucket ACL string (default : "+command.DEFAULT_S3_ACL+")")
	flag.StringVar(&p.CORs, "cors-file", "", "Amazon Web Services S3 CORs YAML file (default : \"\")")
	flag.StringVar(&p.Versioning, "versioning", "disabled", "Amazon Web Services S3 Versioning (default : disabled)")
	flag.StringVar(&p.UseRole, "use-role", "", "Amazon Web Services IAM Role for action (default : \"\")")
	//flag.StringVar(&p.WebSite, "static-website", "", "Amazon Web Service reference Region")
	flag.StringVar(&p.AccessId, "aws-access-id", "", "Amazon Web Services Access Id (default: )")
	flag.StringVar(&p.AccessKey, "aws-access-key", "", "Amazon Web Services Access Key (default: )")
	flag.StringVar(&p.AccessToken, "aws-access-token", "", "Amazon Web Services Access Token (default: )")
	flag.Parse()
	return p.BucketName != ""
}

func (p *S3CreateBucketParser) Parse() []interface{} {
	var arguments []interface{} = make([]interface{}, 0)
	arguments = append(arguments, p.BucketName)
	arguments = append(arguments, p.Region)
	arguments = append(arguments, p.ACL)
	arguments = append(arguments, p.Versioning)
	arguments = append(arguments, p.CORs)
	arguments = append(arguments, p.UseRole)
	//arguments = append(arguments, p.WebSite)
	arguments = append(arguments, p.AccessId)
	arguments = append(arguments, p.AccessKey)
	arguments = append(arguments, p.AccessToken)
	return arguments
}


