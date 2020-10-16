package cmd

import (
	"CLI_StringParse/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strings"
	"time"
)

var calcluateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果: %s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calcluateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			if !strings.Contains(calcluateTime, " ") {
				layout = "2006-01-02"
			}
			currentTimer, err = time.Parse(layout, calcluateTime)
			if err != nil {

				t, _ := timer.GetCalculateTime(currentTimer, duration)
				currentTimer = time.Unix(int64(t), 0)
			}
		}

	},
}
