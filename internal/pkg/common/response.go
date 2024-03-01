package common

type TransactionResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int    `json:"value"`
	Status      Status `json:"status"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
}
