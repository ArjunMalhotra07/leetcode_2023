//! https://leetcode.com/problems/binary-tree-inorder-traversal/ 
//! https://leetcode.com/problems/merge-two-binary-trees/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	return returnInOrderTraversal(root, ans)
}
func returnInOrderTraversal(node *TreeNode, ans []int) []int {
	if node.Left != nil {
		ans = returnInOrderTraversal(node.Left, ans)
	}
	ans = append(ans, node.Val)
	if node.Right != nil {
		ans = returnInOrderTraversal(node.Right, ans)

	}
	return ans
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	return mergeTreesHelperFunc(root1, root2)
}
func mergeTreesHelperFunc(node1 *TreeNode, node2 *TreeNode) *TreeNode {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	node1.Val = node1.Val + node2.Val
	node1.Left = mergeTreesHelperFunc(node1.Left, node2.Left)
	node1.Right = mergeTreesHelperFunc(node1.Right, node2.Right)
	return node1

}
