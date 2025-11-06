/*
 Glock is a simple portmanteau that blends go and block,
 ergo implying a minimal blockchain implementation.

 This blockchain lives in the cli and is designed to be
 ephemeral, however you have some control over its lifespan.

 'glock init' begins it with genesis;
 (every operation in the meantime is saved to a JSON file)
 'glock reset' erases all blockchain data from JSON.

 Integrated and adapted blockchain code from a 2018 article.
 The original code sets up a Go server and a REST API for
 blockchain operations.

 Using flags/arguments, we can easily communicate with it
 in the same fashion, however in a much more convenient manner.

 Command-line components (and other helpers) by github.com/bxavaby

 MIT License © 2025 bxavaby

*/

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

const difficulty = 1

type Block struct {
	Index      int
	Timestamp  string
	BPM        int
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      string
}

var Blockchain []Block

type Message struct {
	BPM int
}

var mutex = &sync.Mutex{}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.BPM) + block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

func generateBlock(oldBlock Block, BPM int) Block {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Difficulty = difficulty

	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		newBlock.Nonce = hex
		if !isHashValid(calculateHash(newBlock), newBlock.Difficulty) {
			fmt.Println(calculateHash(newBlock), " do more work!")
			time.Sleep(time.Second)
			continue
		} else {
			fmt.Println(calculateHash(newBlock), " work done!")
			newBlock.Hash = calculateHash(newBlock)
			break
		}

	}
	return newBlock
}

// Function to print the entire blockchain
func printBlockchain() {
	for i, block := range Blockchain {
		fmt.Printf("\n:.:.:.:.:.:[ Block #%d ]:.:.:.:.:.:", i)
		fmt.Printf("\nIndex: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("BPM: %d\n", block.BPM)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Difficulty: %d\n", block.Difficulty)
		fmt.Printf("Nonce: %s\n", block.Nonce)
	}

	greatSuccess("\nBlockchain has been printed!")
}

// Function to calculate and print blockchain metrics
func showStats() {
	totalDifficulty := 0
	for _, block := range Blockchain {
		totalDifficulty += block.Difficulty
	}

	fmt.Println("\n:.:.:.:.:.:[ Blockchain Stats ]:.:.:.:.:.:")
	fmt.Printf("\nChain Length: %d blocks\n", len(Blockchain))
	fmt.Printf("Total Difficulty: %d\n", totalDifficulty)
	fmt.Printf("Average Difficulty: %.2f\n", float64(totalDifficulty)/float64(len(Blockchain)))
	fmt.Printf("Genesis Block Hash: %s\n", Blockchain[0].Hash)
	fmt.Printf("Latest Block Hash: %s\n\n", Blockchain[len(Blockchain)-1].Hash)

	greatSuccess("Stats have been printed!")
}

// Function to validate the blockchain
func validateBlockchain() {
	state := blockchainState()

	if Blockchain[0].Index != 0 {
		ohNoNoes("Genesis block has invalid index!", nil)
		return
	}

	for i := 1; i < len(Blockchain); i++ {
		if !isBlockValid(Blockchain[i], Blockchain[i-1]) {
			ohNoNoes(fmt.Sprintf("Block #%d is invalid!", i), nil)
			fmt.Printf("  Expected PrevHash: %s\n", Blockchain[i-1].Hash)
			fmt.Printf("  Actual PrevHash:   %s\n", Blockchain[i].PrevHash)
			return
		}
	}

	if state == "genesis-only" {
		greatSuccess("Genesis block is valid!")
	} else {
		greatSuccess(fmt.Sprintf("All %d blocks are valid.", len(Blockchain)))
	}
}
