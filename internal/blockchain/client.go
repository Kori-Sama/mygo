package blockchain

import (
	"context"
	"fmt"
	"math/big"
	"mygo/config"
	mycommon "mygo/internal/pkg/common"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func connectToChain() (*ethclient.Client, *big.Int, *Token, error) {
	conn, err := ethclient.Dial(config.Blockchain.GethClient)
	if err != nil {
		return nil, nil, nil, mycommon.ErrorBlockchainDisconnect
	}

	chainID, err := conn.ChainID(context.Background())
	if err != nil {
		return nil, nil, nil, mycommon.ErrorBlockchainDisconnect
	}

	token, err := NewToken(common.HexToAddress(config.Blockchain.ContractAddress), conn)
	if err != nil {
		return nil, nil, nil, mycommon.ErrorTokenContract
	}

	log.Infof("Connect to blockchain successfully. Chain ID: %s", chainID)
	return conn, chainID, token, nil
}

func Transfer(fromAddress, passphrase, toAddress string, amount *big.Int) error {
	conn, chainID, token, err := connectToChain()
	if err != nil {
		return err
	}
	defer conn.Close()

	ks := keystore.NewKeyStore(config.Blockchain.KeystorePath, keystore.LightScryptN, keystore.LightScryptP)

	fromAccount, err := ks.Find(accounts.Account{Address: common.HexToAddress(fromAddress)})
	if err != nil {
		return fmt.Errorf("failed to find account: %v", err)
	}

	err = ks.Unlock(fromAccount, passphrase)
	if err != nil {
		return fmt.Errorf("failed to unlock account: %v", err)
	}

	auth, err := bind.NewKeyStoreTransactorWithChainID(ks, fromAccount, chainID)
	if err != nil {
		return fmt.Errorf("failed to create authorized transactor: %v", err)
	}

	tx, err := token.Transfer(auth, common.HexToAddress(toAddress), amount)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %v", err)
	}

	log.Infof("Transfer successfully. Transaction hash: %s", tx.Hash().Hex())
	return nil
}

func NewAccount(passphrase string) (string, error) {
	ks := keystore.NewKeyStore(config.Blockchain.KeystorePath, keystore.LightScryptN, keystore.LightScryptP)
	account, err := ks.NewAccount(passphrase)
	if err != nil {
		return "", fmt.Errorf("failed to create account: %v", err)
	}

	log.Infof("Created new account successful. Account address: %s", account.Address.Hex())
	return account.Address.Hex(), nil
}

func BalanceOf(address string) (*big.Int, error) {
	conn, _, token, err := connectToChain()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	res, err := token.BalanceOf(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}, common.HexToAddress(address))
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %v", err)
	}
	return res, nil
}

func Decimal() (*big.Int, error) {
	conn, _, token, err := connectToChain()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	res, err := token.Decimals(nil)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetInt64(int64(res)), nil
}

// Todo! fixing the "invalid sender" error
func SendTransaction(cl *ethclient.Client, toStr string) error {
	const SK = "8443ff36077d13716c4643fbd24a2c166563953f2e9bd07a4ae5473f6327b799"
	const ADDR = "0x81a00791ad9052cd2eb51f81ed2c9bc5ef8c662f"
	var (
		sk       = crypto.ToECDSAUnsafe(common.FromHex(SK))
		to       = common.HexToAddress(toStr)
		value    = big.NewInt(1)
		sender   = common.HexToAddress(ADDR)
		gasLimit = uint64(1000000000)
	)
	// chainid, err := cl.ChainID(context.Background())
	// if err != nil {
	// 	return fmt.Errorf("failed to get chain id: %v", err)
	// }

	nonce, err := cl.NonceAt(context.Background(), sender, nil)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %v", err)
	}

	tx := types.NewTx(
		&types.LegacyTx{
			Nonce:    nonce,
			GasPrice: big.NewInt(1),
			Gas:      gasLimit,
			To:       &to,
			Value:    value,
			Data:     nil,
		})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(88)), sk)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %v", err)
	}

	return cl.SendTransaction(context.Background(), signedTx)
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
