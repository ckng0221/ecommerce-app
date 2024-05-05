package initializers

import "ecommerce-app/utils"

func LoadEnvVariables() {
	requiredEnvs := []string{"DB_URL"}
	utils.LoadEnv(requiredEnvs)
}
