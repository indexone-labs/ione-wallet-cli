package main

import (
	"fmt"
	"math"
	"os"

	"github.com/multiversx/mx-sdk-go/data"
	"github.com/multiversx/mx-sdk-go/interactors"
)

func computeShardID(pubKey []byte) uint64 {
	startingIndex := len(pubKey) - 1

	usedBuffer := pubKey[startingIndex:]

	var addr uint64
	for i := 0; i < len(usedBuffer); i++ {
		addr = (addr << 8) + uint64(usedBuffer[i])
	}

	n := int(math.Ceil(math.Log2(3)))
	maskHigh := (1 << n) - 1
	maskLow := (1 << (n - 1)) - 1

	shard := addr & uint64(maskHigh)
	if shard > 2 {
		shard = addr & uint64(maskLow)
	}

	return shard
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("ERR: File path not given !!")
		return
	}
	pemKeyFile := args[0]

	w := interactors.NewWallet()

	var mnemonic data.Mnemonic
	var err error
	var privateKey []byte

	for {
		mnemonic, err = w.GenerateMnemonic()
		if err != nil {
			fmt.Println("Error generating mnemonic:", err)
			return
		}
		index0 := uint32(0)
		privateKey = w.GetPrivateKeyFromMnemonic(mnemonic, 0, index0)
		shardID := computeShardID(privateKey)
		if shardID == 1 {
			break
		}
	}

	fmt.Println("mnemonic:", mnemonic)
	interactors.NewWallet().SavePrivateKeyToPemFile(privateKey, pemKeyFile)
}
