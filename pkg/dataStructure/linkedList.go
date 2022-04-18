package datastructure

import "fmt"

//节点
type linkedNode struct {
	value interface{} //当前节点的值
	next  *linkedNode //指向下个节点的指针
}

func (node *linkedNode) GetValue() interface{} {
	return node.value
}

//链表结构
type linkedList struct {
	length int
	head   *linkedNode
}

//初始化一个节点
func newLinkedNode(value interface{}) *linkedNode {
	return &linkedNode{value: value, next: nil}
}

//初始化一个链表
func NewLinkedList() *linkedList {
	return &linkedList{0, newLinkedNode(0)}
}

//尾节点后插入
func (l *linkedList) InsertToTail(value interface{}) {
	//找到尾节点
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	l.InsertAfter(cur, value)
}

//从头节点插入
func (l *linkedList) InsertToHead(value interface{}) {
	l.InsertAfter(l.head, value)
}

//从某节点后面插入
func (l *linkedList) InsertAfter(node *linkedNode, value interface{}) bool {
	if node == nil {
		return false
	}
	oldNext := node.next
	//要插入的节点放后面
	node.next = newLinkedNode(value)
	//再将旧节点加上
	node.next.next = oldNext
	l.length++
	return true
}

//从某节点前面插入
func (l *linkedList) InsertBefore(node *linkedNode, value interface{}) bool {
	//先找到节点
	cur := l.head.next
	per := l.head

	for cur != nil {
		if cur == node {
			break
		}
		per = cur
		cur = cur.next
	}
	//空链表，找不到那个节点
	if cur == nil {
		return false
	}
	//在per后面接上新节点
	per.next = newLinkedNode(value)
	//接着，新节点连上旧节点
	per.next.next = cur
	//长度自增
	l.length++
	return true
}

//删除某节点
func (l *linkedList) DelLinkedNode(node *linkedNode, value interface{}) bool {
	//先找到节点
	cur := l.head.next
	per := l.head

	for cur != nil {
		if cur == node {
			break
		}
		per = cur
		cur = cur.next
	}
	//空链表，找不到那个节点
	if cur == nil {
		return false
	}
	//开始删除
	per.next = cur.next
	l.length--
	return true
}

//链表转切片
func (l *linkedList) Print() {
	cur := l.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
