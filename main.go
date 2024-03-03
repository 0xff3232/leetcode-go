// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/solutions/1259453/simple-go-solution-using-recursion/

package main

import (
	"fmt"
    "slices"
)

// TreeNode represents a node in a binary tree.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// buildTree constructs a binary tree from preorder and inorder traversal arrays.
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }

    root := &TreeNode{Val: preorder[0]}
    mid := slices.Index(inorder, preorder[0]) // Use slices.Index for finding the index
    root.Left = buildTree(preorder[1:mid+1], inorder[:mid]) 
    root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
    return root
}

func main() {
    // Example usage
    preorder := []int{3, 9, 20, 15, 7}
    inorder := []int{9, 3, 15, 20, 7}
    root := buildTree(preorder, inorder)
    fmt.Println("Root of the constructed tree:", root.Val)
    // Output should be the root of the tree, for this example: 3
}
