package blockchain

import (
	"context"
	"log"
	"math/big"
	"mygo/config"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func connectToChain() (*ethclient.Client, *big.Int, *Token, error) {
	conn, err := ethclient.Dial(config.Blockchain.GethClient)
	if err != nil {
		return nil, nil, nil, err
	}

	chainID, err := conn.ChainID(context.Background())
	if err != nil {
		return nil, nil, nil, err
	}

	token, err := NewToken(common.HexToAddress(config.Blockchain.ContractAddress), conn)
	if err != nil {
		return nil, nil, nil, err
	}

	return conn, chainID, token, nil
}

func Transfer(fromAddress, toAddress, passphrase string, amount *big.Int) error {
	_, chainID, token, err := connectToChain()
	if err != nil {
		return err
	}
	ks := keystore.NewKeyStore(config.Blockchain.KeystorePath, keystore.LightScryptN, keystore.LightScryptP)
	fromAccount, err := ks.Find(accounts.Account{Address: common.HexToAddress(fromAddress)})
	if err != nil {
		return err
	}
	err = ks.Unlock(fromAccount, passphrase)
	if err != nil {
		return err
	}
	auth, err := bind.NewKeyStoreTransactorWithChainID(ks, fromAccount, chainID)
	if err != nil {
		return err
	}
	tx, err := token.Transfer(auth, common.HexToAddress(toAddress), amount)
	if err != nil {
		return err
	}
	log.Println(tx)
	return nil
}

func NewAccount(passphrase string) (string, error) {
	ks := keystore.NewKeyStore(config.Blockchain.KeystorePath, keystore.LightScryptN, keystore.LightScryptP)
	account, err := ks.NewAccount(passphrase)
	if err != nil {
		return "", err
	}

	return account.Address.Hex(), nil
}

func BalanceOf(address string) (*big.Int, error) {
	_, _, token, err := connectToChain()
	if err != nil {
		return nil, err
	}
	res, err := token.BalanceOf(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Decimal() (uint8, error) {
	_, _, token, err := connectToChain()
	if err != nil {
		return 0, err
	}
	res, err := token.Decimals(nil)
	if err != nil {
		return 0, err
	}
	return res, nil
}

// func DeployContract(conn *ethclient.Client) {
// 	// ks := keystore.NewKeyStore(keystorePath, keystore.LightScryptN, keystore.LightScryptP)
// 	keyin := strings.NewReader(ownerInfo)

// 	auth, err := bind.NewTransactorWithChainID(keyin, ownerPassword, chainID)
// 	if err != nil {
// 		log.Fatalf("Failed to create authorized transactor: %v", err)
// 	}
// 	addr, tx, token, err := DeployToken(auth, conn, "MyGO Token", "MYGO", big.NewInt(1000))
// 	if err != nil {
// 		log.Fatalf("Failed to deploy new token contract: %v", err)
// 	}
// 	log.Println("addr:", addr.Hex())
// 	log.Println("tx:", tx.Hash())
// 	symbol, _ := token.Symbol(nil)
// 	log.Panicln("symbol:", symbol)
// }
