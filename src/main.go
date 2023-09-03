package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

type MerkleTree struct {
	RootNode *MerkleNode
	Tree     [][]*MerkleNode
}
type MerkleNode struct {
	//Left  *MerkleNode
	//Right *MerkleNode
	Data string
}

func makeEven(nodes []*MerkleNode) []*MerkleNode {
	if len(nodes)%2 == 0 {
		return nodes
	}
	return append(nodes, nodes[len(nodes)-1:]...)
}

func GenerateMerkleTree(hashes []string) MerkleTree {
	if len(hashes) == 0 {
		panic("Hash array is empty")
	}
	var tree MerkleTree
	if len(hashes) == 1 {
		tree = MerkleTree{
			RootNode: &MerkleNode{Data: hashes[0]},
			Tree:     [][]*MerkleNode{{&MerkleNode{Data: hashes[0]}}},
		}
		return tree
	}
	bottomLevel := make([]*MerkleNode, len(hashes))
	for index, value := range hashes {
		bottomLevel[index] = &MerkleNode{Data: value}
	}
	bottomLevel = makeEven(bottomLevel)
	//TODO make the hashes even if it's not already

	tree.Tree = append(tree.Tree, bottomLevel)
	genTree(&tree, tree.Tree[0])
	return tree
}

func genTree(tree *MerkleTree, nodes []*MerkleNode) {
	length := len(nodes)

	var hashPairs []*MerkleNode
	for i := 0; i < length-1; i += 2 {
		//concat := append(nodes[i].Data, nodes[i+1].Data...)
		concat := nodes[i].Data + nodes[i+1].Data
		//h := hash(concat)
		h := concat
		hashPairs = append(hashPairs, &MerkleNode{Data: h})
	}
	if len(hashPairs) > 1 {
		hashPairs = makeEven(hashPairs)
	}
	tree.Tree = append(tree.Tree, hashPairs)
	if length == 1 {
		return
	}
	genTree(tree, hashPairs)
}

func (r MerkleTree) Init(size int) {
	for i := 0; i < size; i++ {

	}
}

func main() {
	var r []string
	for i := 0; i < 5; i++ {
		r = append(r, strconv.Itoa(i))
		//r[i] = []byte{byte(i)}
	}

	//r = append(r, []byte("hello world"))
	fmt.Println(r[len(r)-1:])
	mt := GenerateMerkleTree(r)
	fmt.Println(mt)
	for i, row := range mt.Tree {
		fmt.Printf("Row %d:\n", i)
		for j, node := range row {
			fmt.Printf("    Column %d: %v\n", j, *node)
		}
	}
	return
	var a [][]int
	a = [][]int{{1, 4}, {2, 5}}
	a = append(a, []int{1})
	a = append(a, []int{2})
	a[0] = append(a[0], 3)
	a[1] = append(a[1], 4)
	a[2] = append(a[2], 5, 6)
	fmt.Println(a)
	return
	// tree capacity is 2^n, if we want to put 6 leafs we need 2^3 (=8) tree
	// the depth is equal to n
	tree := MerkleTree{}
	treeMatrix := [3][]*MerkleNode{}
	treeMatrix[0][0] = tree.RootNode

	fmt.Println("hello")
	fmt.Println(hash([]byte(("hello"))))
	fmt.Println(hex.EncodeToString([]byte("3338be694f50c5f338814986cdf0686453a888b84f424d792af4b9202398f392")))
}

func hash(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	//return hex.EncodeToString(hasher.Sum(nil))
	return hasher.Sum(nil)
}

//func hashNodes(n1 []byte, n2 []byte) {
//	hasher := sha256.New()
//	hasher.Write(n1)
//	return hex.EncodeToString(hasher.Sum(nil))
//}
