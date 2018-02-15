package main

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
	var block1 = Block{}
	println("block1 hash is: ", block1.hash)

	var block2 = &Block{hash: "block2"}
	println("block2 hash is: ", block2.hash)

	var block3 = &Block{nil, "block3", "block3 data"}
	println("block3 hash is: ", block3.hash)

	var block4 = new(Block)
	block4.hash = "block4"
	println("block4 hash is: ", block4.hash)

	var factoredBlock1 = structFactory(nil, "factored1", "factored1 data")
	println("factoredBlock1 is: ", factoredBlock1.hash)
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
