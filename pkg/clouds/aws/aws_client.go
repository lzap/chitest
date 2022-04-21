package aws

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var Client *ec2.Client

func Initialize() {
	key := os.Getenv("AWS_KEY")
	secret := os.Getenv("AWS_SECRET")
	session := os.Getenv("AWS_SESSION")
	options := ec2.Options{
		Region:      os.Getenv("AWS_REGION"),
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(key, secret, session)),
	}
	Client = ec2.New(options)
}
