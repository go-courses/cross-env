package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"github.com/go-courses/cross-env/bar"
)

func main() {
	var b bar.Twix
	fmt.Println(b.Name())
	spew.Dump(struct {
		Name  string
		Value string
	}{
		Name:  "Hello",
		Value: "world",
	})
}
