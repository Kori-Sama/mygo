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

// This type is user's login information stored in the context of gin.
// The value comes from JWT token, so we can use it to get user's information.
type LoginUser struct {
	Id   int
	Name string
}
