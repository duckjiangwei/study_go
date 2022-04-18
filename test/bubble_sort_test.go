package test

import (
	sort "study/pkg/sort"
	"testing"
)

//冒泡排序测试
func TestBubbleSort(t *testing.T) {
	data := []int{2, 5, 11, 33, 9, 7, 44}
	sortData := sort.BubbleSort(data)
	t.Log(sortData)
}
