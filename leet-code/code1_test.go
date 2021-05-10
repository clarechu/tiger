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

func TestSumRootToLeaf(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	fmt.Println(sumRootToLeaf(root))
}

type X struct {
	x []int
}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	xx := &X{}
	sumRoot1(root, root.Val, xx)
	sum := 0
	for _, x := range xx.x {
		sum = x + sum
	}
	return sum
}

func sumRoot1(root *TreeNode, n int, xx *X) {
	if root.Right == nil && root.Left == nil {
		xx.x = append(xx.x, n)
	}
	if root.Right != nil {
		sumRoot1(root.Right, n<<1+root.Right.Val, xx)
	}
	if root.Left != nil {
		sumRoot1(root.Left, n<<1+root.Left.Val, xx)
	}
}

func TestTkthLargest(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}
	fmt.Println(kthLargest(root, 3))
}

type Result struct {
	val int
}

func kthLargest(root *TreeNode, k int) int {
	r := &Result{}
	n = k
	kt(root, r)
	return r.val
}

var n int

func kt(root *TreeNode, r *Result) {
	if root.Right != nil {
		kt(root.Right, r)
	}
	/*k--
	if k == 1 {
		r.val = root.Val
	}*/
	fmt.Println(root.Val)
	n--
	if n == 0 {
		r.val = root.Val
		return
	}
	if root.Left != nil {
		kt(root.Left, r)
	}
}

func TestLeafSimilar(t *testing.T) {
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	root2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(leafSimilar(root1, root2))
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	list1 = make([]int, 0)
	list2 = make([]int, 0)
	dfs12(root1)
	dfs13(root2)
	return thesame(list1, list2)
}

func thesame(list1, list2 []int) bool {

	if len(list1) != len(list2) {
		return false
	}
	for i, l := range list1 {
		if l != list2[i] {
			return false
		}
	}
	return true
}

var list1 []int
var list2 []int

func dfs12(root *TreeNode) {
	if root.Left == nil && root.Right == nil {
		list1 = append(list1, root.Val)
	}
	if root.Left != nil {
		dfs12(root.Left)
	}

	if root.Right != nil {
		dfs12(root.Right)
	}
}

func dfs13(root *TreeNode) {
	if root.Left == nil && root.Right == nil {
		list2 = append(list2, root.Val)
	}
	if root.Left != nil {
		dfs13(root.Left)
	}

	if root.Right != nil {
		dfs13(root.Right)
	}
}

func maxDepth12(root *Node) int {
	if root == nil {
		return 0
	}
	max := 0
	queue := &Queue{}
	queue.push(root)
	for queue.size() != 0 {
		max++
		cur := &Queue{}
		for queue.size() != 0 {
			node := queue.pop().(*Node)
			for _, c := range node.Children {
				cur.push(c)
			}
		}
		queue = cur
	}
	return max
}

func maxDepth13(root *Node) int {
	if root == nil {
		return 0
	} else if len(root.Children) == 0 {
		return 1
	} else {
		heights := make([]int, 0)
		for _, r := range root.Children {
			heights = append(heights, maxDepth13(r))
		}
		return maxSlice(heights) + 1
	}

}

func maxSlice(heights []int) int {
	max := 0
	for _, h := range heights {
		if h > max {
			max = h
		}
	}
	return max
}

var xx []int

func preorder(root *Node) []int {
	if root == nil {
		return make([]int, 0)
	}
	xx = make([]int, 0)
	preorder1(root)
	return xx
}

func preorder1(root *Node) {
	xx = append(xx, root.Val)
	for _, r := range root.Children {
		preorder1(r)
	}
}

func postorder(root *Node) []int {
	if root == nil {
		return make([]int, 0)
	}
	stack := &Stack{}
	x := make([]int, 0)
	stack.push(root)
	for stack.Size != 0 {
		node := stack.pop().(*Node)
		n := []int{node.Val}
		x = append(n, x...)
		for _, r := range node.Children {
			stack.push(r)
		}
	}
	return x
}

type RightTreeNode struct {
	Right    bool
	TreeNode *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root == nil {
		return sum
	}
	queue := &Queue{}
	tree := &RightTreeNode{
		TreeNode: root,
	}
	queue.push(tree)
	for queue.size() != 0 {
		node := queue.pop().(*RightTreeNode)
		if node.TreeNode.Left == nil && node.TreeNode.Right == nil {
			if node.Right {
				sum = sum + node.TreeNode.Val
			}
		}
		if node.TreeNode.Left != nil {
			node := &RightTreeNode{
				Right:    true,
				TreeNode: node.TreeNode.Left,
			}
			queue.push(node)
		}
		if node.TreeNode.Right != nil {
			node := &RightTreeNode{

				TreeNode: node.TreeNode.Right,
			}
			queue.push(node)
		}
	}
	return sum
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	root := &TreeNode{}
	merge2(root, root1, root2)
	return root
}

func merge2(root *TreeNode, root1 *TreeNode, root2 *TreeNode) {
	if root1 == nil {
		root.Val = root2.Val
		if root2.Left != nil {
			root.Left = root2.Left
		}
		if root2.Right != nil {
			root.Right = root2.Right
		}
	} else if root2 == nil {
		root.Val = root1.Val
		if root1.Left != nil {
			root.Left = root1.Left
		}
		if root1.Right != nil {
			root.Right = root1.Right
		}
	} else {
		root.Val = root1.Val + root2.Val
		if root1.Left != nil || root2.Left != nil {
			left := &TreeNode{}
			root.Left = left
			merge2(left, root1.Left, root2.Left)
		}
		if root1.Right != nil || root2.Right != nil {
			right := &TreeNode{}
			root.Right = right
			merge2(right, root1.Right, root2.Right)
		}
	}
}
