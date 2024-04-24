package orderdetail

type OrderDetail struct {
	ID        string `json:"id"`
	OrderID   string `json:"orderId"`
	ProductID string `json:"productId"`
	Amount    int    `json:"amount"`
}
