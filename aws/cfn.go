package aws

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/hellgate75/aws-architect/util"
)

func CreateStack(service *cloudformation.CloudFormation, stackName string, remoteUrl string, localFilePath string,
	policyURL string, policyBodyLocalFile string, stackRoleArn string, timeoutInMinutes int,
	disableRollback bool, clientRequestToken string,
	notificationArns []*string, resourceTypes []*string, parameters [][]string, tags [][]string) (bool, error) {
	var status bool
	var err error
	var capabilities []*string = make([]*string, 0)
	capabilities = append(capabilities, aws.String("CAPABILITY_IAM"))
	var createStackInput *cloudformation.CreateStackInput = &cloudformation.CreateStackInput{
		Capabilities: capabilities,
		StackName:    aws.String(stackName),
		Tags:         convertCfnTags(tags),
		Parameters:   convertCfnParams(parameters),
	}
	if remoteUrl != "" {
		createStackInput.TemplateURL = aws.String(remoteUrl)
	} else if localFilePath != "" {
		cfnBody, err := util.LoadFileContent(localFilePath)
		if err != nil {
			logger.ErrorS(fmt.Sprintf("Stack Name : %s => Error Loading Local CloudFormation File %s", stackName, localFilePath))
			return false, err
		}
		createStackInput.TemplateBody = aws.String(cfnBody)

	} else {
		return false, errors.New("Cloud formation create command requires CloudFormation url or local CloudFormation file path")
	}
	if stackRoleArn != "" {
		createStackInput.RoleARN = aws.String(stackRoleArn)
	}
	if len(resourceTypes) > 0 {
		createStackInput.ResourceTypes = resourceTypes
	}
	if policyURL != "" {
		createStackInput.StackPolicyURL = aws.String(policyURL)
	}
	if timeoutInMinutes > 0 {
		createStackInput.TimeoutInMinutes = aws.Int64(int64(timeoutInMinutes))
	}

	createStackInput.DisableRollback = aws.Bool(disableRollback)

	if clientRequestToken != "" {
		createStackInput.ClientRequestToken = aws.String(clientRequestToken)
	}

	if len(notificationArns) > 0 {
		createStackInput.NotificationARNs = notificationArns
	}

	if policyBodyLocalFile != "" {
		policyBody, err := util.LoadFileContent(policyBodyLocalFile)
		if err == nil {
			createStackInput.StackPolicyBody = aws.String(policyBody)
		} else {
			logger.Warning(fmt.Sprintf("Stack Name : %s => Error Loading Policy Body File %s", stackName, policyBodyLocalFile))
			logger.WarningE(err)
		}
	}
	stackOutput, err := service.CreateStack(createStackInput)
	if err == nil {
		logger.Log(fmt.Sprintf("Stack Name : %s => CloudFormation Loaded Successfully", stackName))
		logger.Log(fmt.Sprintf("Stack Created => Id: %s", *stackOutput.StackId))
		err = service.WaitUntilStackCreateComplete(&cloudformation.DescribeStacksInput{
			StackName: aws.String(stackName),
		})
		if err != nil {
			logger.ErrorS(fmt.Sprintf("Stack Id : %s => Waiting Creation Error", *stackOutput.StackId))
			logger.Error(err)
			return status, err
		} else {
			status = true
			logger.Log(fmt.Sprintf("Stack Id : %s => Creation Success", *stackOutput.StackId))
		}

	} else {
		logger.ErrorS(fmt.Sprintf("Stack Name : %s => CloudFormation Load Error", stackName))
		logger.Error(err)
	}
	return status, err
}

