package common

import "errors"

var (
	ErrorEmpty           = errors.New("username and password cannot be empty")
	ErrorUnknownUsername = errors.New("unknown username")
	ErrorWrongPassword   = errors.New("wrong password")
	ErrorCreateWallet    = errors.New("failed to create wallet")
	ErrorTransfer        = errors.New("failed to transfer")
	ErrorOperateDatabase = errors.New("failed to operate database")
	ErrorNoWallet        = errors.New("user do not have wallet")
	ErrorGetBalance      = errors.New("failed to get balance")
	ErrorGetDecimals     = errors.New("failed to get decimals")
)
