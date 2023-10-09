package main

import (
    a1 "github.com/zeeshangondal/assignment01bca"
)

func main() {
    blockchain := &a1.Blockchain{}

    // Add some blocks
    blockchain.Blocks = append(blockchain.Blocks, a1.NewBlock("Zeeshan to Asif", 123, ""))
    blockchain.Blocks = append(blockchain.Blocks, a1.NewBlock("Awais to Manan", 456, blockchain.Blocks[0].CurrentHash))
    blockchain.Blocks = append(blockchain.Blocks, a1.NewBlock("Danish to Amanat", 789, blockchain.Blocks[1].CurrentHash))

    // Display all blocks
    blockchain.DisplayBlocks()

    // Change a block's transaction
    blockchain.ChangeBlock(1, "Updated transaction")

    // Verify the blockchain
    isValid := blockchain.VerifyChain()
    if isValid {
        println("Blockchain is valid.")
    } else {
        println("Blockchain is not valid.")
    }
}
