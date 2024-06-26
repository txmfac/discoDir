package main

import (
	"os"
	"fmt"
)

func openFile (name string) *os.File {
	file, err := os.Open(name);
	if (err != nil) {
		fmt.Printf("%sFile %s doesn't exists !%s\n", string(ColorRed), name, string(ColorReset));
		os.Exit(-1);
	}
	return file;
}
