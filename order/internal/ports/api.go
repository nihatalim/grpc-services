package ports

import (
	"context"

	"github.com/nihatalim/grpc-services/order/internal/application/core/domain"
)

type APIPort interface {
	PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error)
}
