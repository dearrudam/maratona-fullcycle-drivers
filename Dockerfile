############################
# STEP 1 build executable binary
############################
FROM golang:latest AS builder
WORKDIR /go/src/github.com/dearrudam/maratona-fullcycle-drivers
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/drivers ./cmd/drivers/main.go

############################
# STEP 2 build a small image
############################
FROM scratch
WORKDIR /root/
COPY --from=builder /go/bin/drivers .
COPY --from=builder /go/src/github.com/dearrudam/maratona-fullcycle-drivers/drivers.json .
ENV DRIVERS_SOURCE=/root/drivers.json
CMD ["./drivers"]