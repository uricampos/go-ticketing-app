package payment

import "time"

type Status string

var (
	StatusPending Status = "pending"
	StatusSuccess Status = "success"
	StatusFailure Status = "failure"
	StatusUnknown Status = "unknown"
)

type PaymentRequest struct {
	OrderID int     `json:"order_id"`
	UserID  int     `json:"user_id"`
	Amount  float64 `json:"amount"`
	Method  string  `json:"method"`
}

type PaymentEvent struct {
	PaymentID string    `json:"payment_id"`
	OrderID   int       `json:"order_id"`
	UserID    int       `json:"user_id"`
	Amount    float64   `json:"amount"`
	Method    string    `json:"method"`
	Status    Status    `json:"status"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
}
