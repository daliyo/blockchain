package blockchain

import (
	"testing"
)

func TestAddBlock(t *testing.T) {
	type args struct {
		newblock Block
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddBlock(tt.args.newblock)
		})
	}
}

func TestPrintBlockChain(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintBlockChain()
		})
	}
}
