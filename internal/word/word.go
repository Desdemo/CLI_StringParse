package word

import (
	"strings"
	"unicode"
)

// 单词全部转换为大写
func ToUpper(s string) string  {
	return strings.ToUpper(s)
}
// 单子全部转换为小写
func ToLower(s string) string  {
	return strings.ToLower(s)
}

// 下划线转写为大写驼峰单词
func UnderscoreToUpperCamelCase(s string) string  {
	// 1、下划线替换为空格
	// 2、对应首字母大写
	// 3、去掉空格
	s = strings.Replace(s, "_"," ", -1)
	s =strings.Title(s)
	return strings.TrimSpace(s)
}

// 下划线转小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string  {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 驼峰单子转下划线单词
func CamelCaseToUnderscore(s string) string  {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}


