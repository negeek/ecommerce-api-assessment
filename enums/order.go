package enums

type OrderStatus struct{}

const (
	Pending   = "pending"
	Cancelled = "cancelled"
	Completed = "completed"
	Shipped   = "shipped"
	Delivered = "delivered"
)

func (o *OrderStatus) IsValid(status string) bool {
	switch status {
	case Pending, Cancelled, Completed, Shipped, Delivered:
		return true
	default:
		return false
	}
}
