package awsclient

import "github.com/hixi-hyi/localstack-go/localstack"

var Client AwsClient = &Aws{}

func UseLocalStack(ls *localstack.LocalStack) {
	Client = ls
}
