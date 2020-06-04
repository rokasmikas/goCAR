# Build Go server
FROM golang:latest AS builder
ADD . /
WORKDIR /
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .


# Build Vue
FROM node:alpine AS node_builder
COPY --from=builder / ./
RUN npm install
RUN npm run build


# Prod
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
COPY --from=node_builder / ./web
RUN chmod +x ./main
EXPOSE 8080
CMD ./main
