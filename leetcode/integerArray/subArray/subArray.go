package subArray

import (
	"fmt"
)

func ZeroSumSub(source []int) (result [2]int) {
	if len(source) == 0 {
		return [2]int{}
	}
	tempSum := make(map[int]int)

	for i, _ := range source {
		sum := + source[i]
		if sum == 0 {
			return [2]int{0, i}
		}
		if tempSum[sum] != 0 {
			return [2]int{tempSum[sum], i}
		}

	}
	return [2]int{}
}

func testZeroSumSub()