package test

import (
	"study/pkg/sort"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	s := []int{9, 11, 77, 7, 8, 3, 2, 9}
	sort.InsertionSort(s)
	t.Log(s)
}
