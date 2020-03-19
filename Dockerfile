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
FROM scratch
COPY --from=builder /go/bin/drivers /go/bin/drivers
# Run the hello binary.
ENTRYPOINT ["/go/bin/drivers"]