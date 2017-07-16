package actions

import (
	"aws-architect/abstract"
	"time"
	"flag"
	"aws-architect/helpers"
	"fmt"
	"aws-architect/aws"
)

type S3CreateBucket struct{
	abstract.Action
	InProgress	bool
	Success			bool
	Message			string
	BucketName	string
	Region			string
	ACL					string
}

func (c S3CreateBucket) Init() (bool) {
	c.InProgress = false
	c.Success = false
	c.Message = ""
	return  true
}

func (c S3CreateBucket) Reset() (bool) {
	c.InProgress = false
	c.Success = false
	c.Message = ""
	return  true
}

func (c *S3CreateBucket) Execute(logChannel chan string) (bool) {
	defer func() {
		// recover from panic caused by writing to a closed channel
		if r := recover(); r != nil {
		}
	}()
	c.InProgress=true
	session := aws.CreateSession();
	awsService := aws.CreateS3Service(session, c.Region)
	logChannel <- fmt.Sprintf("Creating bucket : %s", c.BucketName)
	logChannel <- fmt.Sprintf("Bucket Region : %s", c.Region)
	logChannel <- fmt.Sprintf("Bucket ACL : %s", c.ACL)
	bucketLocation, err := aws.CreateBucket(awsService, c.BucketName, c.Region, c.ACL)
	if err == nil {
		logChannel <- fmt.Sprintf("Bucket '%s' created at locaion : %s", c.BucketName, bucketLocation)
		c.Success=true
		c.InProgress=false
		c.Message = fmt.Sprintf("Creation S3 Bucket %s in Region %s successful!!", c.BucketName,c.Region)
		return  true
	}
	logChannel <- err.Error()
	c.Message = fmt.Sprintf("Error in creating S3 Bucket %s : %s", c.BucketName,err.Error())
	c.Success=false
	c.InProgress=false
	return false
}

func (c *S3CreateBucket) IsInProgress() (bool) {
	return  c.InProgress
}


func (c *S3CreateBucket) GetExitCode() (int) {
	for true {
		if ! c.InProgress {
			break
		}
		time.Sleep(time.Second * 5)
	}
	if c.Success {
		return 0
	}
	return  1
}

func (c *S3CreateBucket) GetCommand() (string) {
	return  c.Command
}

func (c *S3CreateBucket) GetName() (string) {
	return  c.Name
}

func (c *S3CreateBucket) GetUsage() (string) {
	return  c.Usage
}

func (c *S3CreateBucket) AcquireValues() (bool) {
	flag.StringVar(&c.BucketName, "bucket", "", "Amazon Web Services S3 Bucket Name")
	flag.StringVar(&c.Region, "region", defaultAWSRegion, "Amazon Web Service reference Region")
	flag.StringVar(&c.ACL, "acl", defaultS3ACL, "Amazon Web Service reference Region")
	flag.Parse()
	return  c.BucketName != ""
}

func (c *S3CreateBucket) GetLastMessage() (string) {
	return  c.Message
}

func InitS3CreateBucket() {
	var parm1 abstract.Parameter = abstract.Parameter{
		Name: "bucket",
		Description: "Amazon Web Services S3 Bucket Name",
		Mandatory: true,
		HasValue: true,
		SampleValue: "bucket-name",
	}
	var parm2 abstract.Parameter = abstract.Parameter{
		Name: "region",
		Description: "Amazon Web Services reference Region (default : " + defaultAWSRegion + ")",
		Mandatory: false,
		HasValue: true,
		SampleValue: "region-string",
	}
	var parm3 abstract.Parameter = abstract.Parameter{
		Name: "acl",
		Description: "Amazon Web Services S3 ACL string (default : " + defaultS3ACL + ")",
		Mandatory: false,
		HasValue: true,
		SampleValue: "acl-string",
	}
	var Parameters 	[]abstract.Parameter = make([]abstract.Parameter, 0)
	Parameters = append(Parameters, parm1)
	Parameters = append(Parameters, parm2)
	Parameters = append(Parameters, parm3)
	var  S3CreateBucketAction *S3CreateBucket = new (S3CreateBucket)
	S3CreateBucketAction.Parameters= Parameters
	S3CreateBucketAction.Name = "Create S3 Bucket"
	S3CreateBucketAction.Command= "create-bucket"
	S3CreateBucketAction.Description= "Create Amazon Web Services S3 Bucket in Amazon Web Services Region"
	S3CreateBucketAction.Usage=helpers.DefineUsage(S3CreateBucketAction.Command, S3CreateBucketAction.Description, S3CreateBucketAction.Parameters)
	abstract.RegisterAction(S3CreateBucketAction)
}
