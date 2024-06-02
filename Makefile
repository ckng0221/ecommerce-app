# Run Dev Server
# make -j2 server_dev
server_dev: server_dev_backend server_dev_frontend

server_dev_backend:
	cd backend && \
	CompileDaemon -command="./ecommerce-app"

server_dev_frontend:
	cd frontend && \
	npm run dev

# Build
# make -j2 build
build: build_backend build_frontend

build_backend:
	cd backend && \
	go build .

build_frontend:
	cd frontend && \
	npm run build

# Run Prod Server
# make -j2 server_prod
server_prod: server_prod_backend server_prod_frontend

server_prod_backend:
	cd backend && \
	./ecommerce-app

# NOTE: On prod, use nginx to serve static, not available without docker
# Still use dev for local development
server_prod_frontend:
	cd frontend && \
	npm run dev

