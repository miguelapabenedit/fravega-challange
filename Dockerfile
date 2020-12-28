FROM golang:latest

WORKDIR /app

COPY go.mod ./go.mod  
COPY go.sum ./go.sum 

RUN go mod download

COPY . .

RUN go build

CMD ["./fravega-challange"]