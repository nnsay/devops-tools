/*
Copyright © 2023 Jimmy Wang <jimmy.w@aliyun.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// iamCmd represents the iam command
var iamCmd = &cobra.Command{
	Use:   "iam",
	Short: "tools related to iam",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("iam called")
	},
}

func init() {
	rootCmd.AddCommand(iamCmd)
}
