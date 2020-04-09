package blockchain

import "testing"

func TestBlock_CalculateHash(t *testing.T) {
	tests := []struct {
		name  string
		block Block
		want  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.block.CalculateHash(); got != tt.want {
				t.Errorf("Block.CalculateHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
