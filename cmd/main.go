package main

import (
	"github.com/ClareChu/tigger/cmd/pkg"
)

func main() {
	err := pkg.NewRoot().Execute()
	if err != nil {
		return
	}
}
