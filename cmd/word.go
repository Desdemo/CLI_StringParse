package cmd

import (
	"CLI_StringParse/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"

)

const (
	MODE_UPPER = iota + 1   // 单词全部大写
	MODE_LOWER  			// 单词全部小写
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE  // 下划线转为大写驼峰
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE // 下划线转为小写驼峰
	MODE_CAMELCASE_TO_UNDERSCORE // 驼峰转为下划线
)

var desc = strings.Join([]string{
	"该子命令支持全部单子格式转换，模式如下：",
	"1: 全部单词转为大写",
	"2: 全部单词转为小写",
	"3：下划线单词转为大写驼峰",
	"4: 下划线单词转为小写驼峰",
	"5: 驼峰单词转为下划线单词",
}, "\n")

var mode int8
var str string
var wordCmd = &cobra.Command{
	Use: "word",
	Short: "单词格式转换",
	Long: desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)
		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = word.UnderscoreToUpperCamelCase(str)
		case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
			content = word.UnderscoreToLowerCamelCase(str)
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该格式，请执行 help word 查看帮助文档")
		}
		log.Printf("输出结果： %s", content)
	},
}

func init()  {
	wordCmd.Flags().StringVarP(&str, "str","s","","请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m",0,"请输入单词转换模式")
}
