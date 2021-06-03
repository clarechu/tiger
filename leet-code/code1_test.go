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
	"sort"
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
	fmt.Printf("%+v", lowestCommonAncestor(&TreeNode{Val: 1,
		Left: &TreeNode{
			Val: 2,
		}, Right: &TreeNode{
			Val: 3,
		}}, &TreeNode{Val: 2}, &TreeNode{Val: 3}))
	/*	fmt.Printf("%+v", lowestCommonAncestor(&TreeNode{
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
	}, &TreeNode{Val: 2}, &TreeNode{Val: 4}))*/
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

var as []float64

func averageOfLevels(root *TreeNode) []float64 {
	as = make([]float64, 0)
	if root == nil {
		return as
	}
	queue := NewQueue()
	queue.push(root)
	for queue.size() != 0 {
		cur := NewQueue()
		size := queue.size()
		sum := 0
		for queue.size() != 0 {
			node := queue.pop().(*TreeNode)
			sum = sum + node.Val
			if node.Left != nil {
				cur.push(node.Left)
			}
			if node.Right != nil {
				cur.push(node.Right)
			}
		}
		s := float64(sum) / float64(size)
		as = append(as, s)
		queue = cur
	}
	return as
}

var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = make([]string, 0)
	if root == nil {
		return paths
	}
	if root.Left == nil && root.Right == nil {
		paths = append(paths, fmt.Sprintf("%d", root.Val))
	}
	if root.Right != nil {
		dfsPath(fmt.Sprintf("%d", root.Val), root.Right)
	}

	if root.Left != nil {
		dfsPath(fmt.Sprintf("%d", root.Val), root.Left)
	}

	return paths
}

func dfsPath(path string, root *TreeNode) {
	path = path + "->" + fmt.Sprintf("%d", root.Val)
	if root.Left != nil {
		dfsPath(path, root.Left)
	}
	if root.Right != nil {
		dfsPath(path, root.Right)
	}
	if root.Left == nil && root.Right == nil {
		paths = append(paths, path)
	}
}

func TestFindMode(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}
	fmt.Println(findMode(root))
}

var k []int

func findMode(root *TreeNode) []int {
	k = make([]int, 0)
	if root == nil {
		return k
	}
	dfs33(root)
	//最大个数
	max := 0
	//当前个数
	j := 0
	cur := 0
	kk := make([]int, 0)
	for _, i := range k {
		if max == 0 {
			kk = append(kk, i)
			cur = i
			j++
			max++
			continue
		}
		if cur == i {
			j++
		} else {
			j = 1
			cur = i
		}
		if max < j {
			kk = []int{i}
			max = j
		} else if max == j {
			kk = append(kk, i)
		}
	}
	return kk
}

func dfs33(root *TreeNode) {
	if root.Left != nil {
		dfs33(root.Left)
	}
	k = append(k, root.Val)
	if root.Right != nil {
		dfs33(root.Right)
	}
}

func isAnagram(s string, t string) bool {
	ss := []byte(s)
	tt := []byte(t)
	sort.Slice(ss, func(i, j int) bool {
		if ss[i] < ss[j] {
			return true
		}
		return false
	})

	sort.Slice(tt, func(i, j int) bool {
		if tt[i] < tt[j] {
			return true
		}
		return false
	})
	return string(ss) == string(tt)
}

func TestIsAnagxram(t *testing.T) {
	fmt.Println(isAnagram("anagram", "nagaram"))
}

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("raedvmtyxveocfyhluuozodpxlroyjcsfslqmjthsbxhteeinpmnejxxcsyjgugclkehagpemfrnqtrkiropblcqdboztxtsaxqnsktrhzelynbzkxcghhfntrdauyzhzgujhniazijshesigzckgxentepeohcltpydumougjkmgoscchqsryaiamoujnyfpcsbwqtwikedbsjxxtnrpfgeqymwfngixslwlifimdapgzanuqwhwpesaigeoiwoyxzjmxukbsvsjvnjhwdbqzuurfolcngefdpsewrpvwivrsjnttrewkytdvvguatidyemrswpdmeqjrxgfdmcdlrcgiqdkyaaykdqigcrldcmdfgxrjqemdpwsrmeyditaugvvdtykwerttnjsrviwvprwespdfegnclofruuzqbdwhjnvjsvsbkuxmjzxyowioegiasepwhwqunazgpadmifilwlsxignfwmyqegfprntxxjsbdekiwtqwbscpfynjuomaiayrsqhccsogmkjguomudyptlchoepetnexgkczgisehsjizainhjugzhzyuadrtnfhhgcxkzbnylezhrtksnqxastxtzobdqclbporikrtqnrfmepgaheklcgugjyscxxjenmpnieethxbshtjmqlsfscjyorlxpdozouulhyfcoevxytmvdear"))
}

