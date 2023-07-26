/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package monorepo_test

import (
	"testing"

	"github.com/nnsay/devops-tools/cmd/monorepo"
)

func TestFormatCoverageValue(t *testing.T) {
	result := monorepo.FormatCoverageValue("Unknown")
	t.Logf("unknown: %v", result)
	result = monorepo.FormatCoverageValue(90.0)
	t.Logf("green: %v", result)
	result = monorepo.FormatCoverageValue(30.13)
	t.Logf("red: %v", result)
	result = monorepo.FormatCoverageValue(75)
	t.Logf("yellow: %v", result)
}
