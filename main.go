package main

import (
	"fmt"

	"github.com/google/go-github/v48/github"
)

func main() {
	c := github.NewClient(nil)
	fmt.Println(c.BaseURL)
}
