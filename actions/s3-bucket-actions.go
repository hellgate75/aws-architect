package actions

import (
	"github.com/hellgate75/aws-architect/abstract"
	"github.com/hellgate75/aws-architect/helpers"
	"github.com/hellgate75/aws-architect/command/s3-bucket"
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
	S3BucketStatusAction.Description = "Recover Location Constraint of Amazon Web Services S3 Bucket from Amazon Web Services Region"
	S3BucketStatusAction.Usage = helpers.DefineUsage(S3BucketStatusAction.Command, S3BucketStatusAction.Description, S3BucketStatusAction.Parameters)
	S3BucketStatusAction.SetArgumentParser(new(s3_bucket.S3BucketStatusParser))
	S3BucketStatusAction.SetExecutor(new(s3_bucket.S3BucketStatus))
	abstract.RegisterAction(S3BucketStatusAction)
}

func InitListS3Buckets() {
	var parm1 abstract.Parameter = abstract.Parameter{
		Name:        "filter",
		Description: "Bucket name regular expression filter string (default : )",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "bucket-name-fraction-string",
	}
	var Parameters []abstract.Parameter = make([]abstract.Parameter, 0)
	Parameters = append(Parameters, parm1)
	var S3ListBucketsAction *abstract.ActionImpl = new(abstract.ActionImpl)
	S3ListBucketsAction.Parameters = Parameters
	S3ListBucketsAction.Name = "List of S3 Buckets"
	S3ListBucketsAction.Command = "buckets-list"
	S3ListBucketsAction.Description = "Recover List of Amazon Web Services S3 Buckets"
	S3ListBucketsAction.Usage = helpers.DefineUsage(S3ListBucketsAction.Command, S3ListBucketsAction.Description, S3ListBucketsAction.Parameters)
	S3ListBucketsAction.SetArgumentParser(new(s3_bucket.S3ListBucketsParser))
	S3ListBucketsAction.SetExecutor(new(s3_bucket.S3ListBuckets))
	abstract.RegisterAction(S3ListBucketsAction)
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
	var parm4 abstract.Parameter = abstract.Parameter{
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
	var S3DeleteBucketAction *abstract.ActionImpl = new(abstract.ActionImpl)
	S3DeleteBucketAction.Parameters = Parameters
	S3DeleteBucketAction.Name = "Delete S3 Bucket"
	S3DeleteBucketAction.Command = "delete-bucket"
	S3DeleteBucketAction.Description = "Delete Amazon Web Services S3 Bucket from Amazon Web Services Region"
	S3DeleteBucketAction.Usage = helpers.DefineUsage(S3DeleteBucketAction.Command, S3DeleteBucketAction.Description, S3DeleteBucketAction.Parameters)
	S3DeleteBucketAction.SetArgumentParser(new(s3_bucket.S3DeleteBucketParser))
	S3DeleteBucketAction.SetExecutor(new(s3_bucket.S3DeleteBucket))
	abstract.RegisterAction(S3DeleteBucketAction)
}

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
	S3CreateBucketAction.SetArgumentParser(new(s3_bucket.S3CreateBucketParser))
	S3CreateBucketAction.SetExecutor(new(s3_bucket.S3CreateBucket))
	abstract.RegisterAction(S3CreateBucketAction)
}
