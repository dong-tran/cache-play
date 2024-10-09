FROM golang:1.22 as builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o api

FROM debian
WORKDIR /app
COPY --from=builder /src/api .
CMD [ "/app/api" ]
EXPOSE 8080