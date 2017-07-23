package actions

import (
	"github.com/hellgate75/aws-architect/abstract"
	"github.com/hellgate75/aws-architect/command"
	"github.com/hellgate75/aws-architect/helpers"
)

func InitS3BucketStatus() {
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
	var Parameters []abstract.Parameter = make([]abstract.Parameter, 0)
	Parameters = append(Parameters, parm1)
	Parameters = append(Parameters, parm2)
	var S3BucketStatusAction *abstract.ActionImpl = new(abstract.ActionImpl)
	S3BucketStatusAction.Parameters = Parameters
	S3BucketStatusAction.Name = "Status of S3 Bucket"
	S3BucketStatusAction.Command = "bucket-status"
	S3BucketStatusAction.Description = "Recover Location of Amazon Web Services S3 Bucket from Amazon Web Services Region"
	S3BucketStatusAction.Usage = helpers.DefineUsage(S3BucketStatusAction.Command, S3BucketStatusAction.Description, S3BucketStatusAction.Parameters)
	S3BucketStatusAction.SetArgumentParser(new(command.S3BucketStatusParser))
	S3BucketStatusAction.SetExecutor(new(command.S3BucketStatus))
	abstract.RegisterAction(S3BucketStatusAction)
}
