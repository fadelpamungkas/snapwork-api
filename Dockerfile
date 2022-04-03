FROM golang:1.17-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o server main.go

FROM alpine:3.10
WORKDIR /app
COPY --from=build /app/server .
COPY --from=build /app/.env .


EXPOSE 8080

CMD ["/app/server"]

