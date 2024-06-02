package initializers

import (
	"ecommerce-app/models"
	"os"
)

func SynDatabase() {
	env := os.Getenv("ENV")
	if env == "local" {
		// NOTE: may fail to migrate for complex FK dependency all together
		Db.AutoMigrate(
			&models.Address{},
			&models.Product{},
			&models.User{},
			&models.Cart{},
			&models.Order{},
			&models.OrderItem{},
			&models.Payment{},
		)
	}
}
