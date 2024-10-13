FROM golang:1.22.5-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN GOARCH=amd64 go build -o main 

# Ensure the binary has executable permissions
RUN chmod +x main

CMD ["./main"]
