package questions

type Node struct {
	Val      int
	Children []*Node
}
// ! https://leetcode.com/problems/n-ary-tree-postorder-traversal/
func Postorder(root *Node) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	ans = GetPostOrderTraversalForNaryTree(root, ans)
	return ans
}
func GetPostOrderTraversalForNaryTree(node *Node, ans []int) []int {
	if len(node.Children) != 0 {
		for i := 0; i < len(node.Children); i++ {
			ans = GetPostOrderTraversalForNaryTree(node.Children[i], ans)
		}
	}
	if node != nil {
		ans = append(ans, node.Val)
	}
	return ans
}
// ! https://leetcode.com/problems/n-ary-tree-preorder-traversal/
func Preorder(root *Node) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	ans = GetPreOrderTraversalForNaryTree(root, ans)
	return ans
}
func GetPreOrderTraversalForNaryTree(node *Node, ans []int) []int {
	if node != nil {
		ans = append(ans, node.Val)
	}
	if len(node.Children) != 0 {
		for i := 0; i < len(node.Children); i++ {
			ans = GetPreOrderTraversalForNaryTree(node.Children[i], ans)
		}
	}
	return ans
}
