package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

func GetSSMServiceClient() ssmiface.SSMAPI {
	sess := session.Must(session.NewSessionWithOptions(session.Options{}))
	svc := ssm.New(sess)
	return svc
}

func GetParameter(svc ssmiface.SSMAPI, name *string, decrypt bool) (*ssm.GetParameterOutput, error) {
	results, err := svc.GetParameter(
		&ssm.GetParameterInput{
			Name:           name,
			WithDecryption: &decrypt,
		},
	)
	return results, err
}
