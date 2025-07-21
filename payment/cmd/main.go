package main

import (
	"github.com/nihatalim/grpc-services/payment/config"
	"github.com/nihatalim/grpc-services/payment/internal/adapters/grpc"
	"github.com/nihatalim/grpc-services/payment/internal/application/core/api"
)

func main() {
	// dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database. Error: %v", err)
	// }

	// application := api.NewApplication(dbAdapter)
	application := api.NewApplication()
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
