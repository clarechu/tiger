/*
Copyright 2020 The go-harbor Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
*/

package main

import "fmt"

func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[start]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}

/**
* https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	n := &ListNode{}
	cur := n
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			cur.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		} else {
			cur.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1

	} else {
		cur.Next = l2
	}
	return n.Next
}

func addDigits(num int) int {
	i := len(fmt.Sprintf("%d", num))
	cur := 0
	for i != 1 {
		for num > 10 {
			cur = num%10 + cur
			num = num / 10
		}
		num = cur + num
		i = len(fmt.Sprintf("%d", cur))
		cur = 0
	}
	return cur
}

func main() {
	n := []int{9, 7, 11, 31, 78, 1, 7, 66, 5, 4, 82, 100, 33, 22, 73, 234, 8}
	quickSort3(n, 0, 16)
	fmt.Printf("%+v", n)
}

/*func singleNumber(nums []int) []int {
	m := make(map[int]int, 0)
	for _, n := range nums {
		k, ok := m[n]
		if ok {
			delete(m, n)
		} else {
			m[n] = 1
		}
	}
	a := make([]int,0)
	for xx, _ := range m {
		a = append(a, xx)
	}
	return a
}*/

func quickSort1(a []int, start, end int) {
	base := a[start]
	i := start
	j := end
	for i < j {
		for a[j] > base {
			j--
		}
		for a[i] <= base {
			if i == j {
				exchange(a, start, j)
				i++
				j--
				break
			} else {
				i++
			}
		}
		if i < j {
			exchange(a, i, j)
		}
	}
	if start < i {
		quickSort1(a, start, j)
	}
	if j < end {
		quickSort1(a, i, end)
	}
}

func exchange(a []int, i, j int) {
	/*	cc := a[i]
		a[i] = a[j]
		a[j] = cc*/
	a[i], a[j] = a[j], a[i]
}

func quickSort3(a []int, start, end int)  {
	base := a[start]
	i, j := start, end
	for i < j {
		for base < a[j] {
			j--
		}
		for base >= a[i] {
			if i == j {
				a[start], a[j] = a[j], a[start]
				j--
				i++
			} else {
				i++
			}
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}

	if start < j {
		quickSort3(a, start, j)
	}
	if i < end {
		quickSort3(a, i, end)
	}

}