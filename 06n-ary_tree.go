// ! https://leetcode.com/problems/n-ary-tree-preorder-traversal/
// ! https://leetcode.com/problems/n-ary-tree-postorder-traversal/
package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	ans = getPostOrderTraversalForNaryTree(root, ans)
	return ans
}
func getPostOrderTraversalForNaryTree(node *Node, ans []int) []int {
	if len(node.Children) != 0 {
		for i := 0; i < len(node.Children); i++ {
			ans = getPostOrderTraversalForNaryTree(node.Children[i], ans)
		}
	}
	if node != nil {
		ans = append(ans, node.Val)
	}
	return ans
}

func preorder(root *Node) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	ans = getPreOrderTraversalForNaryTree(root, ans)
	return ans
}
func getPreOrderTraversalForNaryTree(node *Node, ans []int) []int {
	if node != nil {
		ans = append(ans, node.Val)
	}
	if len(node.Children) != 0 {
		for i := 0; i < len(node.Children); i++ {
			ans = getPreOrderTraversalForNaryTree(node.Children[i], ans)
		}
	}
	return ans
}
