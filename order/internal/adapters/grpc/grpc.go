package grpc

import (
	"context"

	"github.com/nihatalim/grpc-clients/golang/order"
	"github.com/nihatalim/grpc-services/order/internal/application/core/domain"
)

func (a Adapter) Create(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	orderItems := make([]domain.OrderItem, len(req.OrderItems))

	for _, item := range req.OrderItems {
		orderItem := domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		}

		orderItems = append(orderItems, orderItem)
	}

	domainOrder, err := a.api.PlaceOrder(ctx, domain.NewOrder(req.UserId, orderItems))
	if err != nil {
		return &order.CreateOrderResponse{}, err
	}

	return &order.CreateOrderResponse{
		OrderId: domainOrder.ID,
	}, nil

}

func (a Adapter) Get(ctx context.Context, req *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	return nil, nil
}
