package smt

import (
	"fmt"
	"strconv"
)

// This function calculates zero hash based how many layers are exists
// under current level
func getHashForEmptyNodes(currentDepth int, maxDepth int) *string {
	d := maxDepth - currentDepth
	h := ZeroLeaf
	for i := 0; i < d; i++ {
		h = hash(h, h)
	}
	return &h
}
func getBinaryAddr(index int, depth int) string {
	binaryString := strconv.FormatInt(int64(index), 2)
	return fmt.Sprintf("%0*s", depth, binaryString)
}

func getEdgePositionByAddress(key string, depth int) Position {
	if len(key) == 0 {
		return 0
	}
	r, _ := strconv.Atoi(key[depth : depth+1])
	if r == 0 {
		return Left
	}
	return Right
}
func hash(v1, v2 string) string {
	// used this format for convenient debugging instead of a real hash function
	return fmt.Sprintf(" [%s,%s] ", v1, v2)
}
