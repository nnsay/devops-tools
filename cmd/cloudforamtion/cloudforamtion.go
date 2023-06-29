/*
Copyright © 2023 Jimmy Wang <jimmy.w@aliyun.com>
*/
package cloudforamtion

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cloudformationCmd represents the cloudforamtion command
var CloudformationCmd = &cobra.Command{
	Use:   "cloudformation",
	Short: "tools related to cloudformation",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Please use the cloudformation sub command")
	},
}
