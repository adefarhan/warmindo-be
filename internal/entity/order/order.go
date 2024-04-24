package order

import "time"

type Order struct {
	ID         string     `json:"id"`
	CustomerID string     `json:"customerId"`
	TotalPrice float64    `json:"totalPrice"`
	Status     string     `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
