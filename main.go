package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	Hash     []byte // Hash of the block
	Data     []byte // Data in the block
	PrevHash []byte // Hash of the previous block
}

// Output: this info of block is hashed by sha256 and then converted to hexadecimal -> save to hash property
func (b *Block) DeriveHash() {
	// Create a buffer
	// []byte is a slice of bytes, this is b.Data and b.PrevHash
	// [][]byte is a slice of slices of bytes, this is [b.Data, b.PrevHash]
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// Create a hash of the buffer
	hash := sha256.Sum256(info)
	// Set the hash of the block to the hash of the buffer
	// [:] is a slice of bytes, this is hash[:]
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	// Create a new block
	// []byte{} -> pass parameter to contructor
	// []byte() -> typecast data to []byte
	block := &Block{[]byte{}, []byte(data), prevHash}
	// Derive the hash of the block
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		// print with %x to print in hexadecimal
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		// print with %s to print in string
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
