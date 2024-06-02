# Ecommerce App

Ecommerce App is a proof of concept (POC) e-commerce web application written in [Go](https://go.dev/) and [TypeScript](https://www.typescriptlang.org/).Users can add item to carts, perform checkout and payment. The application integrate with Stripe payment gateway, to perform payment using credit/debit card.

- `Backend`: The backend application server for handling the REST API.
- `Frontend`: The UI of the application.

## Tech Stacks

### Backend

- Programming Language: [Go](https://go.dev/)
- Server Framework: [Chi](https://go-chi.io/)
- Payment Gateway: [Stripe](https://stripe.com)
- ORM: [Gorm](https://gorm.io/)
- Database: [MySQL](https://www.mysql.com/)
- Authentication Protocol: [OIDC](https://openid.net/developers/how-connect-works/)
- Identity Provider: [Google](https://developers.google.com/identity)

### UI

- Programming Language: [TypeScript](https://www.typescriptlang.org/)
- JavaScript library: [React](https://react.dev/)
- CSS Framework: [Tailwind CSS](https://tailwindcss.com/)
- UI Library: [Material UI](https://mui.com/)
- Hosting: [Nginx](https://nginx.org/en/)

### Build

- CI Platform: [GitHub Actions](https://github.com/features/actions)
- Multi-container Tool: [Docker Compose](https://docs.docker.com/compose/)

## Getting Started

### Installation

```bash
# At the project root

$ npm install

# Install Go dev dependencies
$ go get github.com/githubnemo/CompileDaemon
$ go install github.com/githubnemo/CompileDaemon
```

Before running application, rename the `.env.example` files to `.env`, and update the environment variables accordingly.

## Run application

### On local

To run the application locally, ensure that `MySQL` is installed beforehand.

```bash
# At project root
# Development mode
make -j2 server_dev

# Build
make -j2 build

# Production mode
make -j2 server_prod

# Alternatively, you can navigate to the root of each application (e.g., ./apps/api) and run the npm scripts to run the particular application only.
```

### With docker and docker compose

To run the application using Docker, ensure that `Docker` and `Docker Compose` are installed beforehand.

```bash
# Create docker images and run docker containers in detached mode
docker compose up -d

# Stop and remove containers
docker compose down

# To access the docker applications on local, could browser:
http://ecommerce-app.127.0.0.1.nip.io
```
