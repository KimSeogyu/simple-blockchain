package main

import (
	"github.com/seogyugim/simple-blockchain/blockchain"
)

func main() {
	bc := blockchain.NewBlockChain()
	defer bc.Db.Close()

	cli := blockchain.CLI{bc}
	cli.Run()
}
