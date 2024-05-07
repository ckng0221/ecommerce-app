package initializers

import "ecommerce-app/utils"

func LoadEnvVariables() {
	requiredEnvs := []string{"DB_URL", "ENV", "STRIPE_KEY", "FRONTEND_BASE_URL", "STRIPE_CLI_WEBHOOK_SECRET"}
	utils.LoadEnv(requiredEnvs)
}
