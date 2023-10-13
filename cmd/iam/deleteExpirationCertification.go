/*
Copyright © 2023 Jimmy Wang <jimmy.w@aliyun.com>
*/
package iam

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/nnsay/devops-tools/lib"
	"github.com/spf13/cobra"
)

// deleteExpredCertificationCmd represents the deleteExpredCertificationCmd command
var deleteExpredCertificationCmd = &cobra.Command{
	Use:     "delete-expired-certification",
	Aliases: []string{"dec"},
	Short:   "delete all expired server certification",
	Run: func(cmd *cobra.Command, _ []string) {
		experationTime, _ := cmd.Flags().GetInt64("expiration")
		pathPrefix, _ := cmd.Flags().GetString("path-prefix")
		fmt.Printf("experation time: %s \n", time.Unix(experationTime, 0).In(time.Local).Format("2006-01-02 15:04:05"))
		fmt.Printf("path prefix: %s \n", pathPrefix)

		client := lib.GetIamClient()

		output, _ := client.ListServerCertificates(context.TODO(), &iam.ListServerCertificatesInput{PathPrefix: &pathPrefix})
		var wg sync.WaitGroup
		for _, cert := range output.ServerCertificateMetadataList {
			if cert.Expiration.Before(time.Unix(experationTime, 0)) {
				wg.Add(1)
				go func(scm types.ServerCertificateMetadata) {
					defer wg.Done()
					fmt.Printf("the certification(%s) expired at %s\n", *scm.ServerCertificateName, (*scm.Expiration).In(time.Local).Format("2006-01-02 15:04:05"))
					client.DeleteServerCertificate(context.TODO(), &iam.DeleteServerCertificateInput{ServerCertificateName: scm.ServerCertificateName})
				}(cert)
			}
		}
		wg.Wait()
		fmt.Println("deletion done!")
	},
}

func init() {
	IamCmd.AddCommand(deleteExpredCertificationCmd)

	deleteExpredCertificationCmd.Flags().Int64P("expiration", "e", time.Now().Unix(), "expiration unix time, default value is now")
	deleteExpredCertificationCmd.Flags().StringP("path-prefix", "p", "/cloudfront/", "server certification path")
}
