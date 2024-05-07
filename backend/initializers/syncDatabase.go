package initializers

import "ecommerce-app/models"

func SynDatabase() {
	Db.AutoMigrate(
		&models.Product{},
		&models.User{},
		&models.Cart{},
	)
}
