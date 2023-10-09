// Zeeshan Ali
// 20i-0906


package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Transaction   string // The transaction data will be stored in this block.
	Nonce         int    // Lets assume this is a random number which will be used for mining.
	PreviousHash  string // The hash of the previous block in the chain.
	CurrentHash   string // The hash of this block's data.
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.CurrentHash = block.CalculateHash()
	return block
}

func (b *Block) CalculateHash() string {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

type Blockchain struct {
	Blocks []*Block // A slice to store blocks in the blockchain.
}

func (bc *Blockchain) DisplayBlocks() {
	for _, block := range bc.Blocks {
		fmt.Println("Transaction:", block.Transaction)
		fmt.Println("Nonce:", block.Nonce)
		fmt.Println("Previous Hash:", block.PreviousHash)
		fmt.Println("Current Hash:", block.CurrentHash)
		fmt.Println()
	}
}

func (bc *Blockchain) ChangeBlock(blockIndex int, newTransaction string) {
	if blockIndex >= 0 && blockIndex < len(bc.Blocks) {
		bc.Blocks[blockIndex].Transaction = newTransaction
		bc.Blocks[blockIndex].CurrentHash = bc.Blocks[blockIndex].CalculateHash()
	}
}

func (bc *Blockchain) VerifyChain() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		if bc.Blocks[i].PreviousHash != bc.Blocks[i-1].CurrentHash {
			return false
		}
	}
	return true
}

func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	return fmt.Sprintf("%x", hash)
}
