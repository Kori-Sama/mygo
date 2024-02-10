package common

import "errors"

var (
	ErrorEmpty            = errors.New("empty parameter")
	ErrorUnknownUsername  = errors.New("unknown username")
	ErrorWrongPassword    = errors.New("wrong password")
	ErrorNoWallet         = errors.New("user do not have wallet")
	ErrorBalanceNotEnough = errors.New("balance is not enough")
	ErrorInvalidAmount    = errors.New("invalid amount")
	ErrorInvalidToken     = errors.New("invalid token")
	ErrorExpiredToken     = errors.New("expired token")
	ErrorGetInfoFromToken = errors.New("failed to get info from token")
	// internal error
	ErrorOperateDatabase      = errors.New("failed to operate database")
	ErrorBlockchainDisconnect = errors.New("blockchain is disconnected")
	ErrorTokenContract        = errors.New("failed to get token contract")
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
