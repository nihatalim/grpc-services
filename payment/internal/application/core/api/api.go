package api

import (
	"context"

	"github.com/nihatalim/grpc-services/payment/internal/application/core/domain"
	"github.com/nihatalim/grpc-services/payment/internal/ports"
)

type Application struct {
	db ports.DBPort
}

// func NewApplication(db ports.DBPort) *Application {
// 	return &Application{
// 		db: db,
// 	}
// }

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	// paymentId := payment.ID

	// pm, err := a.db.Get(ctx, string(paymentId))
	// if err != nil {
	// 	fmt.Errorf("payment çekilirken hata alındı")
	// }

	pm := domain.Payment{}

	pm.Status = "paid"

	// err = a.db.Save(ctx, &pm)
	// if err != nil {
	// 	fmt.Errorf("payment kaydedilemedi")
	// }

	return pm, nil
}
