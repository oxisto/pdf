# pdf

[![build](https://github.com/oxisto/pdf/actions/workflows/build.yml/badge.svg)](https://github.com/oxisto/pdf/actions/workflows/build.yml)
[![Go
Reference](https://pkg.go.dev/badge/github.com/oxisto.svg)](https://pkg.go.dev/github.com/oxisto/pdf)

This library is a fork of https://github.com/rsc/pdf, which aims to fix some of the outstanding issues.

```bash
go get github.com/oxisto/pdf
```

## Basic Usage

In order to display a very basic plain text representation of the PDF file, the following snippet can be used:

```go
import (
    "fmt"
    "os"

    "github.com/oxisto/pdf"
)

func main() {
    r, err := pdf.Open("test.pdf")
    defer r.Close()
    if err != nil {
        panic(err)
    }
    for i := 1; i <= r.NumPage(); i++ {
        page := r.Page(i)
        fmt.Print(page.Content().Plain())
        fmt.Println()
    }
}
```

The [`pdfcat`](./cmd/pdfcat/pdfcat.go) utility that is also included basically demonstrates the same functionality.

```bash
go install github.com/oxisto/pdf/cmd/pdfcat@latest
pdfcat test.pdf
```
