package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBSTHelper(root, nil, nil)
}

func isBSTHelper(node *TreeNode, min *string, max *string) bool {
	if node == nil {
		return true
	}
	if (min != nil && node.Data <= *min) || (max != nil && node.Data >= *max) {
		return false
	}
	return isBSTHelper(node.Left, min, &node.Data) &&
		isBSTHelper(node.Right, &node.Data, max)
}
