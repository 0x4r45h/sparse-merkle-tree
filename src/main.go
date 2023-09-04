package main

import (
	"fmt"
	smt "merkle-tree/src/smt"
)

func main() {
	sparseTree := smt.NewTree(3)
	fmt.Println(*sparseTree.Root.Data)
	v0 := "10"
	v1 := "25"
	v7 := "71"
	sparseTree.Insert(1, &v1)
	fmt.Println(*sparseTree.Root.Data)
	sparseTree.Insert(0, &v0)
	fmt.Println(*sparseTree.Root.Data)
	sparseTree.Insert(7, &v7)
	fmt.Println(*sparseTree.Root.Data)
}
