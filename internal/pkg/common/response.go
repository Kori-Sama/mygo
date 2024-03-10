package common

type TransactionResponse struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int    `json:"value"`
	Status      Status `json:"status"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Role      Role   `json:"role"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type HistoryResponse struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	TransactionID int    `json:"transaction_id"`
	Action        Action `json:"action"`
	Timestamp     int64  `json:"timestamp"`
	NewValue      string `json:"new_value"`
}
