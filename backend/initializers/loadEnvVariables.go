package initializers

import "ecommerce-app/utils"

func LoadEnvVariables() {
	requiredEnvs := []string{"DB_URL", "ENV", "STRIPE_KEY", "FRONTEND_BASE_URL"}
	utils.LoadEnv(requiredEnvs)
}
