# TOC (Table of Contents) generator

## Table of contents

<!-- GENERATED TOC -->
1. [TOC (Table of Contents) generator](#toc-table-of-contents-generator)
    1. [Table of contents](#table-of-contents)
    2. [Description](#description)
    3. [Usage](#usage)
    4. [Programming language](#programming-language)
    5. [Input and Output](#input-and-output)
        1. [Input](#input)
        2. [Output](#output)
    6. [Docker](#docker)
<!-- GENERATED TOC -->

## Description
Generates the table of contents for a markdown file, based on its headings.

## Usage
```bash
git clone https://github.com/mehix/go/ReadmeTOC.git
cd ReadmeTOC
go build
./ReadmeTOC <filepath>
```

## Programming language
Writen in GO

## Input and Output

### Input
Receives a path to a Markdown file. More help with `-help`

### Output
Prints out the table of contents. This can be then pasted inside the original file.

## Docker

```bash
# Build the image
docker build -t readmetoc:1.0 .
# Run the container
docker run -ti --rm -v $(pwd)/README.md:/tmp/README.md readmetoc:1.0 /tmp/README.md
```