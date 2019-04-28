# aws-client-go
The library can switch original aws and localstack.


## Example
```
import (
    "github.com/hixi-hyi/aws-client-go/awsclient"
    "github.com/hixi-hyi/localstack-go/localstack"
	"github.com/aws/aws-sdk-go/aws/session"
)
func init() {
    if os.Getenv("AWS_SAM_LOCAL") == "true" {
        awsclient.UseLocalStack(localstack.NewLocalStack())
    }
}
func handler() {
    sess := session.Must(session.NewSession())
    awsclient.APIGateway(sess)
    // awsclient return localstack service if you have already called awsclient.UseLocalStack(). if not awsclient return original aws service.
}
```

## Supported
```
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
```


## See Also
* https://github.com/localstack/localstack
* https://github.com/hixi-hyi/localstack-go
