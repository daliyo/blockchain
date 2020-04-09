package blockchain

import (
	"fmt"
	"strconv"
)

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
	newblock.Hash = newblock.CalculateHash()

	chain = append(chain, newblock)
}

// PrintBlockChain 打印区块链内容
func PrintBlockChain() {
	for _, block := range chain {
		fmt.Printf("index=%s, pre=%s, hash=%s, data=%s \n", strconv.FormatInt(block.Index, 10), block.PreviousHash, block.Hash, block.Data)
	}
}
