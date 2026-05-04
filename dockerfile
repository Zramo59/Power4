FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o power4 .


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/power4 .
COPY index.html .
COPY static/ static/

EXPOSE 8080

CMD ["./power4"]
