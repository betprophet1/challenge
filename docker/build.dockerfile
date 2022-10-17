FROM golang:1.18 AS builder
WORKDIR /backend/
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app ./main.go
RUN chmod +x /app

FROM alpine:3.10
WORKDIR /
ARG env
COPY --from=builder /app ./
ENTRYPOINT [ "/app" ]
