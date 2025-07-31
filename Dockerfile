FROM registry.suse.com/bci/golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

FROM registry.suse.com/bci/bci-micro:latest AS final
COPY --from=builder /app/main /usr/bin
WORKDIR /app
ENTRYPOINT [ "main" ]