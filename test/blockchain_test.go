package test

import (
	"mygo/blockchain"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
)

func TestBlockchain(t *testing.T) {
	conn, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		t.Error(err)
	}

	err = blockchain.SendTransaction(conn, "0xbdf642bf296be98aa36b637e3b97b66014d12213")
	if err != nil {
		t.Error(err)
	}
}
