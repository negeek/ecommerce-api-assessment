package services

import (
	"errors"

	orderenum "github.com/negeek/ecommerce-api-assessment/enums"
	"github.com/negeek/ecommerce-api-assessment/repositories"
)

type OrderService struct{}

func (s *OrderService) PlaceOrder(order *repositories.Order) error {
	if order.Status == "" {
		order.Status = orderenum.Pending
	}
	statusChecker := &orderenum.OrderStatus{}
	if !statusChecker.IsValid(order.Status) {
		return errors.New("invalid order status")
	}
	return order.Create()
}

func (s *OrderService) ListOrders(userID int) ([]repositories.Order, error) {
	orders, err := repositories.FindOrdersByUserID(userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *OrderService) CancelOrder(order *repositories.Order, partial bool) error {
	reforder := &repositories.Order{}
	err := reforder.FindByID(order.ID)
	if err != nil {
		return errors.New("order does not exist")
	}
	if reforder.UserID != order.UserID {
		user := &repositories.User{}
		user.FindByID(order.UserID)
		if !user.IsAdmin() {
			return errors.New("access denied: you do not own this order")
		}
	}
	if reforder.Status != orderenum.Pending {
		return errors.New("cannot cancel order: only pending orders can be canceled")
	}
	order.UserID = reforder.UserID
	order.Status = orderenum.Cancelled
	if partial {
		return order.Patch()
	}
	return order.Put()

}

func (s *OrderService) UpdateOrder(order *repositories.Order, partial bool) error {
	if partial {
		return order.Patch()
	}
	return order.Put()
}
