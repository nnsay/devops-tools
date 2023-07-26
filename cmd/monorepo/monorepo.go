/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package monorepo

import (
	"fmt"

	"github.com/spf13/cobra"
)

// monorepoCmd represents the monorepo command
var MonorepoCmd = &cobra.Command{
	Use:   "monorepo",
	Short: "tools related to monorepo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please use the monorepo sub command")
	},
}
