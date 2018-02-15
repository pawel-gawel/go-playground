package main

import "fmt"

type Block struct {
	parent *Block
	hash   string
	data   string
}

type BlockWrapper struct {
	*Block
}

// Structs run structs playground
func Structs() {
	println("\n\n---> Hello Structs!\n")

	newStruct()
	structComposition()
	structFunctions()
}

func newStruct() {
	var block1 = Block{nil, "block1", "data1"}
	fmt.Println("block1: ", block1)

	var block2 = &Block{hash: "block2"}
	fmt.Println("block2: ", *block2)

	var block3 = &Block{nil, "block3", "block3 data"}
	fmt.Println("block3: ", *block3)

	var block4 = new(Block)
	block4.hash = "block4"
	fmt.Println("block4: ", block4)

	var factoredBlock1 = structFactory(nil, "factored1", "factored1 data")
	fmt.Println("factoredBlock1: ", factoredBlock1)
}

func structFactory(parent *Block, hash, data string) *Block {
	return &Block{
		parent: parent,
		hash:   hash,
		data:   data,
	}
}

func structComposition() {
	var block1 = &Block{
		parent: &Block{hash: "parent1"},
	}
	println("block1 parent is: ", block1.parent.hash)

	var blockWrapper1 = BlockWrapper{Block: &Block{hash: "wrapped1"}}
	println("blockWrapper1 indirect hash is: ", blockWrapper1.hash)
}

func (b *Block) log() {
	println("block log: ", b.hash)
}

func structFunctions() {
	block1 := &Block{hash: "block1"}
	block1.log()
}
