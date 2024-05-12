package initializers

import "ecommerce-app/utils"

func LoadEnvVariables() {
	requiredEnvs := []string{"DB_URL", "ENV", "STRIPE_KEY", "FRONTEND_BASE_URL", "STRIPE_WEBHOOK_SECRET", "GOOGLE_CLIENT_ID", "GOOGLE_CLIENT_SECRET"}
	utils.LoadEnv(requiredEnvs)
}
