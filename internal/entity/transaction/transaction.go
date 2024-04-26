package transaction

import "time"

type Transaction struct {
	ID            string     `json:"id"`
	OrderID       string     `json:"orderId"`
	MethodPayment string     `json:"methodPayment"`
	TotalPayment  float64    `json:"totalPayment"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}