func UpdateStack(service *cloudformation.CloudFormation, stackName string, remoteUrl string, localFilePath string, resusePreviousTemplate bool,
	policyURL string, policyBodyLocalFile string, stackRoleArn string, policyUpdateURL string, policyUpdateBodyLocalFile string,
	clientRequestToken string, notificationArns []*string, resourceTypes []*string, parameters [][]string, tags [][]string) (bool, error) {
	var status bool
	var err error
	var capabilities []*string = make([]*string, 0)
	capabilities = append(capabilities, aws.String("CAPABILITY_IAM"))
	var updateStackInput *cloudformation.UpdateStackInput = &cloudformation.UpdateStackInput{
		Capabilities: capabilities,
		StackName:    aws.String(stackName),
		Tags:         convertCfnTags(tags),
		Parameters:   convertCfnParams(parameters),
	}
	if !resusePreviousTemplate {
		if remoteUrl != "" {
			updateStackInput.TemplateURL = aws.String(remoteUrl)
		} else if localFilePath != "" {
			cfnBody, err := util.LoadFileContent(localFilePath)
			if err != nil {
				logger.ErrorS(fmt.Sprintf("Stack Name : %s => Error Loading Local CloudFormation File %s", stackName, localFilePath))
				return false, err
			}
			updateStackInput.TemplateBody = aws.String(cfnBody)

		} else {
			return false, errors.New("Cloud formation create command requires CloudFormation url or local CloudFormation file path")
		}
	}
	if stackRoleArn != "" {
		updateStackInput.RoleARN = aws.String(stackRoleArn)
	}
	if len(resourceTypes) > 0 {
		updateStackInput.ResourceTypes = resourceTypes
	}

	if policyURL != "" {
		updateStackInput.StackPolicyURL = aws.String(policyURL)
	}

	if policyUpdateURL != "" {
		updateStackInput.StackPolicyDuringUpdateURL = aws.String(policyUpdateURL)
	}

	if clientRequestToken != "" {
		updateStackInput.ClientRequestToken = aws.String(clientRequestToken)
	}

	if len(notificationArns) > 0 {
		updateStackInput.NotificationARNs = notificationArns
	}

	updateStackInput.UsePreviousTemplate = aws.Bool(resusePreviousTemplate)

	if policyUpdateBodyLocalFile != "" {
		policyBody, err := util.LoadFileContent(policyUpdateBodyLocalFile)
		if err == nil {
			updateStackInput.StackPolicyDuringUpdateBody = aws.String(policyBody)
		} else {
			logger.Warning(fmt.Sprintf("Stack Name : %s => Error Loading Policy During Update Body File %s", stackName, policyUpdateBodyLocalFile))
			logger.WarningE(err)
		}
	}

	if policyBodyLocalFile != "" {
		policyBody, err := util.LoadFileContent(policyBodyLocalFile)
		if err == nil {
			updateStackInput.StackPolicyBody = aws.String(policyBody)
		} else {
			logger.Warning(fmt.Sprintf("Stack Name : %s => Error Loading Policy Body File %s", stackName, policyBodyLocalFile))
			logger.WarningE(err)
		}
	}
	stackOutput, err := service.UpdateStack(updateStackInput)
	if err == nil {
		logger.Log(fmt.Sprintf("Stack Name : %s => CloudFormation Loaded Successfully", stackName))
		logger.Log(fmt.Sprintf("Stack Updated => Id: %s", *stackOutput.StackId))
		err = service.WaitUntilStackUpdateComplete(&cloudformation.DescribeStacksInput{
			StackName: aws.String(stackName),
		})
		if err != nil {
			logger.ErrorS(fmt.Sprintf("Stack Id : %s => Waiting Update Error", *stackOutput.StackId))
			logger.Error(err)
			return status, err
		} else {
			status = true
			logger.Log(fmt.Sprintf("Stack Id : %s => Update Success", *stackOutput.StackId))
		}

	} else {
		logger.ErrorS(fmt.Sprintf("Stack Name : %s => CloudFormation Load Error", stackName))
		logger.Error(err)
	}
	return status, err
}

func StackStatus(service *cloudformation.CloudFormation, stackName string) (bool, error) {
	var status bool
	var err error
	var input *cloudformation.DescribeStacksInput = &cloudformation.DescribeStacksInput{
		StackName: aws.String(stackName),
	}
	output, err := service.DescribeStacks(input)
	if err == nil {
		status = true
		if len(output.Stacks) > 0 {
			for i := 0; i < len(output.Stacks); i++ {
				stack := output.Stacks[i]
				var stackOutputs []*cloudformation.Output = stack.Outputs
				var outputsString string
				for j := 0; j < len(stackOutputs); j++ {
					var stackOutput *cloudformation.Output = stackOutputs[j]
					outputsString += "Key: " + *stackOutput.OutputKey
					outputsString += ",Value: " + *stackOutput.OutputValue
					outputsString += ",Description: " + *stackOutput.Description + "\n"
				}
				var stackParameters []*cloudformation.Parameter = stack.Parameters
				var parametersString string
				for j := 0; j < len(stackParameters); j++ {
					var stackParameter *cloudformation.Parameter = stackParameters[j]
					parametersString += "Key: " + *stackParameter.ParameterKey
					parametersString += ",Value: " + *stackParameter.ParameterValue
					parametersString += fmt.Sprintf(",UsePreviousValue: %t \n", *stackParameter.UsePreviousValue)
				}
				logger.Log(fmt.Sprintf("Stack Name : %s\nDescription : %s\nCreation : %s\nLast Update : %s\nStatus: %s\nReason: %s\nParameters: %s\nOutputs: %s\n", *stack.StackName, *stack.Description, (*stack.CreationTime).Format("2006-01-02 15:04:05.000"), (*stack.LastUpdatedTime).Format("2006-01-02 15:04:05.000"), *stack.StackStatus, *stack.StackStatusReason, parametersString, outputsString))
			}
		} else {
			logger.Warning(fmt.Sprintf("Stack Name : %s => No record found in Amazon Web Services", stackName))
		}
	} else {
		logger.ErrorS(fmt.Sprintf("Stack Name : %s => Error in recovery information", stackName))
		logger.Error(err)
	}
	return status, err
}

func DeleteStack(service *cloudformation.CloudFormation, stackName string) (bool, error) {
	var status bool
	var err error
	var input *cloudformation.DeleteStackInput = &cloudformation.DeleteStackInput{
		StackName: aws.String(stackName),
	}
	output, err := service.DeleteStack(input)
	if err == nil {
		err = service.WaitUntilStackDeleteComplete(&cloudformation.DescribeStacksInput{
			StackName: aws.String(stackName),
		})
		if err == nil {
			status = true
			logger.Log(fmt.Sprintf("Stack Name : %s => Delete Stack successful", stackName))
			logger.Log(fmt.Sprintf("Response : %s", output.String()))
		} else {
			logger.ErrorS(fmt.Sprintf("Stack Name : %s => Error Waiting Stack Deletion", stackName))
			logger.Error(err)
			return status, err
		}
	} else {
		logger.ErrorS(fmt.Sprintf("Stack Name : %s => Error in Stack Deletion", stackName))
		logger.Error(err)
	}
	return status, err
}
