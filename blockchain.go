package blockchain

import "fmt"

var chain []Block

func init() {
	block := CreateGenesisBlock()
	chain = append(chain, block)
}

func getLatestBlock() Block {
	return chain[len(chain)-1]
}

// AddBlock 添加一个区块
func AddBlock(newblock Block) {
	latestBlock := getLatestBlock()

	newblock.PreviousHash = latestBlock.Hash
	newblock.CalculateHash()

	chain = append(chain, newblock)
}

// PrintBlockChain 打印区块链内容
func PrintBlockChain() {
	fmt.Println(chain)
}
