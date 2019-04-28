package awsclient

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hixi-hyi/localstack-go/localstack"
)

func TestUseLocalStack(t *testing.T) {
	t.Run("Interface", func(t *testing.T) {
		ls := localstack.NewLocalStack()
		ret := Client.APIGateway(session.Must(session.NewSession()))
		if ret.Client.Endpoint == "http://localhost:4567" {
			t.Error("failed wrap to localstack")
		}
		fmt.Printf("%#v", ret.Client.Endpoint)
	})
}
