FROM golang:1.17-buster AS builder

WORKDIR /app
COPY . .

RUN go get
RUN go build

FROM debian:buster-slim AS runner

WORKDIR /app
COPY --from=builder /app/gocrud .

