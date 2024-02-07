package common

import "errors"

var (
	ErrorEmpty                = errors.New("username and password cannot be empty")
	ErrorUnknownUsername      = errors.New("unknown username")
	ErrorWrongPassword        = errors.New("wrong password")
	ErrorOperateDatabase      = errors.New("failed to operate database")
	ErrorNoWallet             = errors.New("user do not have wallet")
	ErrorBlockchainDisconnect = errors.New("blockchain is disconnected")
	ErrorTokenContract        = errors.New("failed to get token contract")
	// ErrorTransfer             = errors.New("failed to transfer")
	// ErrorGetBalance           = errors.New("failed to get balance")
	// ErrorGetDecimals          = errors.New("failed to get decimals")
	// ErrorCreateWallet         = errors.New("failed to create wallet")
)

var internalErrors = []error{
	ErrorOperateDatabase,
	ErrorBlockchainDisconnect,
	ErrorTokenContract,
}

func CheckInternalError(err error) bool {
	if err == nil {
		return false
	}

	for _, e := range internalErrors {
		if err == e {
			return true
		}
	}

	return false
}
