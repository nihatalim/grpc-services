package main

import (
	"github.com/nihatalim/grpc-services/order/config"
	"github.com/nihatalim/grpc-services/order/internal/adapters/db"
	"github.com/nihatalim/grpc-services/order/internal/adapters/grpc"
	"github.com/nihatalim/grpc-services/order/internal/adapters/payment"
	"github.com/nihatalim/grpc-services/order/internal/application/core/api"
)

func main() {
	port := config.GetApplicationPort()
	datasourceUrl := config.GetDataSourceURL()
	paymentServiceUrl := config.GetPaymentServiceUrl()

	dbAdapter, err := db.NewAdapter(datasourceUrl)
	if err != nil {
		panic("database adapter could not be started")
	}

	paymentAdapter := payment.NewAdapter(paymentServiceUrl)

	application := api.NewApplication(dbAdapter, paymentAdapter)

	grpcAdapter := grpc.NewAdapter(port, application)
	grpcAdapter.Run()
}
