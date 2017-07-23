package actions

import (
	"github.com/hellgate75/aws-architect/abstract"
	"github.com/hellgate75/aws-architect/command"
	"github.com/hellgate75/aws-architect/helpers"
)

func InitS3CreateBucket() {
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
		Name:        "acl",
		Description: "Amazon Web Services S3 Bucket ACL string (default : " + defaultS3ACL + ")",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "acl-string",
	}
	var parm4 abstract.Parameter = abstract.Parameter{
		Name:        "versioning",
		Description: "Amazon Web Services S3 Versioning (default : disabled)",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "enabled-or-disabled",
	}
	var parm5 abstract.Parameter = abstract.Parameter{
		Name:        "cors-file",
		Description: "Amazon Web Services S3 CORs YAML file (default : \"\")",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "path-to-cors-yaml-file",
	}
	var parm6 abstract.Parameter = abstract.Parameter{
		Name:        "use-role",
		Description: "Amazon Web Services IAM Role for action (default : \"\")",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "arn-to-role",
	}
	var Parameters []abstract.Parameter = make([]abstract.Parameter, 0)
	Parameters = append(Parameters, parm1)
	Parameters = append(Parameters, parm2)
	Parameters = append(Parameters, parm3)
	Parameters = append(Parameters, parm4)
	Parameters = append(Parameters, parm5)
	Parameters = append(Parameters, parm6)
	var S3CreateBucketAction *abstract.ActionImpl = new(abstract.ActionImpl)
	S3CreateBucketAction.Parameters = Parameters
	S3CreateBucketAction.Name = "Create S3 Bucket"
	S3CreateBucketAction.Command = "create-bucket"
	S3CreateBucketAction.Description = "Create Amazon Web Services S3 Bucket in Amazon Web Services Region"
	S3CreateBucketAction.Usage = helpers.DefineUsage(S3CreateBucketAction.Command, S3CreateBucketAction.Description, S3CreateBucketAction.Parameters)
	S3CreateBucketAction.SetArgumentParser(new(command.S3CreateBucketParser))
	S3CreateBucketAction.SetExecutor(new(command.S3CreateBucket))
	abstract.RegisterAction(S3CreateBucketAction)
}
