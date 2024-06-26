package main

import (
	"ecommerce-app/initializers"
	"ecommerce-app/routers"
	"net/http"

	_ "ariga.io/atlas-provider-gorm/gormschema"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SynDatabase()
}

func main() {
	r := routers.SetupRouter()
	var port = ":8000"
	http.ListenAndServe(port, r)
}
