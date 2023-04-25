/*
Copyright © 2023 Jimmy Wang <jimmy.w@aliyun.com>
*/
package iam

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/nnsay/aws-tools/lib"
	"github.com/spf13/cobra"
)

// checkExpirationCertificationCmd represents the checkExpirationCertification command
var checkExpirationCertificationCmd = &cobra.Command{
	Use:     "check-certification-date",
	Aliases: []string{"ccd"},
	Short:   "check the comming expiration certification",
	Run: func(cmd *cobra.Command, _ []string) {
		leftHours, _ := cmd.Flags().GetInt("expire-hours")
		pathPrefix, _ := cmd.Flags().GetString("path-prefix")
		channel, _ := cmd.Flags().GetString("channel")
		fmt.Printf("left hours: %d \n", leftHours)
		fmt.Printf("path prefix: %s \n", pathPrefix)
		fmt.Printf("channel: %s \n", channel)

		client := lib.GetIamClient()
		output, _ := client.ListServerCertificates(context.TODO(), &iam.ListServerCertificatesInput{PathPrefix: &pathPrefix})
		wg := sync.WaitGroup{}

		for _, cert := range output.ServerCertificateMetadataList {
			if time.Since(*cert.Expiration).Abs().Hours() < float64(leftHours) {
				wg.Add(1)
				go func(scm types.ServerCertificateMetadata) {
					defer wg.Done()
					title := fmt.Sprintf(":rotating_light: _%s_ 证书过期提醒 ", *scm.ServerCertificateName)
					envName, found := os.LookupEnv("ENV_NAME")
					if found {
						title += fmt.Sprintf("(*%s*)", envName)
					}
					message := fmt.Sprintf(":hourglass_flowing_sand: 证书将在 %s 过期, 请及时处理!", (*scm.Expiration).Format("2006-01-02 15:04:05"))
					lib.SendNotification(channel, title, message)
				}(cert)

			}
		}
		wg.Wait()
		fmt.Println("check done!")
	},
}

func init() {
	IamCmd.AddCommand(checkExpirationCertificationCmd)

	checkExpirationCertificationCmd.Flags().IntP("expire-hours", "e", 72, "the left expiration hours of the certification, default value is 72")
	checkExpirationCertificationCmd.Flags().StringP("path-prefix", "p", "/cloudfront/", "server certification path")
	checkExpirationCertificationCmd.Flags().StringP("channel", "c", "#devops", "slack chanel")
}
