# Server
server_dev:
	CompileDaemon -command="./ecommerce-app" 

test_stock_update:
	autocannon -c 100 -a 1000 -m "POST" http://localhost:8000/products/1/stock \
		--headers 'Content-Type: application/json' -b '{"action": "consume", "stock_quantity": 1}'

# Replace with your domain name
domain_name="drake-right-kitten.ngrok-free.app"

# Will forward traffic to localhost:8080
# ngrok http --domain=$(domain_name) 8000
host_server_ngrok:
	ngrok http --domain=$(domain_name) 8000


# ---- Test on local ----
# need to stripe login first
server_stripe_test:
	stripe listen --forward-to localhost:8000/payments/webhook

trigger_stripe_payment:
	stripe trigger payment_intent.succeeded

# -----------------------

# Migrations
include .env
# export

# To avoid conflict, could always only use manual migration only
mk_migration:
	atlas migrate diff --env dev --var dev_url=$(ATLAS_DB_URL_DEV) 

migrate_lint:
	atlas migrate lint --env dev --var dev_url=$(ATLAS_DB_URL_DEV) --latest

# Migrate based on version	
migrate_db:
	atlas migrate apply --env dev --var url=${ATLAS_DB_URL} --baseline 20240526074400

migrate_down:
	atlas migrate down --env dev --var url=${ATLAS_DB_URL} --to-version 20240526074400
# atlas migrate rm 20240526074627

# Migrate based on gorm models
# Will ignore migration files
# Use only either migrate_db or migrate_db_gorm, not both, as may have conflict
migrate_db_gorm:
	atlas schema apply --env local --auto-approve --var url=${ATLAS_DB_URL}

clear_migrations:
	rm -r migrations/