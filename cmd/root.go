/*
Copyright © 2023 Jimmy Wang <jimmy.w@aliyun.com>
*/
package cmd

import (
	"os"

	"github.com/nnsay/aws-tools/cmd/cloudformation"
	"github.com/nnsay/aws-tools/cmd/iam"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aws-tools",
	Short: "a tool collection for aws devops",
	Long:  `a serial tools for usefual and common aws operation`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aws-tools.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(iam.IamCmd)
	rootCmd.AddCommand(cloudformation.CloudforamtionCmd)
}
