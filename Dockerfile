# Build
FROM golang:1.17-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY api ./api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/go-app api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN addgroup -S go-app && adduser -S go-app -G go-app
USER go-app
WORKDIR /home/go-app
COPY --from=build /bin/go-app ./
ENTRYPOINT ["./go-app"]