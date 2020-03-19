############################
# STEP 1 build executable binary
############################
FROM golang:latest AS builder
WORKDIR /go/src/github.com/dearrudam/maratona-fullcycle-drivers
COPY . .
RUN go get -d -v ./...
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/drivers ./cmd/drivers/main.go

############################
# STEP 2 build a small image
############################
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/bin/drivers .
COPY --from=builder /go/src/github.com/dearrudam/maratona-fullcycle-drivers/drivers.json .
ENV DRIVERS_SOURCE=/root/drivers.json
CMD ["./drivers"]