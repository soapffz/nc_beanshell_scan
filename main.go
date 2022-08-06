package main

import (
	"fmt"
	"nc_beanshell_scan/cmd"
	"time"
)

func main() {
	start := time.Now()
	cmd.Execute()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("\ntime used:%s\n", delta)
}
