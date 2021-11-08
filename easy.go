package first_day

import "fmt"

/*Given an integer array nums, return true if any value appears
at least twice in the array, and return false if every element is distinct*/
func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	for _, i := range nums {
		m[i] += 1
		if m[i] >= 2 {
			return true
		}
	}
	return false
}

func test_1() {
	ar1 := []int{1, 2, 3, 4}
	ar2 := []int{1, 2, 3, 1}
	v := [][]int{ar1, ar2}
	for _, i := range v {
		fmt.Println(i)
		fmt.Println(containsDuplicate(i))
	}
}

/*Merge two sorted linked lists and return it as a sorted list.
The list should be made by splicing together the nodes of the first two lists*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = new(ListNode)
	var p = head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	} else {
		p.Next = l2
	}
	return head.Next

}

/*Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times.
You may assume that the majority element always exists in the array*/

func majorityElement(nums []int) int {
	m := map[int]int{}
	for _, i := range nums {
		m[i] += 1
	}
	max_num := m[nums[0]]
	max_key := 0
	for g, j := range m {
		if j >= max_num {
			max_num = j
			max_key = g
		}
	}
	return max_key
}

func test_2() {
	ar1 := []int{3, 2, 3}
	ar2 := []int{2, 2, 1, 1, 1, 2, 2}
	v := [][]int{ar1, ar2}
	for _, i := range v {
		fmt.Println(i)
		fmt.Println(majorityElement(i))
	}
}

/*Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	switch {
	case root == nil:
		return 0
	case root.Left == nil:
		return 1 + minDepth(root.Right)
	case root.Right == nil:
		return 1 + minDepth(root.Left)
	default:
		l := minDepth(root.Left)
		r := minDepth(root.Right)
		switch {
		case l > r:
			return r + 1
		case l < r:
			return l + 1
		}
	}
	return 1
}