package utils

import (
	"fmt"
	"github.com/go-cmd/cmd"
)

type Command struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}

func Run(command *Command) (err error) {
	envCmd := cmd.NewCmd(command.Name, command.Args...)
	gms := <-envCmd.Start()
	if gms.Error != nil {
		fmt.Println(gms.Error)
		return gms.Error
	}
	for _, out := range gms.Stdout {
		fmt.Println(out)
	}
	return
}
