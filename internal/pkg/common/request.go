package common

type TransferRequest struct {
	Passphrase string `json:"passphrase"`
	ToName     string `json:"toName"`
	Amount     string `json:"amount"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

type NewTransactionRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int    `json:"value"`
}

type TransactionRequest struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int    `json:"value"`
}

type CensorRequest struct {
	ID       int  `json:"id"`
	IsPassed bool `json:"is_passed"`
}

// This type is user's login information stored in the context of gin.
// The value comes from JWT token, so we can use it to get user's information.
type LoginUser struct {
	ID   int
	Name string
	Role Role
}
