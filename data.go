// glock/data.go

package main

import (
	"encoding/json"
	"os"
)

const blockchainData = "blockchain.json"

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func loadBlockchain() error {
	if !fileExists(blockchainData) {
		return nil // No data found
	}

	data, err := os.ReadFile(blockchainData)
	if err != nil {
		return nil
	}

	if len(data) == 0 {
		return nil
	}

	return json.Unmarshal(data, &Blockchain)
}

func saveBlockchain() error {
	data, err := json.MarshalIndent(Blockchain, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(blockchainData, data, 0644)
}

func blockchainState() string {
	if !fileExists(blockchainData) {
		return "none"
	}

	if err := loadBlockchain(); err != nil {
		return "error"
	}

	if len(Blockchain) == 0 {
		return "empty"
	} else if len(Blockchain) == 1 {
		return "genesis-only"
	}

	return "populated"
}
