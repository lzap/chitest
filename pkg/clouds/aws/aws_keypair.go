package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ImportSSHKey(body string) (string, error) {
	input := &ec2.ImportKeyPairInput{}
	input.KeyName = aws.String("Red Hat Portal Key")
	input.PublicKeyMaterial = []byte(body)
	output, err := Client.ImportKeyPair(context.TODO(), input)

	if err != nil {
		return "", err
	}

	return aws.ToString(output.KeyPairId), nil
}
