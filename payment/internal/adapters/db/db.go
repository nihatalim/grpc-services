package db

import (
	"context"
	"fmt"

	"github.com/nihatalim/grpc-services/payment/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderID    int64
	TotalPrice float32
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(databaseUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(databaseUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	err := db.AutoMigrate(&Payment{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (a *Adapter) Get(ctx context.Context, id string) (domain.Payment, error) {
	return domain.Payment{}, nil
}

func (a *Adapter) Save(ctx context.Context, payment *domain.Payment) error {
	return nil
}
