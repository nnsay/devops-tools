package lib

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func GetIamClient() *iam.Client {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	client := iam.NewFromConfig(cfg)
	return client
}
