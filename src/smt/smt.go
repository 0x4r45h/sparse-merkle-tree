package smt

import "math"

const ZeroLeaf = "0"

type Position int8

const (
	Left Position = iota
	Right
)

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  *string
}
type SparseMerkleTree struct {
	Root   *MerkleNode
	Leaves map[string]*string
	Depth  int
}

func NewTree(depth int) *SparseMerkleTree {
	root := getHashForEmptyNodes(0, depth)
	leaves := make(map[string]*string)
	return &SparseMerkleTree{Root: &MerkleNode{Data: root}, Depth: depth, Leaves: leaves}
}

func (smt SparseMerkleTree) Insert(index int, value *string) {
	//TODO make index big int to support trees bigger than 2^64
	if float64(index) > math.Pow(2, float64(smt.Depth)) {
		panic("index out of tree range")
	}
	key := getBinaryAddr(index, smt.Depth)
	smt.Leaves[key] = value
	smt.Root = smt.setNode(smt.Root, key, value, 0)
}
func (smt SparseMerkleTree) setNode(node *MerkleNode, key string, value *string, currentDepth int) *MerkleNode {
	maxDepth := smt.Depth
	if currentDepth == maxDepth {
		node.Data = value
		return node
	}
	edge := getEdgePositionByAddress(key, currentDepth)
	childNode := node.getOrCreateNode(currentDepth+1, maxDepth, edge)
	if edge == Left {
		node.Left = smt.setNode(childNode, key, value, currentDepth+1)
	} else {
		node.Right = smt.setNode(childNode, key, value, currentDepth+1)
	}
	node.Data = hashChilds(node.Left, node.Right, currentDepth, maxDepth)
	return node
}

func hashChilds(left *MerkleNode, right *MerkleNode, parentDepth int, maxDepth int) *string {
	var leftData *string
	var rightData *string
	childDepth := parentDepth + 1
	if left != nil {
		leftData = left.Data
	} else {
		leftData = getHashForEmptyNodes(childDepth, maxDepth)
	}
	if right != nil {
		rightData = right.Data
	} else {
		rightData = getHashForEmptyNodes(childDepth, maxDepth)
	}
	h := hash(*leftData, *rightData)
	return &h
}

func (node *MerkleNode) getOrCreateNode(depth, maxDepth int, position Position) *MerkleNode {
	if position == Left {
		if node.Left == nil {
			return &MerkleNode{Data: getHashForEmptyNodes(depth, maxDepth)}
		}
		return node.Left
	} else {
		if node.Right == nil {
			return &MerkleNode{Data: getHashForEmptyNodes(depth, maxDepth)}
		}
		return node.Right
	}
}
