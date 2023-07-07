package awshelp

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func GetSecret(secretId string) (result map[string]string) {
	sess, err := session.NewSession(aws.NewConfig().WithRegion("sa-east-1"))
	if err != nil {
		panic(err)
	}
	svc := secretsmanager.New(sess)
	res, err := svc.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId:     &secretId,
		VersionStage: aws.String("AWSCURRENT"),
	})
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal([]byte(*res.SecretString), &result); err != nil {
		panic(err)
	}
	return
}
