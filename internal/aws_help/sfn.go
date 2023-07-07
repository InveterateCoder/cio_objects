package awshelp

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
)

type SF struct {
	region          string
	stateMachineArn string
	maxResults      int64
	sess            *session.Session
	client          *sfn.SFN
}

func NewSF(stateMachineArn string) (sf SF) {
	sf = SF{
		region:          "sa-east-1",
		stateMachineArn: stateMachineArn,
		maxResults:      1000,
		sess: session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		})),
	}
	sf.client = sfn.New(sf.sess, &aws.Config{Region: &sf.region})
	return
}

func (sf *SF) DescribeExecution(executionArn string) *sfn.DescribeExecutionOutput {
	params := sfn.DescribeExecutionInput{
		ExecutionArn: &executionArn,
	}
	ret, err := sf.client.DescribeExecution(&params)
	if err != nil {
		panic(err)
	}
	return ret
}
