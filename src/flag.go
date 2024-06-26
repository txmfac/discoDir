package main

import "flag"

type flagStruct struct {
	list string
	target string
	suf string
}

func flagInit (test * flagStruct) {
	flag.StringVar(&test.list, "l", "", "file with all the subdir to scan.")
	flag.StringVar(&test.target, "u", "0.0.0.0", "url of the target.")
	flag.StringVar(&test.suf, "s", "", "suffixes to add at the end of the url.")
	flag.Parse()
}
