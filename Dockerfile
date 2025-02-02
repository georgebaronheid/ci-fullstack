FROM golang:1.22.11-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o app .

FROM scratch

COPY --from=builder /app/app .

ENTRYPOINT ["./app"]

# build with: docker build -t georgebaronheid/ci-fullstack:latest .