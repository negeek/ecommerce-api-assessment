package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	middlewareenum "github.com/negeek/ecommerce-api-assessment/enums"
	"github.com/negeek/ecommerce-api-assessment/repositories"
	"github.com/negeek/ecommerce-api-assessment/services"
	"github.com/negeek/ecommerce-api-assessment/utils"
)

type OrdersController struct {
	orderService *services.OrderService
}

func NewOrdersController(orderService *services.OrderService) *OrdersController {
	return &OrdersController{orderService: orderService}
}

func (oc *OrdersController) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var order repositories.Order
	err := utils.Unmarshall(r.Body, &order)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadRequest, "Error reading request data", nil)
		return
	}

	userID := r.Context().Value(middlewareenum.UserContextKey).(int)
	order.UserID = userID

	err = oc.orderService.PlaceOrder(&order)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadGateway, err.Error(), nil)
		return
	}

	utils.JsonResponse(w, true, http.StatusCreated, "Order placed successfully", order)
}

func (oc *OrdersController) ListOrders(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewareenum.UserContextKey).(int)

	orders, err := oc.orderService.ListOrders(userID)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadRequest, "Failed to retrieve orders", nil)
		return
	}

	utils.JsonResponse(w, true, http.StatusOK, "Orders retrieved successfully", orders)
}

func (oc *OrdersController) CancelOrder(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	var order repositories.Order
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadRequest, "Invalid order ID", nil)
		return
	}

	userID := r.Context().Value(middlewareenum.UserContextKey).(int)
	order.ID = id
	order.UserID = userID
	err = oc.orderService.CancelOrder(&order, true)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.JsonResponse(w, true, http.StatusOK, "Order canceled successfully", nil)
}

func (oc *OrdersController) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	var order repositories.Order
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadRequest, "Invalid order ID", nil)
		return
	}
	err = utils.Unmarshall(r.Body, &order)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadRequest, "Error reading request data", nil)
		return
	}
	order.ID = id
	err = oc.orderService.UpdateOrder(&order, true)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusInternalServerError, "Failed to update order status", nil)
		return
	}

	utils.JsonResponse(w, true, http.StatusOK, "Order status updated successfully", nil)
}
