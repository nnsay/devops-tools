package lib

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func GetIamClient() *iam.Client {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	client := iam.NewFromConfig(cfg)
	return client
}

func GetCloudformationClient() *cloudformation.Client {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	client := cloudformation.NewFromConfig(cfg)
	return client
}
