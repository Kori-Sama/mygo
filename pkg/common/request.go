package common

type TransferRequest struct {
	Username   string `json:"username"`
	Passphrase string `json:"passphrase"`
	ToName     string `json:"toName"`
	Amount     string `json:"amount"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUser struct {
	Id   int
	Name string
}
