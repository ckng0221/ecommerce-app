server_dev:
	cd backend && \
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