package subArray

import (
	"fmt"
	"testing"
)

func TestZeroSumSub(t *testing.T) {
	result := ZeroSumSub([]int{1, 2, -1, -2, 3})
	if result == [2]int{0, 3} {
		t.Log("the result is true")
	}

}