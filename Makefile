server_dev:
	cd backend && \
	CompileDaemon -command="./ecommerce-app"

test_stock_update:
	autocannon -c 100 -a 1000 -m "POST" http://localhost:8000/products/1/stock \
		--headers 'Content-Type: application/json' -b '{"action": "consume", "stock_quantity": 1}'
