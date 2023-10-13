/*
Copyright © 2023 Jimmy Wang <jimmy.w@aliyun.com>
*/
package main

import (
	"os"
	"strconv"
	"time"

	"github.com/nnsay/devops-tools/cmd"
)

func main() {
	offset := 0
	offsetHours, isExist := os.LookupEnv("TIME_OFFSET_HOURS")
	if isExist {
		offset, _ = strconv.Atoi(offsetHours)
	}
	time.Local = time.FixedZone("zh-CN", offset*3600)
	cmd.Execute()
}
