package ports

import (
	"context"

	"github.com/nihatalim/grpc-services/order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(context.Context, *domain.Order) error
}
