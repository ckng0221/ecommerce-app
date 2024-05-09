package initializers

import "ecommerce-app/models"

func SynDatabase() {
	Db.AutoMigrate(
		&models.Product{},
		&models.User{},
		&models.Address{},
		&models.Cart{},
		&models.Order{},
		&models.OrderItem{},
	)
}
