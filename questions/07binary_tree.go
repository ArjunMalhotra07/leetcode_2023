package questions

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ! https://leetcode.com/problems/binary-tree-inorder-traversal/
func InorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	return ReturnInOrderTraversal(root, ans)
}
func ReturnInOrderTraversal(node *TreeNode, ans []int) []int {
	if node.Left != nil {
		ans = ReturnInOrderTraversal(node.Left, ans)
	}
	ans = append(ans, node.Val)
	if node.Right != nil {
		ans = ReturnInOrderTraversal(node.Right, ans)

	}
	return ans
}

// ! https://leetcode.com/problems/merge-two-binary-trees/
func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	return MergeTreesHelperFunc(root1, root2)
}
func MergeTreesHelperFunc(node1 *TreeNode, node2 *TreeNode) *TreeNode {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	node1.Val = node1.Val + node2.Val
	node1.Left = MergeTreesHelperFunc(node1.Left, node2.Left)
	node1.Right = MergeTreesHelperFunc(node1.Right, node2.Right)
	return node1

}

// ! https://leetcode.com/problems/invert-binary-tree/description/
func InvertTree(root *TreeNode) *TreeNode {
	InvertTreeHelper(root)
	return root
}
func InvertTreeHelper(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		InvertTreeHelper(root.Left)
	}

	if root.Right != nil {
		InvertTreeHelper(root.Right)
	}
	var temp *TreeNode
	temp = root.Left
	root.Left = root.Right
	root.Right = temp
}
