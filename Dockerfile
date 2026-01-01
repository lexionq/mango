FROM golang:1.24.5-alpine AS builder    

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o mango main.go

FROM alpine
COPY --from=builder /app/mango /usr/local/bin/mango
ENTRYPOINT ["mango"]
CMD ["help"]


