package api

import (
	"context"

	"github.com/nihatalim/grpc-services/order/internal/application/core/domain"
	"github.com/nihatalim/grpc-services/order/internal/ports"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

func (a Application) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	err := a.payment.Charge(ctx, &order)
	if err != nil {
		return order, err
	}

	err = a.db.Save(ctx, &order)
	if err != nil {
		return order, err
	}

	return order, nil
}
