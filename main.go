package main

import (
	"fmt"

	uinf "github.com/pirxthepilot/go-uname/pkg/unameinfo"
)

func main() {
	ui, err := uinf.NewUnameInfo()
	if err == nil {
		fmt.Printf(
			"%s %s %s %s %s %s\n",
			ui.Sysname,
			ui.Nodename,
			ui.Release,
			ui.Version,
			ui.Machine,
			ui.Domainname,
		)
	}
}
