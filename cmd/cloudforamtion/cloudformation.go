/*
Copyright © 2023 Jimmy Wang <jimmy.w@aliyun.com>
*/
/*
Copyright © 2023 Jimmy Wang <jimmy.w@aliyun.com>
*/

package cloudforamtion

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CloudforamtionCmd represents the cloudformation command
var CloudforamtionCmd = &cobra.Command{
	Use:   "cloudforamtion",
	Short: "tools related to cloudforamtion",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Please use the cloudforamtion sub command")
	},
}
