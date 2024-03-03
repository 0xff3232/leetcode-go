// https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/

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

// needed to see console output.
func (bt *TreeNode) InorderPrint() {
    if bt == nil {
        return
    }
    bt.Left.InorderPrint()
    fmt.Println(bt.Val)
    bt.Right.InorderPrint()
}

func (bt *TreeNode) kthSmallest(k int) int {
    ans := -1

    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) { 
       if node == nil {
           return
       } 
        if node.Left != nil && k != 0 {
            dfs(node.Left)
        }
        k--
        if k == 0 {
            ans = node.Val
            return
        }
        if node.Right != nil && k != 0 {
            dfs(node.Right)
        } 
    }
    dfs(bt)
    return ans
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

    fmt.Println(root.kthSmallest(2))
    fmt.Println(root.kthSmallest(5))
    fmt.Println(root.kthSmallest(1))
}
