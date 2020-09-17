
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

ARG USER=toc

RUN apk --no-cache add ca-certificates
# Create a new user with no rights
RUN adduser --disabled-password --no-create-home --shell /no-shell ${USER}

WORKDIR /usr/local/app

COPY --from=builder --chown=${USER} /go/bin/toc ./
COPY --from=builder --chown=${USER} /go/bin/server ./
COPY --chown=${USER} templates ./templates/

RUN chmod 100 -R ./*
RUN chmod 500 -R ./templates

# Remove all users expect application user
RUN sed -i "/^${USER}:/!d" /etc/passwd

# Run unprivileged from here on
USER ${USER}

ENTRYPOINT ["./server"]
LABEL Name=readmetoc Version=0.0.1

