package awsclient

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/aws/aws-sdk-go/service/sms"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/hixi-hyi/localstack-go/localstack"
)

var Client AwsClient = &Aws{}

func UseLocalStack(ls *localstack.LocalStack) {
	Client = ls
}

func APIGateway(s *session.Session) *apigateway.APIGateway {
	return Client.APIGateway(s)
}
func Kinesis(s *session.Session) *kinesis.Kinesis {
	return Client.Kinesis(s)
}

func DynamoDB(s *session.Session) *dynamodb.DynamoDB {
	return Client.DynamoDB(s)
}

func DynamoDBStreams(s *session.Session) *dynamodbstreams.DynamoDBStreams {
	return Client.DynamoDBStreams(s)
}

func S3(s *session.Session) *s3.S3 {
	return Client.S3(s)
}

func Firehose(s *session.Session) *firehose.Firehose {
	return Client.Firehose(s)
}

func Lambda(s *session.Session) *lambda.Lambda {
	return Client.Lambda(s)
}

func SNS(s *session.Session) *sns.SNS {
	return Client.SNS(s)
}

func SQS(s *session.Session) *sqs.SQS {
	return Client.SQS(s)
}

func Redshift(s *session.Session) *redshift.Redshift {
	return Client.Redshift(s)
}

func ElasticsearchService(s *session.Session) *elasticsearchservice.ElasticsearchService {
	return Client.ElasticsearchService(s)
}

func SES(s *session.Session) *ses.SES {
	return Client.SES(s)
}

func Route53(s *session.Session) *route53.Route53 {
	return Client.Route53(s)
}

func CloudFormation(s *session.Session) *cloudformation.CloudFormation {
	return Client.CloudFormation(s)
}

func CloudWatch(s *session.Session) *cloudwatch.CloudWatch {
	return Client.CloudWatch(s)
}

func SMS(s *session.Session) *sms.SMS {
	return Client.SMS(s)
}

func SecretsManager(s *session.Session) *secretsmanager.SecretsManager {
	return Client.SecretsManager(s)
}

func SFN(s *session.Session) *sfn.SFN {
	return Client.SFN(s)
}

func CloudWatchLogs(s *session.Session) *cloudwatchlogs.CloudWatchLogs {
	return Client.CloudWatchLogs(s)
}

func STS(s *session.Session) *sts.STS {
	return Client.STS(s)
}

func IAM(s *session.Session) *iam.IAM {
	return Client.IAM(s)
}
