package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"strings"
)

func convertCfnTags(tags [][]string) []*cloudformation.Tag {
	var cfnTags []*cloudformation.Tag = make([]*cloudformation.Tag, 0)
	for entryIndex := 0; entryIndex < len(tags); entryIndex++ {
		var entry []string = tags[entryIndex]
		if len(entry) > 1 {
			var key string = entry[0]
			var value string = entry[1]
			var tag *cloudformation.Tag = &cloudformation.Tag{
				Key:   aws.String(key),
				Value: aws.String(value),
			}
			cfnTags = append(cfnTags, tag)
		}
	}
	return cfnTags
}

/*
* Parse Expected CloudFormation Tags in format :
* key1=value1/|...|keyN=valueN,
* where key and value parameter, and key-value couple fo values are separated by other ones by "|" character.
* In the command line this group should be bounded by quote marks (")
 */
func parseTags(tagsString string) [][]string {
	var cfnTags [][]string = make([][]string, 0)
	if strings.Index(tagsString, "|") <= 0 || tagsString == "" {
		return cfnTags
	}
	var tokens []string = strings.Split(tagsString, "|")
	for i := 0; i < len(tokens); i++ {
		var tagEntry []string = strings.Split(tokens[i], "=")
		if len(tagEntry) > 1 {
			cfnTags = append(cfnTags, tagEntry)
		}
	}
	return cfnTags
}

func tagsHelper() string {
	return "key1=value1/|...|keyN=valueN,\n" +
		"where key and value represent tag, and key-value couple fo values are separated by other ones by \"|\" character.\n" +
		"In the command line this group should be bounded by quote marks (\")"
}

func convertCfnParams(params [][]string) []*cloudformation.Parameter {
	var cfnParams []*cloudformation.Parameter = make([]*cloudformation.Parameter, 0)
	for entryIndex := 0; entryIndex < len(params); entryIndex++ {
		var entry []string = params[entryIndex]
		if len(entry) > 1 {
			var key string = entry[0]
			var value string = entry[1]
			var usePrevious bool
			if len(entry) > 2 {
				usePrevious = strings.ToLower(entry[2]) == "true" || strings.ToLower(entry[2]) == "yes" || entry[2] == "1"
			}
			var param *cloudformation.Parameter = &cloudformation.Parameter{
				ParameterKey:     aws.String(key),
				ParameterValue:   aws.String(value),
				UsePreviousValue: aws.Bool(usePrevious),
			}
			cfnParams = append(cfnParams, param)
		}
	}
	return cfnParams
}

/*
* Parse Expected CloudFormation Parameters in format :
* key1=value1/[use_previous_value_1]|...|keyN=valueN/[use_previous_value_N],
* where key and value parameter items extended by use_previous_value, that is the option to use previous stack used parameter values.
* key-value-usv tuple value is separated by other ones by "|" character. In the command line this group should be bounded
* by quote marks (")
 */

func parseParams(parametersString string) [][]string {
	var cfnParams [][]string = make([][]string, 0)
	if strings.Index(parametersString, "|") <= 0 || parametersString == "" {
		return cfnParams
	}
	var tokens []string = strings.Split(parametersString, "|")
	for i := 0; i < len(tokens); i++ {
		var tagEntry []string = strings.Split(tokens[i], "=")
		if len(tagEntry) > 1 {
			var tagEntryValue []string = tagEntry
			if strings.Index(tagEntry[1], "\\") >= 0 {
				var tagEntryExtra []string = strings.Split(tagEntry[1], "\\")
				if len(tagEntryExtra) > 1 {
					tagEntryValue[1] = tagEntryExtra[0]
					tagEntryValue = append(tagEntryValue, tagEntryExtra[1])
				}
			}
			cfnParams = append(cfnParams, tagEntryValue)
		}
	}
	return cfnParams
}

func paramsHelper() string {
	return "key1=value1/[use_previous_value_1]|...|keyN=valueN/[use_previous_value_N], \n" +
		"where key and value parameter items extended by use_previous_value, that is the option to use previous stack used \n" +
		"parameter values. key-value-usv tuple value is separated by other ones by \"|\" character. \n" +
		"In the command line this group should be bounded by quote marks (\")"
}

/*
* Parse Expected CloudFormation Resource Types in format :
* ResourceType1,ResourceType2,...,ResourceTypeN
* where ResourceType items are comma separated, in AWS API format (e.g.: AWS::EC2::Instance, AWS::EC2::*, or Custom::MyCustomInstance).
* In the command line this group should be bounded by quote marks (")
 */

func parseResourceTypes(resourceTypesString string) []*string {
	var resTypes []*string = make([]*string, 0)
	if strings.Index(resourceTypesString, ",") <= 0 || resourceTypesString == "" {
		return resTypes
	}
	var tokens []string = strings.Split(resourceTypesString, ",")
	for i := 0; i < len(tokens); i++ {
		var resourceType string = tokens[i]
		resTypes = append(resTypes, aws.String(resourceType))
	}
	return resTypes
}

func resourceTypesHelper() string {
	return "ResourceType1,ResourceType2,...,ResourceTypeN\n" +
		"where ResourceType items are comma separated, in AWS API format (e.g.: AWS::EC2::Instance, AWS::EC2::*, or Custom::MyCustomInstance).\n" +
		"In the command line this group should be bounded by quote marks (\")"
}

/*
* Parse Expected CloudFormation Resource Types in format :
* NotificationArn1,NotificationArn2,...,NotificationArnN
* where NotificationArn items are comma separated, and are reference to notification topics.
* In the command line this group should be bounded by quote marks (")
 */

func parseNotificationArns(notificationArnString string) []*string {
	var notificationArns []*string = make([]*string, 0)
	if strings.Index(notificationArnString, ",") <= 0 || notificationArnString == "" {
		return notificationArns
	}
	var tokens []string = strings.Split(notificationArnString, ",")
	for i := 0; i < len(tokens); i++ {
		var resourceType string = tokens[i]
		notificationArns = append(notificationArns, aws.String(resourceType))
	}
	return notificationArns
}

func notificationArnsHelper() string {
	return "NotificationArn1,NotificationArn2,...,NotificationArnN\n" +
		"where NotificationArn items are comma separated, and are reference to notification topics\n" +
		"In the command line this group should be bounded by quote marks (\")"
}
