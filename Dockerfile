FROM golang:latest

WORKDIR /app

# Copy go mod and sum files and download deps
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN ./gen.sh

# Build the Go app
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]