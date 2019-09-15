FROM golang:1.13.0-buster
LABEL Luis M. <luistm@gmail.com>
RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -i -o main ./cmd/api
CMD ["./main"]