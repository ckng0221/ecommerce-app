package initializers

import (
	"ecommerce-app/models"
	"os"
)

func SynDatabase() {
	env := os.Getenv("ENV")
	if env == "local" {
		Db.AutoMigrate(
			&models.Product{},
			&models.User{},
			&models.Address{},
			&models.Cart{},
			&models.Order{},
			&models.OrderItem{},
			&models.Payment{},
		)
	}
}
