/*
Copyright © 2023 Jimmy Wang <jimmy.w@aliyun.com>
*/
package cmd

import (
	"os"

	"github.com/nnsay/devops-tools/cmd/cloudforamtion"
	"github.com/nnsay/devops-tools/cmd/iam"
	"github.com/nnsay/devops-tools/cmd/monorepo"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "devops-tools",
	Short: "a tool collection for devops",
	Long:  `a serial tools for usefual and common devops operation`,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.devops-tools.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(iam.IamCmd)
	rootCmd.AddCommand(cloudforamtion.CloudformationCmd)
	rootCmd.AddCommand(monorepo.MonorepoCmd)
}
