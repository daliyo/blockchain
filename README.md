# blockchain sample

```go
package main

import (
	"strconv"
	"time"

	"github.com/glink4/blockchain"
)

func main() {

	for i := 1; i < 10; i++ {
		newblock := blockchain.Block{}
		newblock.Index = int64(i)
		newblock.Timestamp = time.Now().String()
		newblock.Data = strconv.Itoa(i) + " blockchain"
		blockchain.AddBlock(newblock)
	}

	blockchain.PrintBlockChain()
}
```