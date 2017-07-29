package cfn

import (
	"github.com/hellgate75/aws-architect/abstract"
	"github.com/hellgate75/aws-architect/command"
	"fmt"
	"github.com/hellgate75/aws-architect/awslet"
	"flag"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"strconv"
)

type CreateStack struct {
}

func (p *CreateStack) Execute(action *abstract.ActionImpl, arguments []interface{}, logChannel chan string) bool {
	var stackName string = fmt.Sprintf("%v", arguments[0])
	var awsRegion string = fmt.Sprintf("%v", arguments[1])
	var cfnUrl string = fmt.Sprintf("%v", arguments[2])
	var cfnPath string = fmt.Sprintf("%v", arguments[3])
	var policyUrl string = fmt.Sprintf("%v", arguments[4])
	var policyFile string = fmt.Sprintf("%v", arguments[5])
	var stackRoleArn string = fmt.Sprintf("%v", arguments[6])
	var timeout int
	var errN error
	timeout, errN = strconv.Atoi(fmt.Sprintf("%v", arguments[7]))
	if errN != nil {
		timeout = 0
	}
	var disableRollback bool
	disableRollback, errN = strconv.ParseBool(fmt.Sprintf("%v", arguments[8]))
	if errN != nil {
		disableRollback = false
	}
	var cliToken string = fmt.Sprintf("%v", arguments[9])
	var notificationArns []*string = arguments[10].([]*string)
	var resourceTypes []*string = arguments[11].([]*string)
	var parameters [][]string = arguments[12].([][]string)
	var tags [][]string = arguments[13].([][]string)
	var iamRole string = fmt.Sprintf("%v", arguments[14])
	defer func() {
		// recover from panic caused by writing to a closed channel
		if r := recover(); r != nil {
		}
	}()
	action.InProgress = true
	session := awslet.CreateSession()
	var awsService *cloudformation.CloudFormation
	if iamRole == "" {
		awsService = awslet.CreateCfnService(session, awsRegion)
	} else {
		awsService = awslet.CreateCfnServiceAssumeRole(session, awsRegion, iamRole)

	}
	logChannel <- fmt.Sprintf("Creating Stack : %s", stackName)
	logChannel <- fmt.Sprintf("Bucket Region : %s", awsRegion)
	success, err := awslet.CreateStack(awsService, stackName, cfnUrl, cfnPath, policyUrl, policyFile, stackRoleArn,
									timeout, disableRollback, cliToken, notificationArns, resourceTypes, parameters, tags)
	if err == nil {
		logChannel <- fmt.Sprintf("CloudFormation Stack '%s' created : %t", stackName, success)
		action.Success = true
		action.InProgress = false
		if success {
			action.Message = fmt.Sprintf("Creation of CloudFormation Stack %s in Region %s successful!!", stackName, awsRegion)
		} else {
			action.Message = fmt.Sprintf("Creation of CloudFormation Stack %s in Region %s FAILED!!", stackName, awsRegion)
		}
		return true
	}
	logChannel <- err.Error()
	action.Message = fmt.Sprintf("Error creating CloudFormation Stack %s : %s", stackName, err.Error())
	action.Success = false
	action.InProgress = false
	return false
}

type CreateStackParser struct {
	StackName 				string
	Region    				string
	CfnRemoteUrl			string
	CfnFilePath				string
	PolicyUrl					string
	PolicyFile				string
	RoleArn						string
	TimeoutMinutes		int
	DisableRollback		bool
	ClientToken				string
	NotificationARNs	[]*string
	ResourceTypes			[]*string
	Parameters				[][]string
	Tags							[][]string
	UseRole   				string
}

func (p *CreateStackParser) Validate() bool {
	var notificationArnString, resourceTypesString, parametersString, tagsString string
	flag.StringVar(&p.StackName, "stack-name", "", "Amazon Web Services CloudFormation Stack Name")
	flag.StringVar(&p.Region, "region", command.DEFAULT_AWS_REGION, "Amazon Web Service reference Region (default : "+command.DEFAULT_AWS_REGION+")")
	flag.StringVar(&p.CfnRemoteUrl, "cfn-url", "", "Amazon Web Services CloudFormation file URL (default : \"\")")
	flag.StringVar(&p.CfnFilePath, "cfn-path", "", "Amazon Web Services CloudFormation local file Path (default : \"\")")
	flag.StringVar(&p.PolicyUrl, "policy-url", "", "Amazon Web Services Policy Document URL (default : \"\")")
	flag.StringVar(&p.PolicyFile, "policy-file", "", "Amazon Web Services Policy Document local file path (default : \"\")")
	flag.StringVar(&p.RoleArn, "stack-role-arn", "", "Amazon Web Services IAM Role Arn for Stack operations (default : \"\")")
	flag.IntVar(&p.TimeoutMinutes, "timeout", 0, "Timeout related to Stack creation process time, no timeout when zero (default : 0)")
	flag.BoolVar(&p.DisableRollback, "disable-rollback", false, "Disable OnError Rollback (default : false)")
	flag.StringVar(&p.ClientToken, "client-token", "", "Client Request Token (default : \"\")")
	flag.StringVar(&notificationArnString, "notification-arns", "", "List of Amazon Web Services SNS topic ARNs (default : []) - format :\n" + awslet.NotificationArnsHelper())
	flag.StringVar(&resourceTypesString, "resource-types", "", "List of Amazon Web Services Resource Types (default : []) - format :\n" + awslet.ResourceTypesHelper())
	flag.StringVar(&parametersString, "parameters", "", "List of Stack Input Parameters (default : []) - format :\n" + awslet.ParamsHelper())
	flag.StringVar(&tagsString, "tags", "", "List of Stack Tags (default : []) - format :\n" + awslet.TagsHelper())
	flag.StringVar(&p.UseRole, "use-role", "", "Amazon Web Services IAM Role for action (default : \"\")")
	flag.Parse()
	p.NotificationARNs = awslet.ParseNotificationArns(notificationArnString)
	p.ResourceTypes = awslet.ParseResourceTypes(resourceTypesString)
	p.Parameters = awslet.ParseParams(parametersString)
	p.Tags = awslet.ParseTags(tagsString)
	return p.StackName != "" && (p.CfnRemoteUrl != "" || p.CfnFilePath != "")
}

func (p *CreateStackParser) Parse() []interface{} {
	var arguments []interface{} = make([]interface{}, 0)
	arguments = append(arguments, p.StackName)
	arguments = append(arguments, p.Region)
	arguments = append(arguments, p.CfnRemoteUrl)
	arguments = append(arguments, p.CfnFilePath)
	arguments = append(arguments, p.PolicyUrl)
	arguments = append(arguments, p.PolicyFile)
	arguments = append(arguments, p.RoleArn)
	arguments = append(arguments, p.TimeoutMinutes)
	arguments = append(arguments, p.DisableRollback)
	arguments = append(arguments, p.ClientToken)
	arguments = append(arguments, p.NotificationARNs)
	arguments = append(arguments, p.ResourceTypes)
	arguments = append(arguments, p.Parameters)
	arguments = append(arguments, p.Tags)
	arguments = append(arguments, p.UseRole)
	return arguments
}

