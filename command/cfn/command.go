package command

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cfn "github.com/aws/aws-sdk-go/service/cloudformation"
	log "github.com/cultureamp/aws-nquire/logging"
	"os"
	"strings"
)

func Run(stack string, key string, region string) {
	svc := cfn.New(session.New(&aws.Config{Region: aws.String(region)}))
	resp := queryStackOutputs(svc, stack)
	outputs := getOutputs(resp.Stacks)
	findByKey(key, outputs)
}

func findByKey(k string, outputs []*cfn.Output) {
	for _, output := range outputs {
		if strings.EqualFold(*output.OutputKey, k) {
			fmt.Println(*output.OutputValue)
		}
	}
}

func queryStackOutputs(svc *cfn.CloudFormation, stack string) *cfn.DescribeStacksOutput {
	inputs := params(stack)
	resp, err := svc.DescribeStacks(inputs)
	if err != nil {
		log.Error("Error in describing stacks")
		panic(err)
	}

	return resp
}

func getOutputs(stacks []*cfn.Stack) []*cfn.Output {
	if len(stacks) != 1 {
		log.Error("Either too little or too many stacks found")
		os.Exit(1)
	}
	return stacks[0].Outputs
}

func params(name string) *cfn.DescribeStacksInput {
	return &cfn.DescribeStacksInput{
		NextToken: aws.String("NextToken"),
		StackName: aws.String(name),
	}
}
