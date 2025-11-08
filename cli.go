// glock/cli.go

package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func Logo() string {
	logo := `

        __              __
.-----.|  |.-----.----.|  |--.
|  _  ||  ||  _  |  __||    <
|___  ||__||_____|____||__|__|
|_____| v0.1.0 ==============+

:::::::::::::::::::::::::::::::
      >_ ARR bxavaby 2025     +
:::::::::::::::::::::::::::::::

+-----------------------------+
|       ̿'\̵͇̿̿\з=(◕_◕)=ε/̵͇̿̿/'̿'̿      |
|      go blockchain cli      |
+-----------------------------+

`
	return logo
}

func Verlogo() string {
	logo := `

        __              __
.-----.|  |.-----.----.|  |--.
|  _  ||  ||  _  |  __||    <
|___  ||__||_____|____||__|__|
|_____| v0.1.0 ==============+

Run 'glock init' to begin genesis.
Run 'glock help' to see all commands.

`
	return logo
}

func Help() string {
	help := `
Usage: glock [options]

Options:
  -h, --help          Display this help message

  -a, --add           Mine and add a new block
  -i, --init          Initialize genesis block
  -p, --print         Print entire blockchain
  -r, --reset         Delete blockchain instance
  -s, --stats         Show all chain statistics
  -v, --validate      Check blockchain integrity
`
	return help
}

func Run() int {
	if len(os.Args) < 2 {
		fmt.Println(Verlogo())
		return 0
	}

	if len(os.Args) > 2 {
		ohNoes("Use only one argument at a time!")
	}

	arg := strings.ToLower(os.Args[1])

	switch arg {
	case "-h", "--help", "help":
		fmt.Println(Logo())
		fmt.Println(Help())
		return 0

	case "-a", "--add", "add":
		state := blockchainState()

		if state == "none" || state == "empty" {
			ohNoNoes("Genesis does not exist! Run 'glock init' to create it.", nil)
			return 1
		}

		if len(Blockchain) == 0 {
			ohNoNoes("No blockchain found! Initialize with 'glock init'.", nil)
			return 1
		}

		fmt.Print("Enter BPM value: ")
		var bpm int
		_, err := fmt.Scanf("%d", &bpm)
		if err != nil {
			ohNoNoes("Invalid BPM value!", err)
			return 1
		}

		singleWell("Mining new block...")
		mutex.Lock()
		newBlock := generateBlock(Blockchain[len(Blockchain)-1], bpm)
		Blockchain = append(Blockchain, newBlock)
		mutex.Unlock()

		if err := saveBlockchain(); err != nil {
			ohNoNoes("Failed to save block:", err)
			return 1
		}

		spew.Dump(newBlock)
		greatSuccess("Block has been mined!")
		return 0

	case "-i", "--init", "init":
		Wiper()

		state := blockchainState()

		if state == "populated" || state == "genesis-only" {
			ohNoNoes("Blockchain already exists! Run 'glock reset' to reset it.", nil)
			return 1
		}

		singleWell("Initializing genesis...")
		t := time.Now()
		genesisBlock := Block{
			Index:      0,
			Timestamp:  t.String(),
			BPM:        0,
			Hash:       "",
			PrevHash:   "",
			Difficulty: difficulty,
			Nonce:      "",
		}
		genesisBlock.Hash = calculateHash(genesisBlock)

		mutex.Lock()
		Blockchain = append(Blockchain, genesisBlock)
		mutex.Unlock()

		if err := saveBlockchain(); err != nil {
			ohNoNoes("Failed to save genesis block:", err)
			return 1
		}

		spew.Dump(genesisBlock)
		greatSuccess("Genesis has begun!")
		return 0

	case "-p", "--print", "print":
		Wiper()

		state := blockchainState()

		if state == "none" || state == "empty" {
			ohNoNoes("Genesis does not exist! Run 'glock init' to create it.", nil)
			return 1
		}

		if err := loadBlockchain(); err != nil {
			ohNoNoes("Failed to load blockchain:", err)
			return 1
		}

		if state == "genesis-only" {
			singleWell("Blockchain contains only the genesis block:")
		} else {
			singleWell(fmt.Sprintf("Blockchain contains %d blocks:", len(Blockchain)))
		}

		printBlockchain()
		return 0

	case "-r", "--reset", "reset":
		Wiper()

		state := blockchainState()

		singleWell("Resetting blockchain...\n")

		if state == "none" || state == "empty" {
			tripleWell("Blockchain is already empty.")
			return 0
		}

		if !YesOrNo("This will DELETE the entire blockchain. Are you sure?") {
			singleWell("Reset cancelled.")
			return 0
		}

		if err := os.Remove(blockchainData); err != nil {
			ohNoNoes("Failed to delete data:", err)
			return 1
		}

		mutex.Lock()
		Blockchain = []Block{}
		mutex.Unlock()

		greatSuccess("Blockchain has been reset!")
		return 0

	case "-s", "--stats", "stats":
		Wiper()

		state := blockchainState()

		if state == "none" || state == "empty" {
			ohNoes("No data available!")
			return 0
		}

		if err := loadBlockchain(); err != nil {
			ohNoNoes("Failed to load blockchain:", err)
			return 1
		}

		showStats()
		return 0

	case "-v", "--validate", "validate":
		Wiper()
		singleWell("Validating blockchain integrity...\n")

		state := blockchainState()

		if state == "none" || state == "empty" {
			tripleWell("Nothing to validate!")
			return 0
		}

		if err := loadBlockchain(); err != nil {
			ohNoNoes("Failed to load blockchain:", err)
			return 1
		}

		validateBlockchain()
		return 0

	default:
		Wiper()
		fmt.Printf("Unknown argument: %v\n", os.Args[1])
		fmt.Println(Help())
		return 1
	}
}
