package main

import (
	"fmt"
	"os"

	"github.com/oxisto/pdf"
)

func main() {
	r, err := pdf.Open(os.Args[1])
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
