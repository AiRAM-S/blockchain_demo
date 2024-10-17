package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	command := flag.String("command", "", "Command to execute (listblocks, newblock)")
	data := flag.String("data", "", "Data for the new block")

	flag.Parse()

	t := time.Now()

	bc := NewBlockchain()
	switch *command {
	case "listblocks":
		bci := bc.Iterator()
		for {
			block := bci.Next()

			fmt.Printf("Prev. hash: %x\n", block.PrevHash)
			fmt.Printf("Data: %s\n", block.Data)
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Printf("Pow: %v\n", NewProofOfWork(block).Validate())
			fmt.Println()
			if len(block.PrevHash) == 0 {
				break
			}
		}
	case "newblock":
		bc.AddBlock(*data)
	default:
		fmt.Println("Invalid command")
	}
	fmt.Println("Time using:", time.Since(t))

	// ex 2.4
	// bc := NewBlockchain()

	// bc.AddBlock("Send 1 BTC to Ivan")
	// bc.AddBlock("Send 2 more BTC to Ivan")

	// bci := bc.Iterator()
	// for {
	// 	block := bci.Next()

	// 	fmt.Printf("Prev. hash: %x\n", block.PrevHash)
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	fmt.Printf("Pow: %v\n", NewProofOfWork(block).Validate())
	// 	fmt.Println()
	// 	if len(block.PrevHash) == 0 {
	// 		break
	// 	}
	// }

	// fmt.Println("Time using:", time.Since(t))

	// for _, block := range bc.blocks {
	// 	fmt.Printf("Prev. hash: %x\n", block.PrevHash)
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	fmt.Printf("Pow: %v\n",NewProofOfWork(block).Validate())
	// 	fmt.Println()
	// }

}
