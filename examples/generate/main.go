package main

import (
	"fmt"
	"github.com/xiecat/wsm/lib/shell"
	"github.com/xiecat/wsm/lib/shell/behinder"
)

func main() {
	pass, c := behinder.GenRandShell(shell.AspScript)
	fmt.Println(pass)
	fmt.Println(c)
}
