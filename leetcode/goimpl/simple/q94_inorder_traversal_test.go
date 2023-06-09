package simple

import (
	"fmt"
	"goimpl/common"
	"testing"
)

func inorderTraversal(root *common.TreeNode) (nums []int) {
	inorder(root, &nums)
	return nums
}

func inorder(root *common.TreeNode, nums *[]int) []int {
	if root == nil {
		return *nums
	}

	inorder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inorder(root.Right, nums)

	return *nums
}

func Test_inorderTraversal(t *testing.T) {
	tree := &common.TreeNode{
		Val: 6,
		Left: &common.TreeNode{
			Val: 9,
			Left: &common.TreeNode{
				Val: 10,
			},
			Right: &common.TreeNode{
				Val: 7,
			},
		},
		Right: &common.TreeNode{
			Val: 4,
			Left: &common.TreeNode{
				Val: 5,
			},
			Right: &common.TreeNode{
				Val: 2,
				Left: &common.TreeNode{
					Val: 3,
				},
				Right: &common.TreeNode{
					Val: 1,
				},
			},
		},
	}

	nums := inorderTraversal(tree)
	fmt.Printf("%v\n", nums)
}
