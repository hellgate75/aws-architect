package command

import (
	"github.com/hellgate75/aws-architect/abstract"
	"flag"
	"fmt"
	"github.com/hellgate75/aws-architect/aws"
	"strconv"
	"strings"
)

type S3CreateBucket struct {
}

func (p *S3CreateBucket) Execute(action *abstract.ActionImpl,arguments []interface{},logChannel chan string) (bool) {
	var bucketName string = fmt.Sprintf("%v", arguments[0])
	var awsRegion string = fmt.Sprintf("%v", arguments[1])
	var awsAcl string = fmt.Sprintf("%v", arguments[2])
	var s3Versioning string = fmt.Sprintf("%v", arguments[3])
	var s3Cors string = fmt.Sprintf("%v", arguments[4])
	//var s3WebSite string = fmt.Sprintf("%v", arguments[5])
	defer func() {
		// recover from panic caused by writing to a closed channel
		if r := recover(); r != nil {
		}
	}()
	action.InProgress = true
	session := aws.CreateSession()
	awsService := aws.CreateS3Service(session, awsRegion)
	logChannel <- fmt.Sprintf("Creating bucket : %s", bucketName)
	logChannel <- fmt.Sprintf("Bucket Region : %s", awsRegion)
	logChannel <- fmt.Sprintf("Bucket ACL : %s", awsAcl)
	logChannel <- fmt.Sprintf("Bucket Versioning : %s", s3Versioning)
	logChannel <- fmt.Sprintf("Bucket Cors file : %s", s3Cors)
	bucketLocation, err := aws.CreateBucket(awsService, bucketName, awsRegion, awsAcl, strings.ToLower(s3Versioning)=="enabled", s3Cors)
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
	BucketName string
	Region     string
	ACL        string
	Versioning string
	WebSite		 string
	CORs			 string
}

func (p *S3CreateBucketParser) Validate() (bool) {
	flag.StringVar(&p.BucketName, "bucket", "", "Amazon Web Services S3 Bucket Name")
	flag.StringVar(&p.Region, "region", defaultAWSRegion, "Amazon Web Service reference Region (default : " + defaultAWSRegion + ")")
	flag.StringVar(&p.ACL, "acl", defaultS3ACL, "Amazon Web ServiceS3 Bucket ACL string (default : " + defaultS3ACL + ")")
	flag.StringVar(&p.CORs, "cors-file", "", "Amazon Web Services S3 CORs YAML file (default : \"\")")
	flag.StringVar(&p.Versioning, "versioning", "disabled", "Amazon Web Services S3 Versioning (default : disabled)")
	//flag.StringVar(&p.WebSite, "static-website", "", "Amazon Web Service reference Region")
	flag.Parse()
	return p.BucketName != ""
}

func (p *S3CreateBucketParser) Parse() ([]interface{}) {
	var arguments []interface{} = make([]interface{},0)
	arguments = append(arguments, p.BucketName)
	arguments = append(arguments, p.Region)
	arguments = append(arguments, p.ACL)
	arguments = append(arguments, p.Versioning)
	arguments = append(arguments, p.CORs)
	//arguments = append(arguments, p.WebSite)
	return arguments
}

type S3DeleteBucket struct {
}

func (p *S3DeleteBucket) Execute(action *abstract.ActionImpl,arguments []interface{},logChannel chan string) (bool) {
	var bucketName string = fmt.Sprintf("%v", arguments[0])
	var awsRegion string = fmt.Sprintf("%v", arguments[1])
	s3Recursive,_ := strconv.ParseBool(fmt.Sprintf("%v", arguments[2]))
	defer func() {
		// recover from panic caused by writing to a closed channel
		if r := recover(); r != nil {
		}
	}()
	action.InProgress = true
	session := aws.CreateSession()
	awsService := aws.CreateS3Service(session, awsRegion)
	logChannel <- fmt.Sprintf("Deleting bucket : %s", bucketName)
	logChannel <- fmt.Sprintf("Bucket Region : %s", awsRegion)
	var removed bool
	var err error
	if s3Recursive {
		removed, err = aws.DeleteBucketRecursive(awsService, bucketName)
	} else {
		removed, err = aws.DeleteBucket(awsService, bucketName)
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
}

func (p *S3DeleteBucketParser) Validate() (bool) {
	flag.StringVar(&p.BucketName, "bucket", "", "Amazon Web Services S3 Bucket Name")
	flag.StringVar(&p.Region, "region", defaultAWSRegion, "Amazon Web Service reference Region")
	flag.BoolVar(&p.Recursive, "recursive", false, "Delete recursively keys from S3 Bucket")
	flag.Parse()
	return p.BucketName != ""
}

func (p *S3DeleteBucketParser) Parse() ([]interface{}) {
	var arguments []interface{} = make([]interface{},0)
	arguments = append(arguments, p.BucketName)
	arguments = append(arguments, p.Region)
	arguments = append(arguments, p.Recursive)
	return arguments
}