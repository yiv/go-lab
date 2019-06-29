package main

import "fmt"

type ListNode struct {
	Key  int32
	Next *ListNode
}

func GetLastListNodes(head *ListNode, k int32) *ListNode {
	var theNode *ListNode
	var all []*ListNode
	pos := head
	for {
		if pos == nil {
			break
		}
		all = append(all, pos)
		pos = pos.Next
	}
	theNode = all[k]
	return theNode
}

func main() {
	var head *ListNode
	for i := 50; i > 0; i-- {
		newNode := &ListNode{Key: int32(i), Next: nil}
		if head == nil {
			head = newNode
		} else {
			newNode.Next = head
			head = newNode
		}
	}
	fmt.Println("有 50 个节点的链表，从头到尾它们的 Key 依次是 49,48,47......1,0")
	node := GetLastListNodes(head, 10)
	fmt.Println("获取它倒数第10个节点的 Key 值为：", node.Key)
}
