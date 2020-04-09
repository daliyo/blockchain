package blockchain

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// Block 区块
type Block struct {
	Index        int64
	Timestamp    string
	Data         string
	PreviousHash string
	Hash         string
}

// CreateGenesisBlock 创建创世区块
func CreateGenesisBlock() Block {
	return Block{0, time.Now().String(), "I'm J.", "000", ""}
}

// CalculateHash 计算当前区块的HASH值
func (block Block) CalculateHash() string {
	sha256 := sha256.New()
	sha256.Write([]byte(strconv.FormatInt(block.Index, 10) + block.Timestamp + block.Data + block.PreviousHash))

	return fmt.Sprintf("%x", sha256.Sum(nil))
}
