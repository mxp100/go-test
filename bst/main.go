/**
Library for working with Binary Search Tree

@author Yuriy Potemkin <ysoft2000@mail.ru>
@see https://ru.wikipedia.org/wiki/%D0%94%D0%B2%D0%BE%D0%B8%D1%87%D0%BD%D0%BE%D0%B5_%D0%B4%D0%B5%D1%80%D0%B5%D0%B2%D0%BE_%D0%BF%D0%BE%D0%B8%D1%81%D0%BA%D0%B0
*/
package bst

import (
	"../logger"
	"errors"
)

type any interface{}

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (tree *TreeNode) Insert(value int) error {

	if tree == nil {
		return errors.New("tree not init")
	}

	logger.JSON(map[string]any{
		"Type":      "BST",
		"Method":    "Insert",
		"Operation": "Compare",
		"Data": map[string]any{
			"Tree value": tree.Value,
			"Value":      value,
		},
	})
	if tree.Value == value {
		return nil
	}

	if tree.Value < value {
		if tree.Right == nil {
			logger.JSON(map[string]any{
				"Type":      "BST",
				"Method":    "Insert",
				"Operation": "Create right node",
				"Data": map[string]any{
					"Tree value": tree.Value,
					"Value":      value,
				},
			})
			tree.Right = &TreeNode{Value: value}
			return nil
		}
		logger.JSON(map[string]any{
			"Type":      "BST",
			"Method":    "Insert",
			"Operation": "Insert into right node",
			"Data": map[string]any{
				"Tree value": tree.Value,
				"Value":      value,
			},
		})
		return tree.Right.Insert(value)
	}

	if tree.Value > value {
		if tree.Left == nil {
			logger.JSON(map[string]any{
				"Type":      "BST",
				"Method":    "Insert",
				"Operation": "Create left node",
				"Data": map[string]any{
					"Tree value": tree.Value,
					"Value":      value,
				},
			})
			tree.Left = &TreeNode{Value: value}
			return nil
		}
		logger.JSON(map[string]any{
			"Type":      "BST",
			"Method":    "Insert",
			"Operation": "Insert into left node",
			"Data": map[string]any{
				"Tree value": tree.Value,
				"Value":      value,
			},
		})
		return tree.Left.Insert(value)
	}
	return nil
}

func (tree *TreeNode) Find(value int) (TreeNode, bool) {

	if tree == nil {
		logger.JSON(map[string]any{
			"Type":      "BST",
			"Method":    "Find",
			"Operation": "Check null",
			"Data":      map[string]any{},
		})
		return TreeNode{}, false
	}

	logger.JSON(map[string]any{
		"Type":      "BST",
		"Method":    "Find",
		"Operation": "Compare",
		"Data": map[string]any{
			"Tree value": tree.Value,
			"Value":      value,
		},
	})

	switch {
	case value == tree.Value:
		logger.JSON(map[string]any{
			"Type":      "BST",
			"Method":    "Find",
			"Operation": "Value founded",
			"Data": map[string]any{
				"Tree value": tree.Value,
				"Value":      value,
			},
		})
		return *tree, true
	case value < tree.Value:
		logger.JSON(map[string]any{
			"Type":      "BST",
			"Method":    "Find",
			"Operation": "Find in left",
			"Data": map[string]any{
				"Tree value": tree.Value,
				"Value":      value,
			},
		})
		return tree.Left.Find(value)
	default:
		logger.JSON(map[string]any{
			"Type":      "BST",
			"Method":    "Find",
			"Operation": "Find in right",
			"Data": map[string]any{
				"Tree value": tree.Value,
				"Value":      value,
			},
		})
		return tree.Right.Find(value)
	}
}

func (tree *TreeNode) Remove(value int) *TreeNode {
	logger.JSON(map[string]any{
		"Type":      "BST",
		"Method":    "Remove",
		"Operation": "Check null",
		"Data":      map[string]any{},
	})
	if tree == nil {
		return nil
	}

	logger.JSON(map[string]any{
		"Type":      "BST",
		"Method":    "Remove",
		"Operation": "Compare",
		"Data": map[string]any{
			"Tree value": tree.Value,
			"Value":      value,
		},
	})
	if value < tree.Value {
		logger.JSON(map[string]any{
			"Type":      "BST",
			"Method":    "Remove",
			"Operation": "Remove in left node",
			"Data": map[string]any{
				"Tree value": tree.Value,
				"Value":      value,
			},
		})
		tree.Left = tree.Left.Remove(value)
		return tree
	}
	if value > tree.Value {
		logger.JSON(map[string]any{
			"Type":      "BST",
			"Method":    "Remove",
			"Operation": "Remove in right node",
			"Data": map[string]any{
				"Tree value": tree.Value,
				"Value":      value,
			},
		})
		tree.Right = tree.Right.Remove(value)
		return tree
	}

	if tree.Left == nil && tree.Right == nil {
		logger.JSON(map[string]any{
			"Type":      "BST",
			"Method":    "Remove",
			"Operation": "Node end(empty left and right)",
			"Data": map[string]any{
				"Tree value": tree.Value,
				"Value":      value,
			},
		})
		tree = nil
		return nil
	}

	if tree.Left == nil {
		tree = tree.Right
		return tree
	}
	if tree.Right == nil {
		tree = tree.Left
		return tree
	}

	minRight := tree.Right
	for {
		if minRight != nil && minRight.Left != nil {
			minRight = minRight.Left
		} else {
			break
		}
	}

	logger.JSON(map[string]any{
		"Type":      "BST",
		"Method":    "Remove",
		"Operation": "Copy min right to current and remove right",
		"Data": map[string]any{
			"Tree value": tree.Value,
			"Value":      value,
		},
	})
	tree.Value = minRight.Value
	tree.Right = tree.Right.Remove(tree.Value)
	return tree
}
