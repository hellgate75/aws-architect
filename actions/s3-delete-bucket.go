package actions

import (
	"github.com/hellgate75/aws-architect/abstract"
	"github.com/hellgate75/aws-architect/helpers"
	"github.com/hellgate75/aws-architect/command"
)

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
	var S3DeleteBucketAction *abstract.ActionImpl = new(abstract.ActionImpl)
	S3DeleteBucketAction.Parameters = Parameters
	S3DeleteBucketAction.Name = "Delete S3 Bucket"
	S3DeleteBucketAction.Command = "delete-bucket"
	S3DeleteBucketAction.Description = "Delete Amazon Web Services S3 Bucket from Amazon Web Services Region"
	S3DeleteBucketAction.Usage = helpers.DefineUsage(S3DeleteBucketAction.Command, S3DeleteBucketAction.Description, S3DeleteBucketAction.Parameters)
	S3DeleteBucketAction.SetArgumentParser(new(command.S3DeleteBucketParser))
	S3DeleteBucketAction.SetExecutor(new(command.S3DeleteBucket))
	abstract.RegisterAction(S3DeleteBucketAction)
}
