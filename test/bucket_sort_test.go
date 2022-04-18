package test

import (
	sort "study/pkg/sort"
	"testing"
)

func TestBucketSort(t *testing.T) {
	s := []int{9, 11, 77, 7, 8, 3, 2, 9}
	sort.SelectionSort(s)
	t.Log(s)
}
