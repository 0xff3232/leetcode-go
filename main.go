// https://leetcode.com/problems/delete-node-in-a-bst/description/

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

// Insert for BTS.
func (bt *TreeNode) Insert(val int) *TreeNode {
    if bt == nil {
        return NewTreeNode(val)
    }
    if val < bt.Val {
        bt.Left = bt.Left.Insert(val)
    } else {
        bt.Right = bt.Right.Insert(val)
    }
    return bt 
}

func (bt *TreeNode) FindMin() int {
    for bt.Left != nil {
        bt = bt.Left
    }
    return bt.Val
}

func (bt *TreeNode) Remove(val int) *TreeNode {
    if bt == nil {
        return nil
    }

    if val < bt.Val {
        bt.Left = bt.Left.Remove(val)
    } else if val > bt.Val {
        bt.Right = bt.Right.Remove(val)
    } else {
        if bt.Left == nil {
            return bt.Right
        } else if bt.Right == nil {
            return bt.Left
        } else {
            minNode := bt.FindMin()
            bt.Val = minNode
            bt.Right = bt.Right.Remove(minNode)
        }
    }

    return bt
}

// BTS print, meaning in order of Binary Search.
func (bt *TreeNode) Print() {
    if bt == nil {
        return
    }
    bt.Left.Print()
    fmt.Println(bt.Val)
    bt.Right.Print()
}

func main() {
    // creating the cases
    var root *TreeNode
    nums := []int{4,2,7,1,3}
    for i := range nums {
        root = root.Insert(nums[i])
    }

    n := 7
    
    fmt.Println("before removal:")
    root.Print()

    root.Remove(n)
    fmt.Println("after removal:")

    root.Print()
}

/*
LeetCode answer 

// TreeNode represents a node in a binary tree.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// deleteNode removes a node with the specified key from the BST.
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }

    // Search for the node to delete.
    if key > root.Val {
        // If key is greater, go right.
        root.Right = deleteNode(root.Right, key)
    } else if key < root.Val {
        // If key is lesser, go left.
        root.Left = deleteNode(root.Left, key)
    } else {
        // Found the node to delete.
        if root.Left == nil {
            // Node has no left child.
            return root.Right
        } else if root.Right == nil {
            // Node has no right child.
            return root.Left
        } else {
            // Node has two children, find the minimum node on the right subtree,
            // replace the current node's value with the min value,
            // then delete the minimum node in the right subtree.
            minVal := findMin(root.Right)
            root.Val = minVal
            root.Right = deleteNode(root.Right, minVal)
        }
    }
    return root
}

// findMin finds the minimum value in a subtree rooted at the given node.
func findMin(root *TreeNode) int {
    for root.Left != nil {
        root = root.Left
    }
    return root.Val
}

*/