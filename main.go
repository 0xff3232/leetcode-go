// https://leetcode.com/problems/binary-tree-inorder-traversal/submissions/1191279308/

package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
    return &TreeNode{
        Val: val,
        Left: nil,
        Right: nil, 
    }
}

// needed to see console output, also answers question.
func (bt *TreeNode) InorderPrint() {
    if bt == nil {
        return
    }
    bt.Left.InorderPrint()
    fmt.Println(bt.Val)
    bt.Right.InorderPrint()
}

func main() {
    /*
    example
            4
           / \
          2   7
         / \
        1   3
    */
    root := &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, 
            &TreeNode{3, nil, nil}}, &TreeNode{7, nil, nil}}  
    fmt.Println("Inorder traversal--")
    root.InorderPrint()
    fmt.Println("--")
}

/*
LeetCode answer:

func inorderTraversal(root *TreeNode) []int {    
    var result []int
    inorder(root, &result)
    return result
}

func inorder(root *TreeNode, arr *[]int) {
    if root == nil {
        return
    }
    inorder(root.Left, arr)
    *arr = append(*arr, root.Val)
    inorder(root.Right, arr)
}
*/