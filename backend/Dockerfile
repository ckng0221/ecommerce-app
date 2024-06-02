# Build the application from source
FROM golang:1.22 AS builder
ARG service

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /${service}

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS production
ARG service
ENV ENV=production

WORKDIR /

COPY --from=builder /${service} /go-app

EXPOSE 8000

# USER nonroot:nonroot

ENTRYPOINT ["/go-app"]

# Eg.
# docker build --build-arg="service=backend" -f ./backend/Dockerfile -t ecommerce-backend
# docker run -p 8000:8000 --env-file ./backend/.env ecommerce-backend