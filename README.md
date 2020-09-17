# TOC (Table of Contents) generator

## Table of contents

<!-- GENERATED TOC -->
1. [TOC (Table of Contents) generator](#toc-table-of-contents-generator)
    1. [Table of contents](#table-of-contents)
    2. [Description](#description)
    3. [Programming language](#programming-language)
    4. [Build the project](#build-the-project)
    5. [Usage](#usage)
        1. [Run it for only one file](#run-it-for-only-one-file)
        2. [Run as webserver](#run-as-webserver)
    6. [Build and run with Docker](#build-and-run-with-docker)
<!-- GENERATED TOC -->

## Description
Generates the table of contents for a markdown document, based on its headings. Can be used in 2 ways:
* as a web-server accepting a URL path to the location of the document
* as a one off for local files or remote URL's

## Programming language
Writen in GO

## Build the project
```bash
git clone https://github.com/mehiX/ReadmeTOC.git
cd ReadmeTOC
go get -d -v ./...
# command line
go build ./cmd/cli/*.go
# http server
go build ./cmd/http/*.go
```

## Usage

### Run it for only one file
The path parameter can be a local file or a URL

```bash
./toc URL
```

Prints out the table of contents. This can be then pasted inside the original file.


### Run as webserver

```bash
# Listen on 0.0.0.0:8080
./server :8080
```

The URL can be passed as a query parameter:

```bash
curl http://localhost:8080/query?path=https://raw.githubusercontent.com/mehiX/ReadmeTOC/master/README.md
```

or it can be send as part of a json request body:

```bash
curl -v -X GET \
  -d '{"url":"https://raw.githubusercontent.com/mehiX/ReadmeTOC/master/README.md"}' \
  -H "Content-Type: application/json" \
  http://localhost:8080/json
```


## Build and run with Docker

```bash
# Build the image
docker build -t readmetoc:1.1 .

# Run the container for 1 file (local)
docker run \
  -t \
  --rm \
  -v $(pwd)/README.md:/tmp/README.md \
  --entrypoint "/app/toc"
  readmetoc:1.1 /tmp/README.md

# Run for a URL
docker run \
  -t \
  --rm \
  --entrypoint "/app/toc" \
  readmetoc:1.1 https://github.com/mehiX/ReadmeTOC/raw/master/README.md
```