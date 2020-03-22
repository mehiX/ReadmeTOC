# TOC (Table of Contents) generator

## Table of contents

<!-- GENERATED TOC -->
1. [TOC (Table of Contents) generator](#toc-table-of-contents-generator)
    1. [Table of contents](#table-of-contents)
    2. [Description](#description)
    3. [Programming language](#programming-language)
    4. [Usage](#usage)
        1. [Provide URL on command line](#provide-url-on-command-line)
            1. [Input and Output](#input-and-output)
                1. [Input](#input)
                2. [Output](#output)
        2. [Run as webserver](#run-as-webserver)
    5. [Docker](#docker)
<!-- GENERATED TOC -->

## Description
Generates the table of contents for a markdown file, based on its headings.

## Programming language
Writen in GO

## Usage
```bash
git clone https://github.com/mehiX/ReadmeTOC.git
cd ReadmeTOC
go build
```

### Provide URL on command line
```bash
./ReadmeTOC -path URL
```

#### Input and Output

##### Input
Receives a url or file path to a Markdown document. More help with `-help`

##### Output
Prints out the table of contents. This can be then pasted inside the original file.


### Run as webserver
```bash
# Listen on 0.0.0.0:8080
./ReadmeTOC -serve :8080
```

## Docker

```bash
# Build the image
docker build -t readmetoc:1.0 .
# Run the container
docker run -ti --rm -v $(pwd)/README.md:/tmp/README.md readmetoc:1.0 /tmp/README.md
```