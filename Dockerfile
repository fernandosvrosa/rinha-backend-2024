FROM golang:1.22 as build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o server /app/api/cmd/main.go

FROM alpine:3.12
WORKDIR /app
COPY --from=build /app/server .
CMD ["./server"]
