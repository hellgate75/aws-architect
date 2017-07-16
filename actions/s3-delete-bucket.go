package actions

import (
	"aws-architect/abstract"
	"aws-architect/aws"
	"aws-architect/helpers"
	"flag"
	"fmt"
	"time"
)

type S3DeleteBucket struct {
	abstract.Action
	InProgress bool
	Success    bool
	Message    string
	BucketName string
	Region     string
	Recursive  bool
}

func (c S3DeleteBucket) Init() bool {
	c.InProgress = false
	c.Success = false
	c.Message = ""
	return true
}

func (c S3DeleteBucket) Reset() bool {
	c.InProgress = false
	c.Success = false
	c.Message = ""
	return true
}

func (c *S3DeleteBucket) Execute(logChannel chan string) bool {
	defer func() {
		// recover from panic caused by writing to a closed channel
		if r := recover(); r != nil {
		}
	}()
	c.InProgress = true
	session := aws.CreateSession()
	awsService := aws.CreateS3Service(session, c.Region)
	logChannel <- fmt.Sprintf("Deleting bucket : %s", c.BucketName)
	logChannel <- fmt.Sprintf("Bucket Region : %s", c.Region)
	var removed bool
	var err error
	if c.Recursive {
		removed, err = aws.DeleteBucketRecursive(awsService, c.BucketName)
	} else {
		removed, err = aws.DeleteBucket(awsService, c.BucketName)
	}
	if err == nil {
		logChannel <- fmt.Sprintf("Bucket '%s' removed : %t", c.BucketName, removed)
		c.Success = true
		c.InProgress = false
		c.Message = fmt.Sprintf("Deletion of S3 Bucket %s in Region %s successful!!", c.BucketName, c.Region)
		return true
	}
	logChannel <- err.Error()
	c.Message = fmt.Sprintf("Error in deleting S3 Bucket %s : %s", c.BucketName, err.Error())
	c.Success = false
	c.InProgress = false
	return false
}

func (c *S3DeleteBucket) IsInProgress() bool {
	return c.InProgress
}

func (c *S3DeleteBucket) GetExitCode() int {
	for true {
		if !c.InProgress {
			break
		}
		time.Sleep(time.Second * 5)
	}
	if c.Success {
		return 0
	}
	return 1
}

func (c *S3DeleteBucket) GetCommand() string {
	return c.Command
}

func (c *S3DeleteBucket) GetName() string {
	return c.Name
}

func (c *S3DeleteBucket) GetUsage() string {
	return c.Usage
}

func (c *S3DeleteBucket) AcquireValues() bool {
	flag.StringVar(&c.BucketName, "bucket", "", "Amazon Web Services S3 Bucket Name")
	flag.StringVar(&c.Region, "region", defaultAWSRegion, "Amazon Web Service reference Region")
	flag.BoolVar(&c.Recursive, "recursive", false, "Delete recursively keys from S3 Bucket")
	flag.Parse()
	return c.BucketName != ""
}

func (c *S3DeleteBucket) GetLastMessage() string {
	return c.Message
}

func InitS3DeleteBucket() {
	var parm1 abstract.Parameter = abstract.Parameter{
		Name:        "bucket",
		Description: "Amazon Web Services S3 Bucket Name",
		Mandatory:   true,
		HasValue:    true,
		SampleValue: "bucket-name",
	}
	var parm2 abstract.Parameter = abstract.Parameter{
		Name:        "region",
		Description: "Amazon Web Services reference Region (default : " + defaultAWSRegion + ")",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "region-string",
	}
	var parm3 abstract.Parameter = abstract.Parameter{
		Name:        "recursive",
		Description: "Delete recursively keys from S3 Bucket (default: false)",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "true-or-false",
	}
	var Parameters []abstract.Parameter = make([]abstract.Parameter, 0)
	Parameters = append(Parameters, parm1)
	Parameters = append(Parameters, parm2)
	Parameters = append(Parameters, parm3)
	var S3DeleteBucketAction *S3DeleteBucket = new(S3DeleteBucket)
	S3DeleteBucketAction.Parameters = Parameters
	S3DeleteBucketAction.Name = "Delete S3 Bucket"
	S3DeleteBucketAction.Command = "delete-bucket"
	S3DeleteBucketAction.Description = "Delete Amazon Web Services S3 Bucket from Amazon Web Services Region"
	S3DeleteBucketAction.Usage = helpers.DefineUsage(S3DeleteBucketAction.Command, S3DeleteBucketAction.Description, S3DeleteBucketAction.Parameters)
	abstract.RegisterAction(S3DeleteBucketAction)
}
