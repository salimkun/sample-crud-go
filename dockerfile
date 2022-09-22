FROM golang

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o user

ENTRYPOINT ["/app/user"]

EXPOSE 8080