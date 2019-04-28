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

type Aws struct {
	AwsClient
}

func (c *Aws) APIGateway(s *session.Session) *apigateway.APIGateway {
	return apigateway.New(s)
}
func (c *Aws) Kinesis(s *session.Session) *kinesis.Kinesis {
	return kinesis.New(s)
}

func (c *Aws) DynamoDB(s *session.Session) *dynamodb.DynamoDB {
	return dynamodb.New(s)
}

func (c *Aws) DynamoDBStreams(s *session.Session) *dynamodbstreams.DynamoDBStreams {
	return dynamodbstreams.New(s)
}

func (c *Aws) S3(s *session.Session) *s3.S3 {
	return s3.New(s)
}

func (c *Aws) Firehose(s *session.Session) *firehose.Firehose {
	return firehose.New(s)
}

func (c *Aws) Lambda(s *session.Session) *lambda.Lambda {
	return lambda.New(s)
}

func (c *Aws) SNS(s *session.Session) *sns.SNS {
	return sns.New(s)
}

func (c *Aws) SQS(s *session.Session) *sqs.SQS {
	return sqs.New(s)
}

func (c *Aws) Redshift(s *session.Session) *redshift.Redshift {
	return redshift.New(s)
}

func (c *Aws) ElasticsearchService(s *session.Session) *elasticsearchservice.ElasticsearchService {
	return elasticsearchservice.New(s)
}
func (c *Aws) SES(s *session.Session) *ses.SES {
	return ses.New(s)
}

func (c *Aws) Route53(s *session.Session) *route53.Route53 {
	return route53.New(s)
}

func (c *Aws) CloudFormation(s *session.Session) *cloudformation.CloudFormation {
	return cloudformation.New(s)
}

func (c *Aws) CloudWatch(s *session.Session) *cloudwatch.CloudWatch {
	return cloudwatch.New(s)
}

func (c *Aws) SMS(s *session.Session) *sms.SMS {
	return sms.New(s)
}

func (c *Aws) SecretsManager(s *session.Session) *secretsmanager.SecretsManager {
	return secretsmanager.New(s)
}

func (c *Aws) SFN(s *session.Session) *sfn.SFN {
	return sfn.New(s)
}

func (c *Aws) CloudWatchLogs(s *session.Session) *cloudwatchlogs.CloudWatchLogs {
	return cloudwatchlogs.New(s)
}

func (c *Aws) STS(s *session.Session) *sts.STS {
	return sts.New(s)
}

func (c *Aws) IAM(s *session.Session) *iam.IAM {
	return iam.New(s)
}
