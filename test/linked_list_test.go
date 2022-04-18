package test

import (
	datastructure "study/pkg/dataStructure"
	"testing"
)

func TestLinkedList(t *testing.T) {
	//初始化链表
	linkdlist := datastructure.NewLinkedList()
	//循环插入5个数字
	for i := 0; i < 5; i++ {
		linkdlist.InsertToTail(i + 1)
	}
	linkdlist.Print()
	//从链表头部插入
	linkdlist.InsertToHead(9)
	linkdlist.Print()
}
