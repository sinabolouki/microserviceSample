package utils

import "os"

const ()

var (
	UserServiceAddress      = getEnv("USER_SERVICE_ADDRESS", "localhost:50051")
	CatalogueServiceAddress = getEnv("CATALOGUE_SERVICE_ADDRESS", "localhost:50052")
	OrderServiceAddress     = getEnv("ORDER_SERVICE_ADDRESS", "localhost:50053")
)

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
