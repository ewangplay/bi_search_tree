package bstree

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBiSearchTree(t *testing.T) {

	bt := BiSearchTree{}

	max := 50
	for i := 0; i < max; i++ {
		bt.Add(rand.Int63n(int64(max)))
	}

	if bt.IsEmpty() {
		t.Fatal("The binary search tree should be not empty.")
	}

	bt.InOrderTravel(func(v int64) {
		fmt.Printf("%v -> ", v)
	})
	fmt.Println("nil")

	fmt.Println("tree deepth: ", bt.GetDeepth())

	fmt.Println("min value: ", bt.GetMin())

	fmt.Println("max value: ", bt.GetMax())

	bt.Clear()
	if !bt.IsEmpty() {
		t.Fatal("The binary search tree should be empty.")
	}
}
