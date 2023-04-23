/*
Copyright © 2023 Jimmy Wang <jimmy.w@aliyun.com>
*/
package iam

import (
	"fmt"

	"github.com/spf13/cobra"
)

// iamCmd represents the iam command
var IamCmd = &cobra.Command{
	Use:   "iam",
	Short: "tools related to iam",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Please use the iam sub command")
	},
}
