package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func throughtFile (flag flagStruct) {
	file := openFile(flag.list);
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := flag.target + scanner.Text() + flag.suf;
		fmt.Printf("url: %s\n", url)
		res, err := http.Get(url)
		if (err != nil) {
			fmt.Printf("%s[%s]: error with request%s\n", string(ColorRed), url, string(ColorReset))
		}
		if (res.StatusCode >= 200 && res.StatusCode <= 226) {
			fmt.Printf("%s[%s]: status code: %d%s\n", string(ColorGreen), url, res.StatusCode, string(ColorReset))
		}
		if (res.StatusCode >= 400 && res.StatusCode <= 451) {
			fmt.Printf("%s[%s]: status code: %d%s\n", string(ColorYellow), url, res.StatusCode, string(ColorReset))
		}
	}
}
