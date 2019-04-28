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
)

type AwsClient interface {
	APIGateway(*session.Session) *apigateway.APIGateway
	Kinesis(*session.Session) *kinesis.Kinesis
	DynamoDB(*session.Session) *dynamodb.DynamoDB
	DynamoDBStreams(*session.Session) *dynamodbstreams.DynamoDBStreams
	S3(*session.Session) *s3.S3
	Firehose(*session.Session) *firehose.Firehose
	Lambda(*session.Session) *lambda.Lambda
	SNS(*session.Session) *sns.SNS
	SQS(*session.Session) *sqs.SQS
	Redshift(*session.Session) *redshift.Redshift
	ElasticsearchService(*session.Session) *elasticsearchservice.ElasticsearchService
	SES(*session.Session) *ses.SES
	Route53(*session.Session) *route53.Route53
	CloudFormation(*session.Session) *cloudformation.CloudFormation
	CloudWatch(*session.Session) *cloudwatch.CloudWatch
	SMS(*session.Session) *sms.SMS
	SecretsManager(*session.Session) *secretsmanager.SecretsManager
	SFN(*session.Session) *sfn.SFN
	CloudWatchLogs(*session.Session) *cloudwatchlogs.CloudWatchLogs
	STS(*session.Session) *sts.STS
	IAM(*session.Session) *iam.IAM
}
