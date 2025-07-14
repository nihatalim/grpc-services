package payment

import (
	"context"

	"github.com/nihatalim/grpc-clients/golang/payment"
	"github.com/nihatalim/grpc-services/order/internal/application/core/domain"
)

type Adapter struct {
	payment payment.PaymentClient
}

func (a Adapter) Charge(context.Context, *domain.Order) error {
	return nil
}
