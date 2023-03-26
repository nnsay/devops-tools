/*
Copyright © 2023 Jimmy Wang <jimmy.w@aliyun.com>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/spf13/cobra"
)

// deleteExpredCertificationCmd represents the deleteExpredCertificationCmd command
var deleteExpredCertificationCmd = &cobra.Command{
	Use:     "delete-expired-certification",
	Aliases: []string{"dec"},
	Short:   "delete all expired server certification",
	Run: func(cmd *cobra.Command, args []string) {
		experationTime, _ := cmd.Flags().GetInt64("expiration")
		pathPrefix, _ := cmd.Flags().GetString("path-prefix")
		fmt.Printf("experation time: %s \n", time.Unix(experationTime, 0).Format("2006-01-02 15:04:05"))
		fmt.Printf("path prefix: %s \n", pathPrefix)
		cfg, _ := config.LoadDefaultConfig(context.TODO())
		client := iam.NewFromConfig(cfg)
		output, _ := client.ListServerCertificates(context.TODO(), &iam.ListServerCertificatesInput{PathPrefix: &pathPrefix})
		for _, cert := range output.ServerCertificateMetadataList {
			if cert.Expiration.Before(time.Unix(experationTime, 0)) {
				fmt.Printf("the certification(%s) expired at %s\n", *cert.ServerCertificateName, (*cert.Expiration).Format("2006-01-02 15:04:05"))
				client.DeleteServerCertificate(context.TODO(), &iam.DeleteServerCertificateInput{ServerCertificateName: cert.ServerCertificateName})
			}
		}
		fmt.Println("deletion done!")
	},
}

func init() {
	iamCmd.AddCommand(deleteExpredCertificationCmd)

	deleteExpredCertificationCmd.Flags().Int64P("expiration", "e", time.Now().Unix(), "expiration unix time, default value is now")
	deleteExpredCertificationCmd.Flags().StringP("path-prefix", "p", "/cloudfront/", "server certification path")
}
