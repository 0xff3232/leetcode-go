// https://leetcode.com/problems/search-in-a-binary-search-tree/description/

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

// BTS print, meaning in order of Binary Search.
func (bt *TreeNode) Print() {
    if bt == nil {
        return
    }
    bt.Left.Print()
    fmt.Println(bt.Val)
    bt.Right.Print()
}

func (bt *TreeNode) Search(target int) *TreeNode{
    if bt == nil {
        return nil 
    }

    if (target < bt.Val) {
        return bt.Left.Search(target)
    } else if (target > bt.Val) {
        return bt.Right.Search(target)
    } 
    return bt
    
}

func main() {
    // creating the cases
    var root *TreeNode
    nums := []int{4, 2, 7, 1, 3}
    for i := range nums {
        root = root.Insert(nums[i])
    }
    root.Print()
    s := root.Search(7)
    if s != nil {
        fmt.Println("target found", s.Val) 
    } else {
        fmt.Println("target not found")
    }
}


/*
LeetCode struct and solution:

 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }


func searchBST(root *TreeNode, val int) *TreeNode {
   if root == nil {
       return nil
   } 

   if (val < root.Val) {
       return searchBST(root.Left, val)
   } else if (val > root.Val) {
       return searchBST(root.Right, val)
   } 
   return root
}
*/