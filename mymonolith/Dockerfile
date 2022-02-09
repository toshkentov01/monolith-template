FROM golang:1.17.5-alpine3.14
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . ./
RUN go build -o main ./cmd/main.go 
RUN go mod tidy
CMD ["./main"]