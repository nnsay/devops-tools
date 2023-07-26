/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package monorepo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func formatPercent(v float64) string {
	if v < 65 {
		return fmt.Sprintf("🔴 %.2f%%", v)
	} else if v > 80 {
		return fmt.Sprintf("🟢 %.2f%%", v)
	} else {
		return fmt.Sprintf("🟡 %.2f%%", v)
	}
}

func FormatCoverageValue(v interface{}) string {
	if vv, isStr := v.(string); isStr && vv == "Unknown" {
		return `🔴 Unknown`
	} else if vv, isFloat64 := v.(float64); isFloat64 {
		return formatPercent(vv)
	} else {
		return formatPercent(float64(v.(int)))
	}
}

type ReportItemType struct {
	Total   float64     `json:"total"`
	Covered float64     `json:"covered"`
	Skipped float64     `json:"skipped"`
	Pct     interface{} `json:"pct"`
}
type ReportSummary struct {
	Total struct {
		Lines      ReportItemType `json:"lines"`
		Statements ReportItemType `json:"statements"`
		Functions  ReportItemType `json:"functions'`
		Branches   ReportItemType `json:"branches"`
	} `json:"total"`
}

// codeCoverageReportCmd represents the codeCoverageReport command
var codeCoverageReportCmd = &cobra.Command{
	Use:     "codeCoverageReport",
	Aliases: []string{"ccr"},
	Short:   "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		coverageDirs, _ := cmd.Flags().GetStringArray("coverageDir")
		reportPath, _ := cmd.Flags().GetString("reportPath")
		limitTarget, _ := cmd.Flags().GetInt("limitTarget")
		exitFilePath := fmt.Sprintf("%s.exit", reportPath)

		fmt.Printf("coverageDir: %v\n", coverageDirs)
		fmt.Printf("reportPath: %s\n", reportPath)
		fmt.Printf("exitFilePath: %s\n", exitFilePath)
		fmt.Printf("limitTarget: %d\n", limitTarget)

		var reportMarkdownString = []string{
			"# Coverage report",
			"|Project|Lines|Statements|Branches|Functions|",
			"|---|---|---|---|---|",
		}
		os.Remove(exitFilePath)
		for _, coverageDir := range coverageDirs {
			dirs, err := ioutil.ReadDir(coverageDir)
			if err != nil {
				panic(err)
			}
			for _, dir := range dirs {
				reportPath := fmt.Sprintf("%s/%s/coverage-summary.json", coverageDir, dir.Name())
				reportBytes, err := ioutil.ReadFile(reportPath)
				if err != nil {
					panic(err)
				}
				var reportSummary ReportSummary
				json.Unmarshal(reportBytes, &reportSummary)
				projectCoverage := fmt.Sprintf("|%s|%s|%s|%s|%s|",
					dir.Name(),
					FormatCoverageValue(reportSummary.Total.Lines.Pct),
					FormatCoverageValue(reportSummary.Total.Statements.Pct),
					FormatCoverageValue(reportSummary.Total.Functions.Pct),
					FormatCoverageValue(reportSummary.Total.Branches.Pct),
				)
				reportMarkdownString = append(reportMarkdownString, projectCoverage)
				ptcFloat, isFloat64 := reportSummary.Total.Statements.Pct.(float64)
				// ptcStr, isStr := reportSummary.Total.Statements.Pct.(string)
				if isFloat64 && ptcFloat < float64(limitTarget) /*|| (isStr && ptcStr == "Unknown")*/ {
					exitReasonMsg := fmt.Sprintf("%s statements coverage %v is lower than the limition target %d !\n", dir.Name(), reportSummary.Total.Statements.Pct, limitTarget)
					exitFile, _ := os.OpenFile(exitFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
					defer exitFile.Close()
					exitFile.WriteString(exitReasonMsg)
				}
			}
		}
		reportMarkdownFile, err := os.OpenFile(reportPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
		if err != nil {
			panic(err)
		}
		defer reportMarkdownFile.Close()
		reportMarkdownFile.WriteString(strings.Join(reportMarkdownString, "\n"))
	},
}

func init() {
	MonorepoCmd.AddCommand(codeCoverageReportCmd)

	codeCoverageReportCmd.Flags().StringArrayP("coverageDir", "d", []string{"tmp/packages"}, "coverage report folder")
	codeCoverageReportCmd.Flags().StringP("reportPath", "r", "tmp/report.md", "output code coverage report file path")
	codeCoverageReportCmd.Flags().IntP("limitTarget", "l", -1, "failure if statements code coverage less than this limit target value")
}
