// glock/cli.go

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func Logo() string {
	logo := `           

        __              __    
.-----.|  |.-----.----.|  |--.
|  _  ||  ||  _  |  __||    < 
|___  ||__||_____|____||__|__|
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

func Help() string {
	help := `
Usage: glock [options]

Options:
  -h, --help          Display this help message
  -v, --version       Display the version number

  -a, --add           Mine and add a new block
  -i, --init          Execute genesis block
  -p, --print         Print entire blockchain
  -s, --stats         Show all chain stats
  -v, --validate      Check blockchain integrity
`
	return help
}

/*
func Version() string {
	version := "v0.1.0"

	return version
}
*/

func Run() int {
	if len(os.Args) < 2 {
		Wiper()
		fmt.Println(Logo())
		singleWell("")
		return 0
	}

	if len(os.Args) > 2 {
		tripleWell("Use only one argument at a time!")
	}

	arg := strings.ToLower(os.Args[1])

	switch arg {
	case "-h", "--help", "help":
		fmt.Println(Logo())
		fmt.Println(Help())
		return 0
	case "-a", "--add", "add":
		singleWell("Opening your config...")
		generateBlock()
		return 0
	case "-i", "--init", "init":
		singleWell("Creating genesis...")
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			t := time.Now()
			genesisBlock := Block{}
			genesisBlock = Block{0, t.String(), 0, calculateHash(genesisBlock), "", difficulty, ""}
			spew.Dump(genesisBlock)

			mutex.Lock()
			Blockchain = append(Blockchain, genesisBlock)
			mutex.Unlock()
		}()

		log.Fatal(run())
		return 0
	case "-p", "--print", "print":
		// Logic to print the entire blockchain
		return 0
	case "-s", "--stats", "stats":
		// Logic to print all chain stats
		// Such as length, total difficulty,
		// avg. mining time, etc
		return 0
	case "-v", "--validate", "validate":
		// Logic to check and print blockchain integrity
		return 0
	default:
		fmt.Errorf("Unknown argument: %v", os.Args[1])
		fmt.Println(Help())
		return 1
	}
}
