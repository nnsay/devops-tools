/*
Copyright © 2023 Jimmy Wang <jimmy.w@aliyun.com>
*/
package cloudforamtion

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/nnsay/aws-tools/lib"
	"github.com/spf13/cobra"
)

// checkExpirationCloudformationCmd represents the checkExpirationCloudformation command
var checkExpirationCloudformationCmd = &cobra.Command{
	Use:     "checkExpirationCloudformation",
	Aliases: []string{"cec"},
	Short:   "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		days, _ := cmd.Flags().GetInt("days")
		channel, _ := cmd.Flags().GetString("channel")
		fmt.Printf("days: %d\n", days)
		fmt.Printf("channel: %s\n", channel)

		client := lib.GetCloudformationClient()
		output, _ := client.ListStacks(context.TODO(), &cloudformation.ListStacksInput{
			StackStatusFilter: []types.StackStatus{
				types.StackStatusCreateFailed,
				types.StackStatusCreateComplete,

				types.StackStatusDeleteFailed,

				types.StackStatusUpdateFailed,
				types.StackStatusUpdateComplete,

				types.StackStatusUpdateRollbackFailed,
				types.StackStatusUpdateRollbackComplete,

				types.StackStatusRollbackFailed,
				types.StackStatusRollbackComplete,
			},
		})

		whiteStackNames, emptyWhiteStackName := os.LookupEnv("WHITE_STACK_NAMES")
		if emptyWhiteStackName {
			fmt.Println("WHITE_STACK_NAMES is empty")
		}
		for _, summarie := range output.StackSummaries {
			lastUpdatedTime := *summarie.LastUpdatedTime
			description := *summarie.TemplateDescription
			noChangeDays := (int)(time.Since(lastUpdatedTime).Abs().Hours() / 24)
			name := *summarie.StackName
			if strings.Contains(whiteStackNames, name) || noChangeDays < days {
				continue
			}
			title := ":mag: Cloudformation提醒 "
			envName, found := os.LookupEnv("ENV_NAME")
			if found {
				title += fmt.Sprintf("(*%s*)", envName)
			}
			message := fmt.Sprintf("名称: _%s_ \n", name)
			message += fmt.Sprintf("描述: _%s_ \n", description)
			message += fmt.Sprintf("该资源自从 _%s_ 未更新, 请考虑处理!", lastUpdatedTime.Format("2006-01-02 15:04:05"))
			lib.SendNotification(channel, title, message)
		}
	},
}

func init() {
	CloudformationCmd.AddCommand(checkExpirationCloudformationCmd)

	checkExpirationCloudformationCmd.Flags().IntP("days", "d", 10, "the max no change days, default value is 10")
	checkExpirationCloudformationCmd.Flags().StringP("channel", "c", "#devops", "slack chanel")
}
