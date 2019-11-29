package main

import (
	"fmt"
)

// input：字符串source
// output：翻转过的字符串
// function：将字符串source里空格分隔开的单词词序翻转
func reverseStr(source string) (result string) {
	if source == "" {
		return ""
	}
	length := len(source)
	var tempStr string
	for i := length - 1; i >= 0; i-- {
		// 从字符串尾部开始遍历，分离出每一个单词（翻转）
		if string(source[i]) != " " {
			tempStr = tempStr + string(source[i])
		}
		
	}

}

func main() {

}
