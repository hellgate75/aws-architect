package s3_bucket

import (
	"github.com/hellgate75/aws-architect/awslet"
	"fmt"
	"flag"
	"github.com/hellgate75/aws-architect/command"
	"github.com/hellgate75/aws-architect/abstract"
	"github.com/aws/aws-sdk-go/service/s3"
	"regexp"
)

type S3ListBuckets struct {
}

func (p *S3ListBuckets) Execute(action *abstract.ActionImpl, arguments []interface{}, logChannel chan string) bool {
	var filter string = fmt.Sprintf("%v", arguments[0])
	var accessId string = fmt.Sprintf("%v", arguments[1])
	var accessKey string = fmt.Sprintf("%v", arguments[2])
	var accessToken string = fmt.Sprintf("%v", arguments[3])

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
	} else {
		awsService = awslet.CreateS3Service(session, awslet.DEFAULT_AWS_REGION)
	}
	logChannel <- "List of all Buckets :"
	if filter != "" {
		logChannel <- fmt.Sprintf("Regular Expression : %s", filter)
	} else  {
		logChannel <- "Name contains : *"
	}
	var buckets []*s3.Bucket
	var err error
	buckets, err = awslet.ListBuckets(awsService)
	if err == nil {
		if len(buckets) > 0 {
			var bucketNames string =""
			var numShownBucket int = 0
			for i:=0; i < len(buckets); i++ {
				var bucket *s3.Bucket = buckets[i]
				var matched bool
				matched, err = regexp.MatchString(filter,*bucket.Name)
				if filter == "" || ( err==nil && matched ) {
					var bucketData string  = fmt.Sprintf("Bucket : %s - Created : %s", *bucket.Name, bucket.CreationDate.Format("2006-01-02 15:04:05.000"))
					logChannel <- bucketData
					bucketNames += "\n" + bucketData
					numShownBucket++
				}
			}
			logChannel <- fmt.Sprintf("Number of buckets: %v", numShownBucket)
			if numShownBucket == 0 {
				logChannel <- "No Bucket found in S3 ..."
				action.Message = "S3 does not contain buckets with selected criteria!!"
			} else {
				action.Message = fmt.Sprintf("\nList of S3 Buckets :  %s", bucketNames)
			}
		} else  {
			logChannel <- "No Bucket found in S3 ..."
			action.Message = "S3 does not contain buckets!!"
		}
		action.Success = true
		action.InProgress = false
		return true
	}
	logChannel <- err.Error()
	action.Message = fmt.Sprintf("Error in recoverying S3 Bucket list : %s", err.Error())
	action.Success = false
	action.InProgress = false
	return false
}

type S3ListBucketsParser struct {
	Region     		string
	Filter		 		string
	AccessId			string
	AccessKey			string
	AccessToken		string
}

func (p *S3ListBucketsParser) Validate() bool {
	flag.StringVar(&p.Filter, "filter", command.DEFAULT_AWS_REGION, "Bucket name regular expression filter string (default : )")
	flag.StringVar(&p.AccessId, "aws-access-id", "", "Amazon Web Services Access Id (default: )")
	flag.StringVar(&p.AccessKey, "aws-access-key", "", "Amazon Web Services Access Key (default: )")
	flag.StringVar(&p.AccessToken, "aws-access-token", "", "Amazon Web Services Access Token (default: )")
	flag.Parse()
	return true
}

func (p *S3ListBucketsParser) Parse() []interface{} {
	var arguments []interface{} = make([]interface{}, 0)
	arguments = append(arguments, p.Filter)
	arguments = append(arguments, p.AccessId)
	arguments = append(arguments, p.AccessKey)
	arguments = append(arguments, p.AccessToken)
	return arguments
}
