package actions

import (
	"github.com/hellgate75/aws-architect/abstract"
	"github.com/hellgate75/aws-architect/awslet"
	"github.com/hellgate75/aws-architect/command/cfn"
	"github.com/hellgate75/aws-architect/helpers"
)

func InitCreateCfnStack() {
	var Parameters []abstract.Parameter = make([]abstract.Parameter, 0)
	var parm1 abstract.Parameter = abstract.Parameter{
		Name:        "stack-name",
		Description: "Amazon Web Services CloudFormation Stack Name",
		Mandatory:   true,
		HasValue:    true,
		SampleValue: "stack-name",
	}
	var parm2 abstract.Parameter = abstract.Parameter{
		Name:        "region",
		Description: "Amazon Web Services reference Region (default : " + defaultAWSRegion + ")",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "region-string",
	}
	var parm3 abstract.Parameter = abstract.Parameter{
		Name:        "cfn-url",
		Description: "Amazon Web Services CloudFormation file URL (default : \"\")",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "http-url-string",
	}
	var parm4 abstract.Parameter = abstract.Parameter{
		Name:        "cfn-path",
		Description: "Amazon Web Services CloudFormation local file Path (default : \"\")",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "local-file-full-path",
	}
	var parm5 abstract.Parameter = abstract.Parameter{
		Name:        "policy-url",
		Description: "Amazon Web Services Policy Document URL (default : \"\")",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "http-url-string",
	}
	var parm6 abstract.Parameter = abstract.Parameter{
		Name:        "policy-file",
		Description: "Amazon Web Services Policy Document local file path (default : \"\")",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "local-file-full-path",
	}
	var parm7 abstract.Parameter = abstract.Parameter{
		Name:        "stack-role-arn",
		Description: "Amazon Web Services IAM Role Arn for Stack operations (default : \"\")",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "arn-to-role",
	}
	var parm8 abstract.Parameter = abstract.Parameter{
		Name:        "timeout",
		Description: "Amazon Web Services IAM Role Arn for Stack operations (default : \"\")",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "minutes",
	}
	var parm9 abstract.Parameter = abstract.Parameter{
		Name:        "disable-rollback",
		Description: "Disable OnError Rollback (default : false)",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "bool",
	}
	var parm10 abstract.Parameter = abstract.Parameter{
		Name:        "client-token",
		Description: "Client Request Token (default : \"\")",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "string",
	}
	var parm11 abstract.Parameter = abstract.Parameter{
		Name:        "notification-arns",
		Description: "List of Amazon Web Services SNS topic ARNs (default : []) - format :\n" + awslet.NotificationArnsHelper(),
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "list",
	}
	var parm12 abstract.Parameter = abstract.Parameter{
		Name:        "resource-types",
		Description: "List of Amazon Web Services Resource Types (default : []) - format :\n" + awslet.ResourceTypesHelper(),
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "list",
	}
	var parm13 abstract.Parameter = abstract.Parameter{
		Name:        "parameters",
		Description: "List of Stack Input Parameters (default : []) - format :\n" + awslet.ParamsHelper(),
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "list",
	}
	var parm14 abstract.Parameter = abstract.Parameter{
		Name:        "tags",
		Description: "List of Stack Tags (default : []) - format :\n" + awslet.TagsHelper(),
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "list",
	}
	var parm15 abstract.Parameter = abstract.Parameter{
		Name:        "use-role",
		Description: "Amazon Web Services IAM Role for action (default : \"\")",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "arn-to-role",
	}
	var parm16 abstract.Parameter = abstract.Parameter{
		Name:        "aws-access-id",
		Description: "Amazon Web Services Access Id (default: )",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "string",
	}
	var parm17 abstract.Parameter = abstract.Parameter{
		Name:        "aws-access-key",
		Description: "Amazon Web Services Access Key (default: )",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "string",
	}
	var parm18 abstract.Parameter = abstract.Parameter{
		Name:        "aws-access-token",
		Description: "Amazon Web Services Access Token (default: )",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "string",
	}
	Parameters = append(Parameters, parm1)
	Parameters = append(Parameters, parm2)
	Parameters = append(Parameters, parm3)
	Parameters = append(Parameters, parm4)
	Parameters = append(Parameters, parm5)
	Parameters = append(Parameters, parm6)
	Parameters = append(Parameters, parm7)
	Parameters = append(Parameters, parm8)
	Parameters = append(Parameters, parm9)
	Parameters = append(Parameters, parm10)
	Parameters = append(Parameters, parm11)
	Parameters = append(Parameters, parm12)
	Parameters = append(Parameters, parm13)
	Parameters = append(Parameters, parm14)
	Parameters = append(Parameters, parm15)
	Parameters = append(Parameters, parm16)
	Parameters = append(Parameters, parm17)
	Parameters = append(Parameters, parm18)

	var CfnCreateStackAction *abstract.ActionImpl = new(abstract.ActionImpl)
	CfnCreateStackAction.Parameters = Parameters
	CfnCreateStackAction.Name = "Create new CloudFormation Stack"
	CfnCreateStackAction.Command = "create-stack"
	CfnCreateStackAction.Description = "Create CloudFormation Stack from local file or remote URL"
	CfnCreateStackAction.Usage = helpers.DefineUsage(CfnCreateStackAction.Command, CfnCreateStackAction.Description, CfnCreateStackAction.Parameters)
	CfnCreateStackAction.SetArgumentParser(new(cfn.CreateStackParser))
	CfnCreateStackAction.SetExecutor(new(cfn.CreateStack))
	abstract.RegisterAction(CfnCreateStackAction)
}
