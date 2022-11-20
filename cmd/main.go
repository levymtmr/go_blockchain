package main

import (
	"fmt"
	"strconv"

	"github.com/levymtmr/go_blockchain/internal"
)

func main() {
	chain := internal.InitBlockChain()

	chain.AddBlock("First block after genesis block")
	chain.AddBlock("Second block after genesis block")
	chain.AddBlock("Third block after genesis block")

	for _, block := range chain.Blocks {
		fmt.Printf("PrevHash: %x \n", block.PrevHash)
		fmt.Printf("Data: %s \n", block.Data)
		fmt.Printf("Hash: %x \n", block.Hash)

		pow := internal.NewProof(block)

		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
