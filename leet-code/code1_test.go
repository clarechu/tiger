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

package leet_code

import (
	"fmt"
	"strings"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("%+v", twoSum(nums, target))

}

/*
func TestSum(t *testing.T) {
	n := []int{9, 7, 11, 31, 78, 1, 7, 66, 5, 4, 82}
	quickSort1(n, 0, 10)
	fmt.Printf("%+v", n)
}
*/
func Test(t *testing.T) {
	//fmt.Println(lengthOfLongestSubstring("qrsvbspk"))
	qt := &TreeNode{
		Val:   1,
		Right: nil,
		Left: &TreeNode{
			Val: 2,
		},
	}
	pt := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
		},
	}
	fmt.Printf("%+v", isSameTree(qt, pt))

}

func findRepeatNumber(nums []int) int {
	m := make(map[int]int, 0)
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return m[n]
		}
		m[n] = 1
	}
	return 0
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	a := []byte(s)
	mlen := 1
	clen := 1
	max := []byte{}
	for k, i := range a {
		if k == len(a)-1 {
			if mlen <= 1 {
				mlen = clen
				clen = 1
				max = []byte{}
			}
		}
		max = append(max, i)
		for _, j := range a[k+1:] {
			if !strings.Contains(string(max), string(j)) {
				clen++
				max = append(max, j)
			} else {
				if mlen <= clen {
					mlen = clen
				}
				clen = 1
				max = []byte{}
				break
			}
		}

	}
	return mlen
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	ps := make([]int, 0)
	qs := make([]int, 0)
	ps = ltree(p, ps)
	qs = ltree(q, qs)
	if len(ps) != len(qs) {
		return false
	}
	for i := 0; i < len(ps); i++ {
		if ps[i] != qs[i] {
			return false
		}
	}
	return true
}

func ltree(p *TreeNode, a []int) []int {
	a = append(a, p.Val)
	if p.Left != nil {
		a = ltree(p.Left, a)
	} else {
		a = append(a, 0)
	}

	if p.Right != nil {
		a = ltree(p.Right, a)
	} else {
		a = append(a, 0)
	}
	return a
}

//isSymmetric https://leetcode-cn.com/problems/symmetric-tree/solution/dui-cheng-er-cha-shu-by-leetcode-solution/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetrics(root, root)
}

func isSymmetrics(q *TreeNode, p *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}
	if q == nil || p == nil {
		return false
	}
	if q.Val != p.Val {
		return false
	}
	return isSymmetrics(q.Right, p.Left) && isSymmetrics(q.Left, p.Right)
}