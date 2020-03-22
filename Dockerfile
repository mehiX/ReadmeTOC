
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...
RUN ls -lart /go/bin

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/ReadmeTOC /app
ENTRYPOINT ["./app"]
CMD ["-help"]
LABEL Name=readmetoc Version=0.0.1

