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
	
	for i := length - 1; i >= 0; i-- {
		// 从字符串尾部开始遍历，分离出每一个单词（翻转）
		if string(source[i]) != " " {
			tempStr := tempStr + string(source[i])
			continue	
		}
		if string(source[i]) == " " && len(tempStr) != 0 {
			result = append(result, tempStr)
			if result != "" {
				result = append(result, " ")
			}
			
		}

		tempStr = ""
	}
	return result
}

func palindrome(source string, left, right int) (result string) {
	if left >= 0 && right < len(source) && source[left] == source[right] {
		left--
		right++
	}
	return source[left:right + 1]
}

func longestPalindrSub(source string) (result string) {
	if source == "" {
		return ""
	}
	var tempSub string
	for i, _ := range source {
		tempSub = palindrome(source, i, i)
		if len(tempSub) > len(result) {
			result = tempSub
		}
		tempSub = palindrome(source, i, i + 1)
		if len(tempSub) > len(result) {
			result = tempSub
		}
	}
	return result
}

func spaceRep(source string) (result string) {
	
}

func main() {

}
