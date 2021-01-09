FROM golang:latest

WORKDIR /app

# Copy go mod and sum files and download deps
COPY go.mod go.sum ./
RUN go mod download

COPY . /app

WORKDIR /app

RUN make gen

# Build the Go app
RUN go build -o server .

EXPOSE 5000

CMD ["/app/server"]