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
	"math"
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

func Test2(t *testing.T) {
	//fmt.Println(lengthOfLongestSubstring("qrsvbspk"))
	qt := &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 9,
				},
				Left: &TreeNode{
					Val: 8,
				},
			},
			Left: &TreeNode{
				Val:   4,
				Right: nil,
				Left:  nil,
			},
		},
		Left: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 8,
				},
				Left: &TreeNode{
					Val: 9,
				},
			},
			Left: &TreeNode{
				Val:   5,
				Right: nil,
				Left:  nil,
			},
		},
	}

	fmt.Printf("%+v", maxDepth(qt))

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

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := NewQueue()
	queue.push(root.Left)
	queue.push(root.Right)
	stack := make([]int, 0)
	for queue.size() != 0 {
		cq := NewQueue()
		for queue.size() != 0 {
			i := queue.pop().(*TreeNode)
			if i == nil {
				stack = append(stack, -1)
				continue
			}
			cq.push(i.Left)
			cq.push(i.Right)
			stack = append(stack, i.Val)
		}
		if !isa(stack) {
			return false
		}
		stack = make([]int, 0)
		queue = cq
	}
	return true
}

func isa(a []int) bool {
	if len(a)&1 != 0 {
		return false
	}
	n := len(a) - 1
	for i := len(a) / 2; i < len(a); i++ {
		if a[i] != a[n-i] {
			return false
		}
	}
	return true
}

type a struct {
	max int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cur := 0
	x := a{}
	dfs(root, &x, cur)
	return x.max
}

func dfs(root *TreeNode, max *a, cur int) {
	cur++
	if max.max < cur {
		max.max = cur
	}
	if root.Left != nil {
		dfs(root.Left, max, cur)
	}

	if root.Right != nil {
		dfs(root.Right, max, cur)
	}
}

func Test3(t *testing.T) {
	tree := sortedArrayToBST([]int{1, 2, 3, 4, 5, 6, 0, 8})
	fmt.Printf("%+v", tree)
	fmt.Printf("%+v", isBalanced(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: nil,
		},
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
			Left: &TreeNode{
				Val:   4,
				Right: nil,
				Left: &TreeNode{
					Val: 8,
				},
			},
		},
	}))

	fmt.Printf("%+v", lowestCommonAncestor(&TreeNode{
		Val: 6,
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
				},
				Left: &TreeNode{
					Val: 3,
				},
			},
			Left: &TreeNode{
				Val: 0,
			},
		},
	}, &TreeNode{Val: 2}, &TreeNode{Val: 4}))
}

/*
https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
*/
func sortedArrayToBST(nums []int) *TreeNode {
	return toBST(nums, 0, len(nums)-1)
}

func toBST(nums []int, start, end int) (tree *TreeNode) {
	if start > end {
		return
	}
	i := (end-start)/2 + start
	tree = &TreeNode{
		Val: nums[i],
	}
	tree.Left = toBST(nums, start, i-1)
	tree.Right = toBST(nums, i+1, end)
	return tree
}

/*
https://leetcode-cn.com/problems/balanced-binary-tree/
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	less := make([]int, 0)
	dfs1(&less, 0, root)
	min, max := less[0], less[0]
	for _, i := range less {
		if min > i {
			min = i
		}
		if max < i {
			max = i
		}
	}
	if max-min > 1 {
		return false
	}
	return true
}

func dfs1(less *[]int, d int, root *TreeNode) {
	d++
	if root.Left == nil {
		*less = append(*less, d)
	} else {
		dfs1(less, d, root.Left)
	}

	if root.Right == nil {
		*less = append(*less, d)
	} else {
		dfs1(less, d, root.Right)
	}
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	min := math.MaxInt32
	if root.Left != nil {
		min = fmin(minDepth(root.Left), min)
	}
	if root.Right != nil {
		min = fmin(minDepth(root.Right), min)
	}
	return min + 1
}

func fmin(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		if targetSum == 0 {
			return true
		}
		return false
	}
	a := false
	if root.Left != nil {
		a = hasPathSum(root.Left, targetSum-root.Val)
		if a {
			return a
		}
	}
	if root.Right != nil {
		a = hasPathSum(root.Right, targetSum-root.Val)
		if a {
			return a
		}
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			return true
		}
	}
	return false
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	p := &TreeNode{}
	invert(root, p)
	return p
}

func invert(root *TreeNode, p *TreeNode) {
	p.Val = root.Val
	if root.Left != nil {
		p.Right = &TreeNode{}
		invert(root.Left, p.Right)
	}
	if root.Right != nil {
		p.Left = &TreeNode{}
		invert(root.Right, p.Left)
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := NewQueue()
	queue.push(root)
	for queue.size() != 0 {
		a := queue.pop().(*TreeNode)
		m := make(map[int]int, 0)
		m[a.Val] = 1
		if a.Left != nil {
			m[a.Left.Val] = 1
			queue.push(a.Left)
		}
		if a.Right != nil {
			m[a.Right.Val] = 1
			queue.push(a.Right)
		}
		_, ok := m[p.Val]
		_, ok2 := m[q.Val]
		if ok && ok2 {
			return a
		}
	}
	return nil
}

func Test33(t *testing.T) {
	isHappy(100)
}

func isHappy(n int) bool {
	c := 0
	if n < 10 {
		c = n
	}
	a := make([]int, 0)
	for n >= 10 {
		a = append(a, n%10)
		n = n / 10
		if n < 10 {
			a = append(a, n)
		}
	}
	for _, cc := range a {
		c = cc*cc + c
	}
	if c < 10 {
		if c == 1 || c == 7 {
			return true
		} else {
			return false
		}
	} else {
		return isHappy(c)
	}
}
