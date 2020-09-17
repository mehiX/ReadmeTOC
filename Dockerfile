
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./cmd/cli/*.go
RUN go install -v ./cmd/http/*.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
# Create a new user with no rights
RUN adduser --disabled-password --no-create-home --shell /no-shell toc

WORKDIR /usr/local/app

COPY --from=builder --chown=toc /go/bin/toc ./
COPY --from=builder --chown=toc /go/bin/server ./
COPY --chown=toc templates ./templates/

RUN chmod 100 -R ./*
RUN chmod 500 -R ./templates

# Run unprivileged from here on
USER toc

ENTRYPOINT ["./server"]
LABEL Name=readmetoc Version=0.0.1

