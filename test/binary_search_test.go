package test

import (
	sort "study/pkg/sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	s := []int{9, 11, 55, 2, 7, 33, 19}
	length := len(s) - 1
	position, err := sort.BinarySearch(s, 2, 0, length)
	if err == nil {
		t.Logf("位置是 %d", position)
	} else {
		t.Log("未找到")
	}
}
