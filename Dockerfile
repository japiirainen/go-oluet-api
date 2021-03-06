FROM golang:latest AS builder

WORKDIR /go/src/github.com/japiirainen/go-oluet-api

# Copy go mod and sum files and download deps
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make gen
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/japiirainen/go-oluet-api/app .
COPY --from=builder /go/src/github.com/japiirainen/go-oluet-api/.env.production .env
COPY --from=builder /go/src/github.com/japiirainen/go-oluet-api/db/migrations/postgres ./db/migrations/postgres
COPY --from=builder /go/src/github.com/japiirainen/go-oluet-api/exel/data ./exel/data
COPY --from=builder /go/src/github.com/japiirainen/go-oluet-api/public/html ./public/html

EXPOSE 5000

CMD ["./app"]