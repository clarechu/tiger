package main

import (
	"github.com/ClareChu/tiger/cmd/pkg"
)

func main() {
	err := pkg.NewRoot().Execute()
	if err != nil {
		return
	}
}
