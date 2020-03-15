# building stage
FROM golang:1.13-alpine as builder
LABEL maintainer="Dmitry Shorokhov <dm-shorokhov@yandex.ru>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# running application stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/locale/translate.* ./locale/
EXPOSE 8080
CMD ["./main"]
