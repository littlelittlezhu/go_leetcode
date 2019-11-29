package main

import (
	"fmt"
)

// input：字符串source，target
// output：-1/1
// function：判断target是否是source的子串
func subStr(source, target string) int {
	if source == "" || target == "" || len(source) < len(target){ //此处也可以使用len函数判空，但是查资料显示两种方式代价无明显差别
		return -1
	}
	num := len(target)
	temp := 0
	//此处转化为byte数组对于中文字符可能会出问题，因此最好应该转为rune数组
	sources := []rune(source)
	targets := []rune(target)
	// 扫描一遍即可，类似kmp算法，还可优化使循环上界为source长度减去target长度
	for _, sub := range  sources {
		if sub == targets[temp] {
			temp++
		}else {
			temp = 0
		}
		if num == temp {
			return 1
		}
	}
	return -1
}

// input：字符串source，target
// output：-1/1
// function：判断target在source的最长子串（注意target子串都要匹配一次）长度
func longestSubStr(source, target string) int {
	//此处也可以使用len函数判空，但是查资料显示两种方式代价无明显差别
	if source == "" || target == "" {
		return 0
	}

	var longest, tempL int

	sources := []rune(source)
	targets := []rune(target)
	// 对于target的每个字串都要进行一次kmp，注意此处与subStr的区别
	for i, _ := range  sources {
		for _, subT := range  targets {
			//每匹配一次，下一次将和source的下一个字符进行匹配
			if sources[i + tempL] == subT {
				tempL++
			}else {
				tempL = 0
			}
			if longest < tempL {
				longest = tempL
			}
		}

	}
	return longest
}

// input：字符串source，target
// output：false/true
// function：判断target和source是否互为变位词(修改一定条件后也可以变成compareStr函数)
// 此处判断使用了map哈希表结构，另外也可以直接使用一个字符串长度大小的数组，分别统计个字符数量
func anaGram(source, target string) bool {
	//此处也可以使用len函数判空，但是查资料显示两种方式代价无明显差别
	if source == "" || target == "" || len(source) != len(target){
		return false
	}
	//此处使用byte到int的映射对于编码大于8位的字符可能会有隐患，因此推荐使用更通用的rune类型
	//先将字符串转化为rune数组
	sources := []rune(source)
	targets := []rune(target)
	mapR := make(map[rune]int)
	mapT := make(map[rune]int)
	for i,_ := range source {
		mapR[sources[i]]++
		mapT[targets[i]]++
	}

	for j,_ := range mapR {
		//fmt.Println(j,map1[j])
		if mapR[j] != mapT[j] {
			return false
		}
	}
	return true
}

// input：字符串source，target
// output：false/true
// function：判断target里所有字符在source是否都出现
func cmpStr(source, target string) bool {
	//此处也可以使用len函数判空，但是查资料显示两种方式代价无明显差别
	if source == "" || target == "" || len(source) < len(target){
		return false
	}
	//此处使用byte到int的映射对于编码大于8位的字符可能会有隐患，因此推荐使用更通用的rune类型
	//先将字符串转化为rune数组
	sources := []rune(source)
	targets := []rune(target)
	mapR := make(map[rune]int)
	mapT := make(map[rune]int)

	for i,_ := range source {
		mapR[sources[i]]++
	}
	for j,_ := range target {
		mapT[targets[j]]++
	}

	for k,_ := range mapT {
		//fmt.Println(j,map1[j])
		if mapR[k] != mapT[k] {
			return false
		}
	}
	return true
}

// input：字符串数组source，target
// output：anaGram字符串数组
// function：将字符串数组source里出现的变位词都分成一组并输出
func anaGrams(source []string) (result []string) {
	// 此处赋值并非拷贝赋值，而是引用，因此之前的想法拷贝一份source先行处理是错误的
	temps := source
	mapR := make(map[string][]int)

	// 统计排序后的、互为变位词的字符串的数组下标集合
	for i,temp := range temps {
		mapR[strSort(temp)] = append(mapR[strSort(temp)],i)
	}

	//对于数组长度大于等于2的统计并作为结果输出
	for _,temp := range mapR {
		if len(temp) > 1 {
			for _,index := range temp {
				// 在go语言里要注意append函数的返回值一定要有变量去承接
				result = append(result, source[index])
			}
		}
	}
	return result
}

// input：字符串source，int型的offset
// output：翻转offset位的字符串
// function：将字符串翻转offset位
func rotate(source string, offset int) (result string) {
	sourceSub1 := reverse(source[:len(source) - offset])
	sourceSub2 := reverse(source[len(source) - offset:])
	sourceSub := sourceSub1 + sourceSub2
	result = reverse(sourceSub)
	return result
}

// input：字符串source
// output：翻转后的字符串
// function：将字符串翻转
func reverse(source string) (result string) {
	if source == "" {
		return ""
	}
	sources := []rune(source)
	length := len(sources)
	for i := 0; i < length / 2; i++ {
		sources[i], sources[length - (i + 1)] = sources[length - (i + 1)], sources[i]
	}
	return string(sources)
}
// input：字符串source
// output：有序字符串
// function：将无序字符串按ASCII码从小到大排序
func strSort(source string) string {
	//使用rune数组更加通用
	sources := []rune(source)

	for i,_ := range sources {
		for j := i+1; j < len(sources); j++ {
			if sources[i] > sources[j] {
				sources[i], sources[j] = sources[j], sources[i]
			}
		}
	}
	return string(sources)
}

func main() {
	//subResult := longestSubStr("asdfgh", "afga")
	//fmt.Println(subResult)
	//anaResult := cmpStr("中国ren", "国中h")
	//fmt.Println(anaResult)
	//source := []string{"lint", "intl", "inlt", "code"}
	//result := anaGrams(source)
	//fmt.Println(result)
	fmt.Println("asdfg")
	result := rotate("asdfg", 2)
	fmt.Println(result)
}