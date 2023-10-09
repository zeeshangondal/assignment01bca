// Zeeshan Ali
// 20i-0906

package main

import (
    a1 "github.com/zeeshangondal/assignment01bca"
)

func main() {
    blockchain := &a1.Blockchain{}

    // Add some blocks
    blockchain.Blocks = append(blockchain.Blocks, a1.NewBlock("Zeeshan to Asif", 5456, ""))
    blockchain.Blocks = append(blockchain.Blocks, a1.NewBlock("Awais to Manan", 90385, blockchain.Blocks[0].CurrentHash))
    blockchain.Blocks = append(blockchain.Blocks, a1.NewBlock("Danish to Amanat", 846923, blockchain.Blocks[1].CurrentHash))
    blockchain.Blocks = append(blockchain.Blocks, a1.NewBlock("Danish to Zeeshan", 8434523, blockchain.Blocks[1].CurrentHash))
    blockchain.Blocks = append(blockchain.Blocks, a1.NewBlock("Danish to Khizer", 84476473, blockchain.Blocks[1].CurrentHash))
    blockchain.Blocks = append(blockchain.Blocks, a1.NewBlock("Umer to Aneesa", 847462, blockchain.Blocks[1].CurrentHash))
    blockchain.Blocks = append(blockchain.Blocks, a1.NewBlock("Sara to Ibrahim", 23456, blockchain.Blocks[1].CurrentHash))
    blockchain.Blocks = append(blockchain.Blocks, a1.NewBlock("Zeeshan to Danish", 67579563, blockchain.Blocks[1].CurrentHash))

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
