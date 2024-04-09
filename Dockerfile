FROM golang

RUN mkdir /app

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -v -o /app/bina ./cmd/api

EXPOSE 8080

CMD ["/app/bina"]