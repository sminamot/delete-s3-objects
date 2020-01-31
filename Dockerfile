FROM golang:1.13 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o delete-s3-objects

FROM alpine:3.11
WORKDIR /app
COPY --from=builder /app/delete-s3-objects .

ENTRYPOINT ["/app/delete-s3-objects"]
