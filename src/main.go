package main

import (
	"fmt"
	"os"
)

func main() {
	displayBanner()
	test := flagStruct {}
	if (len(os.Args) < 2) {
		fmt.Println("Usage: ./discoDir -u [url of target (default 0.0.0.0)] -l [payload file] -u suffixe.")
		os.Exit(-1)
	}
	flagInit(&test)
	throughtFile(test)
}