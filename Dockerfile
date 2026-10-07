FROM golang:1.27

WORKDIR /app

COPY . .

RUN go build -o my-go-app .

EXPOSE 8086

CMD ["./my-go-app"]
