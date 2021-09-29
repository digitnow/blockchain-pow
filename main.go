package main

import (
    "fmt"
    "bytes"
    "crypto/sha256"
)

// version01: implementing av simple linked list model to simulate a blockchain
type BlockChain struct {
    blocks []*Block
}

type Block struct {
    Hash        []byte // Derived from Data and PrevHash
    Data        []byte // Can be passed around in a dist db
    PrevHash    []byte // Last block's hash; links blocks together
}

func (b *Block) DeriveHash() {
    info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
    hash := sha256.Sum256(info)
    b.Hash = hash[:]
}

// sha256 is simple compared to the real implementations
func CreateBlock(data string, prevHash []byte) *Block {
    block := &Block{[]byte{}, []byte(data), prevHash}
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
    chain.AddBlock("Block #1")
    chain.AddBlock("Block #2")
    chain.AddBlock("Block #3")
    for _, block := range chain.Blocks {
        fmt.Printf("------------------------------------\n")
        fmt.Printf("Previous Hash: %x\n", block.PrevHash)
        fmt.Printf("Data in Block: %s\n", block.Data)
        fmt.Printf("Hash: %x\n", block.Hash)        
        
        fmt.Println()
    }
}
