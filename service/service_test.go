package service

import (
	"math/big"
	"testing"
)

func TestCalcToken(t *testing.T) {
	t.Run("01", func(t *testing.T) {
		balance := big.NewInt(1000000000000000000)
		decimal := uint8(18)
		want := 1.0
		if got := calcToken(balance, decimal); got != want {
			t.Errorf("calcToken() = %v, want %v", got, want)
		}
	})
	t.Run("02", func(t *testing.T) {
		balance := big.NewInt(100)
		decimal := uint8(18)
		want := 0.0000000000000001
		if got := calcToken(balance, decimal); got != want {
			t.Errorf("calcToken() = %v, want %v", got, want)
		}
	})
}