func longestPalindrome(s string) string {
	ss := []byte(s)
	len := len(ss)
	if len < 2 {
		return s
	}
	max := 1
	begin := 0
	kk := make([][]bool, len)
	for jj := range ss {
		kk[jj][jj] = true
	}
	for l := 2; l <= len; l++ {
		for jj := 0; jj < len; jj++ {
			jjj := jj + l - 1
			if jjj >= len {
				break
			}

			if ss[jjj] != ss[jj] {
				kk[jj][jjj] = false
			} else {
				if jjj-jj <= 2 {
					kk[jj][jjj] = true
				} else {
					kk[jj][jjj] = kk[jj+1][jjj-1]
				}
			}
			//
			if jjj-jj+1 >= max && kk[jj][jjj] {
				max = jjj - jj + 1
				begin = jj
			}
		}
	}
	return string(ss[begin : max+begin])
}

func TestTrap(t *testing.T) {
	//fmt.Println(trap([]int{4,2,3}))
	fmt.Println(waysToStep(61))
}
func waysToStep(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	dp[2] = 4
	for i := 3; i < n; i++ {
		dp[i] = dp[i-1]%1000000007 + dp[i-2]%1000000007 + dp[i-3]%1000000007
		dp[i] = dp[i] % 1000000007
	}
	return dp[n-1]
}

func TestConstructor(t *testing.T) {
	fmt.Println(numDecodings("226"))
}

func numDecodings(s string) int {
	ss := []byte(s)
	if len(ss) == 1 {
		return 1
	}
	max := 0
	for i := 0; i < len(ss)-1; i++ {
		n := ss[i] - 48
		if n == 0 {
			continue
		}
		m := ss[i+1] - 48
		if m == 0 {
			max = max - 1
		}
		if n*10+m <= 26 {
			max = max + 2
		} else {
			max += 1
		}
	}
	return max
}

func TestWordBreak(t *testing.T) {
	fmt.Println(wordBreak("cars", []string{"car", "ca", "rs"}))
}

func wordBreak(s string, wordDict []string) bool {
	n := make([]string, 0)
	n = append(n, s)
	sort.Slice(wordDict, func(i, j int) bool {
		if len([]byte(wordDict[i])) < len([]byte(wordDict[j])) {
			return true
		}
		return false
	})
	for len(n) != 0 {
		for _, nn := range n {
			flag := false
			cur := make([]string, 0)
			for _, w := range wordDict {
				if strings.Contains(nn, w) {
					flag = true
					a := strings.Split(nn, w)
					for _, aa := range a {
						if aa != "" {
							cur = append(cur, aa)
						}
					}
					break
				}
			}
			if !flag {
				return false
			}
			n = cur
		}

	}

	return true
}

func TestNumArray_SumRange(t *testing.T) {
	fmt.Println(removeElement([]int{1, 1, 2}, 1))
}

func removeDuplicates(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i-1], nums[i:]...)
			i--
			continue
		}
	}
	return len(nums)
}

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			continue
		}
	}
	return len(nums)
}

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		n := digits[i] + 1
		if n == 10 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
				return digits
			}
		} else {
			digits[i] = n
			return digits
		}
	}
	return digits
}

var i []string

func readBinaryWatch(turnedOn int) []string {
	if turnedOn == 0 {
		return make([]string, 0)
	}
	i = make([]string, 0)
	read(turnedOn, 0, 0, 0, 1)
	return i
}

func read(turnedOn, h int, m int, cur int, begin int) {
	if h > 11 || m > 59 {
		return
	}
	if turnedOn == cur {
		mm := ""
		if m == 0 {
			mm = "00"
		} else if m < 10 {
			mm = fmt.Sprintf("0%d", m)
		} else {
			mm = fmt.Sprintf("%d", m)
		}
		i = append(i, fmt.Sprintf("%d:%s", h, mm))
		return
	}
	ma := map[int]int{
		1:  1,
		2:  2,
		3:  4,
		4:  8,
		5:  1,
		6:  2,
		7:  4,
		8:  8,
		9:  16,
		10: 32,
	}
	for i := begin; i <= 10; i++ {
		if i <= 4 {
			read(turnedOn, h+ma[i], m, cur+1, i+1)
		} else {
			read(turnedOn, h, m+ma[i], cur+1, i+1)
		}
	}
}

func TestReadBinaryWatch(t *testing.T) {
	fmt.Println(numberOfMatches(7))
}

/*
https://leetcode-cn.com/problems/count-of-matches-in-tournament/
*/
func numberOfMatches(n int) int {
	if n == 1 {
		return 0
	}
	c := make([]int, 0)
	i := 0
	for n >= 2 {
		if i == 0 {
			c = append(c, n/2)
		} else {
			c = append(c, c[i-1]+n/2)
		}
		i++
		if n%2 != 0 {
			n = n/2 + 1
		} else {
			n = n / 2
		}
	}
	return c[len(c)-1]
}

func TestSubsetXORSum(t *testing.T) {
	fmt.Println(subsetXORSum([]int{3, 4, 5, 6, 7, 8}))
	letterCombinations("6")
}

func subsetXORSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := 2
	nu := make([]int, 0)
	nu = append(nu, nums...)
	for len(nums)+1 > n {
		num2 = make([]int, 0)
		sub(nums, 0, 0, 1, n)
		nu = append(nu, num2...)
		n++
	}
	nn := 0
	for _, i := range nu {
		nn += i
	}
	return nn
}

