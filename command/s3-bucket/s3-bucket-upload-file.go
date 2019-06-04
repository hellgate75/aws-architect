package s3_bucket

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hellgate75/aws-architect/abstract"
	"github.com/hellgate75/aws-architect/awslet"
	"github.com/hellgate75/aws-architect/command"
	"time"
)

type S3BucketUpload struct {
}

func (p *S3BucketUpload) Execute(action *abstract.ActionImpl, arguments []interface{}, logChannel chan string) bool {
	var bucketName string = fmt.Sprintf("%v", arguments[0])
	var awsRegion string = fmt.Sprintf("%v", arguments[1])
	var iamRole string = fmt.Sprintf("%v", arguments[2])
	var keyName string = fmt.Sprintf("%v", arguments[3])
	var filePath string = fmt.Sprintf("%v", arguments[4])
	var acl string = fmt.Sprintf("%v", arguments[5])
	var storageClass string = fmt.Sprintf("%v", arguments[6])
	var multipart bool = arguments[7].(bool)
	var accessId string = fmt.Sprintf("%v", arguments[8])
	var accessKey string = fmt.Sprintf("%v", arguments[9])
	var accessToken string = fmt.Sprintf("%v", arguments[10])

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
		awsService, credErr = awslet.CreateS3ServiceWithCredentials(session, awsRegion, accessId, accessKey, accessToken)
		if credErr != nil {
			logChannel <- credErr.Error()
			action.Message = fmt.Sprintf("Credential Error for ID :  %s connectiong to S3 Bucket %s : %s", accessId, bucketName, credErr.Error())
			action.Success = false
			action.InProgress = false
			return false
		}
	} else if iamRole == "" {
		awsService = awslet.CreateS3Service(session, awsRegion)
	} else {
		awsService = awslet.CreateS3ServiceAssumeRole(session, awsRegion, iamRole)
	}
	logChannel <- fmt.Sprintf("S3 Bucket : %s", bucketName)
	time.Sleep(200 * time.Microsecond)
	logChannel <- fmt.Sprintf("S3 Bucket Region : %s", awsRegion)
	time.Sleep(200 * time.Microsecond)
	logChannel <- fmt.Sprintf("Key Name : %s", keyName)
	time.Sleep(200 * time.Microsecond)
	logChannel <- fmt.Sprintf("Local File : %s", filePath)
	time.Sleep(200 * time.Microsecond)
	logChannel <- fmt.Sprintf("ACL : %s", acl)
	time.Sleep(200 * time.Microsecond)
	logChannel <- fmt.Sprintf("Storage Class : %s", storageClass)
	time.Sleep(200 * time.Microsecond)
	logChannel <- fmt.Sprintf("Is Multipart : %t", multipart)
	var exists bool
	var err error
	exists, err = awslet.BucketUpload(awsService, awsRegion, acl, filePath, keyName, storageClass, multipart)
	if err == nil {
		logChannel <- fmt.Sprintf("Bucket key '%s' with file Upload '%s' uploaded : %s", keyName, filePath, exists)
		action.Message = fmt.Sprintf("Upload file %s in S3 Bucket %s in Region %s : %t!!", keyName, bucketName, awsRegion, exists)
		action.Success = exists
		action.InProgress = false
		return true
	}
	logChannel <- err.Error()
	action.Message = fmt.Sprintf("Error in uploading file %s to S3 Bucket %s : %s", filePath, bucketName, err.Error())
	action.Success = false
	action.InProgress = false
	return false
}

type S3BucketUploadParser struct {
	BucketName   string
	Region       string
	UseRole      string
	KeyName      string
	FilePath     string
	Acl          string
	StorageClass string
	Multipart    bool
	AccessId     string
	AccessKey    string
	AccessToken  string
}

func (p *S3BucketUploadParser) Validate() bool {
	flag.StringVar(&p.BucketName, "bucket", "", "Amazon Web Services S3 Bucket Name")
	flag.StringVar(&p.Region, "region", command.DEFAULT_AWS_REGION, "Amazon Web Service reference Region (default : "+command.DEFAULT_AWS_REGION+")")
	flag.StringVar(&p.UseRole, "use-role", "", "Amazon Web Services IAM Role for action (default : \"\")")
	flag.StringVar(&p.KeyName, "key", "", "Amazon Web Services S3 Bucket Key name (default : )")
	flag.StringVar(&p.FilePath, "file", "", "Full qualified file path to upload (default : )")
	//var aclCtrlVals string = "READ WRITE READ_ACP WRITE_ACP FULL_CONTROL"
	flag.StringVar(&p.Acl, "acl", command.DEFAULT_ACL_TYPE, "Amazon Web Services S3 Bucket Object Acl (Available Values : "+command.ACL_TYPE_STRING+" - default: "+command.DEFAULT_ACL_TYPE+")")
	flag.StringVar(&p.StorageClass, "storage-class", command.DEFAULT_STORAGE_TYPE, "Amazon Web Services S3 Storage Class (Available Values: "+command.STORAGE_TYPE_STRING+" default: "+command.DEFAULT_STORAGE_TYPE+")")
	flag.StringVar(&p.AccessId, "aws-access-id", "", "Amazon Web Services Access Id (default: )")
	flag.StringVar(&p.AccessKey, "aws-access-key", "", "Amazon Web Services Access Key (default: )")
	flag.StringVar(&p.AccessToken, "aws-access-token", "", "Amazon Web Services Access Token (default: )")
	flag.Parse()
	return p.BucketName != ""
}

func (p *S3BucketUploadParser) Parse() []interface{} {
	var arguments []interface{} = make([]interface{}, 0)
	arguments = append(arguments, p.BucketName)
	arguments = append(arguments, p.Region)
	arguments = append(arguments, p.UseRole)
	arguments = append(arguments, p.KeyName)
	arguments = append(arguments, p.FilePath)
	arguments = append(arguments, p.Acl)
	arguments = append(arguments, p.StorageClass)
	arguments = append(arguments, p.Multipart)
	arguments = append(arguments, p.AccessId)
	arguments = append(arguments, p.AccessKey)
	arguments = append(arguments, p.AccessToken)
	return arguments
}
