package awsclient

import (
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hixi-hyi/localstack-go/localstack"
)

func TestGlobal(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		os.Setenv("AWS_REGION", "us-east-1")
		ret := APIGateway(session.Must(session.NewSession()))
		if ret.Client.Endpoint != "https://apigateway.us-east-1.amazonaws.com" {
			t.Error("failed default client")
		}
		fmt.Printf("%#v\n", ret.Client.Endpoint)
	})
	t.Run("UseLocalStack", func(t *testing.T) {
		ls := localstack.NewLocalStack()
		UseLocalStack(ls)
		ret := APIGateway(session.Must(session.NewSession()))
		if ret.Client.Endpoint != "http://localhost:4567" {
			t.Error("failed wrap to localstack")
		}
		fmt.Printf("%#v\n", ret.Client.Endpoint)
	})
}