var num2 []int

func sub(cur []int, begin, c, n, q int) {
	for j := begin; j < len(cur); j++ {
		b := c ^ cur[j]
		if n == q {
			num2 = append(num2, b)
		} else {
			sub(cur, j+1, b, n+1, q)
		}
	}
}

var m12 = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	cc = make([]string, 0)
	b := []byte(digits)
	letter(b, 0, "")
	return cc
}

var cc []string

func letter(b []byte, n int, cur string) {
	if len(b) == n {
		cc = append(cc, cur)
		return
	}
	mm := m12[b[n]]
	for i := 0; i < len(mm); i++ {
		cc := cur + mm[i]
		letter(b, n+1, cc)
	}
}

var combinationSums [][]int

func combinationSum(candidates []int, target int) [][]int {
	combinationSums = make([][]int, 0)
	sums(candidates, make([]int, 0), target, 0, 0)
	return combinationSums
}

func sums(candidates, cur []int, target, sum, i int) {
	if sum == target {
		combinationSums = append(combinationSums, append([]int(nil), cur...))
		return
	}
	if sum > target {
		return
	}
	for j := i; j < len(candidates); j++ {
		dd := append(cur, candidates[j])
		s := sum + candidates[j]
		sums(candidates, dd, target, s, j)
	}
}

func TestSubsetXORSum1(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

func permute(nums []int) [][]int {
	combinationSums = make([][]int, 0)
	sums1(nums, make([]int, 0))
	return combinationSums
}

func sums1(candidates, cur []int) {
	if len(candidates) == len(cur) {
		combinationSums = append(combinationSums, append([]int(nil), cur...))
	}
	for j := 0; j < len(candidates); j++ {
		flag := false
		for _, c := range cur {
			if c == candidates[j] {
				flag = true
			}
		}
		if !flag {
			dd := append(cur, candidates[j])
			sums1(candidates, dd)
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

var l *ListNode

func TestConstructor2(t *testing.T) {
	fmt.Printf("%+v", addTwoNumbers(&ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: &ListNode{},
			},
		},
	}, nil))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l = nil
	stack1 := NewStack()
	for l1 != nil {
		stack1.push(l1.Val)
		l1 = l1.Next
	}
	stack2 := NewStack()
	for l2 != nil {
		stack2.push(l2.Val)
		l2 = l2.Next
	}
	addStack(stack1, stack2, 0)
	return l
}

func addStack(stack1, stack2 *Stack, i int) {
	if stack1.Size > 0 && stack2.Size > 0 {
		v1 := stack1.pop().(int)
		v2 := stack2.pop().(int)
		l = &ListNode{
			Val:  (v1 + v2 + i) % 10,
			Next: l,
		}
		if v1+v2+i >= 10 {
			addStack(stack1, stack2, 1)
		} else {
			addStack(stack1, stack2, 0)
		}
	} else if stack1.Size > 0 {
		v1 := stack1.pop().(int)
		l = &ListNode{
			Val:  (v1 + i) % 10,
			Next: l,
		}
		if v1+i >= 10 {
			addStack(stack1, stack2, 1)
		} else {
			addStack(stack1, stack2, 0)
		}
	} else if stack2.Size > 0 {
		v2 := stack2.pop().(int) + i
		l = &ListNode{
			Val:  v2 % 10,
			Next: l,
		}
		if v2 >= 10 {
			addStack(stack1, stack2, 1)
		} else {
			addStack(stack1, stack2, 0)
		}
	} else if i != 0 {
		l = &ListNode{
			Val:  1,
			Next: l,
		}
	}
}

func TestNewQueue(t *testing.T) {
	fmt.Printf("%+v", lowestCommonAncestor(&TreeNode{
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
		}}, &TreeNode{
		Val: 9,
	}, nil))
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parent := map[int]*TreeNode{}
	dfs1111(root, parent)
	visited := map[int]bool{}
	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}

func dfs1111(r *TreeNode, parent map[int]*TreeNode) {
	if r == nil {
		return
	}
	if r.Left != nil {
		parent[r.Left.Val] = r
		dfs1111(r.Left, parent)
	}
	if r.Right != nil {
		parent[r.Right.Val] = r
		dfs1111(r.Right, parent)
	}
}

var ss []int

func findTarget(root *TreeNode, k int) bool {
	ss = make([]int, 0)
	toSlice(root)
	for k2, i := range ss {
		for k1, j := range ss {
			if (i+j) == k && k2 != k1 {
				return true
			}
		}
	}
	return false
}

func toSlice(root *TreeNode) {
	if root == nil {
		return
	}
	toSlice(root.Left)
	ss = append(ss, root.Val)
	toSlice(root.Right)
}

func TestConstructor3(t *testing.T) {
	fmt.Println(numTrees(3))
}

func numTrees(n int) int {
	a := make([]int, n+1)
	a[0] = 1
	a[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			a[i] += a[j-1] * a[i-j]
		}
	}
	return a[n]
}
