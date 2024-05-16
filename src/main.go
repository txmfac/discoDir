package main

import (
	"flag"
	"fmt"
	"os"
)

type flagStruct struct {
	list string
	target string
}

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func displayBanner() {
	banner := "					                                                              \n                                                                                  \n                                                       ,---,                      \n      ,---,  ,--,                                    .'  .' `\\    ,--,            \n    ,---.'|,--.'|                           ,---.  ,---.'     \\ ,--.'|    __  ,-. \n    |   | :|  |,      .--.--.              '   ,'\\ |   |  .`\\  ||  |,   ,' ,'/ /| \n    |   | |`--'_     /  /    '     ,---.  /   /   |:   : |  '  |`--'_   '  | |' | \n  ,--.__| |,' ,'|   |  :  /`./    /     \\.   ; ,. :|   ' '  ;  :,' ,'|  |  |   ,' \n /   ,'   |'  | |   |  :  ;_     /    / ''   | |: :'   | ;  .  |'  | |  '  :  /   \n.   '  /  ||  | :    \\  \\    `. .    ' / '   | .; :|   | :  |  '|  | :  |  | '    \n'   ; |:  |'  : |__   `----.   \\'   ; :__|   :    |'   : | /  ; '  : |__;  : |    \n|   | '/  '|  | '.'| /  /`--'  /'   | '.'|\\   \\  / |   | '` ,/  |  | '.'|  , ;    \n|   :    :|;  :    ;'--'.     / |   :    : `----'  ;   :  .'    ;  :    ;---'     \n \\   \\  /  |  ,   /   `--'---'   \\   \\  /          |   ,.'      |  ,   /          \n  `----'    ---`-'                `----'           '---'         ---`-'           \n                                                                                  "
	fmt.Println(string(ColorGreen), banner, string(ColorReset))
	fmt.Println(string(ColorBlue), "[Github := @txmfac]", string(ColorReset))
	fmt.Println(string(ColorRed), "Don't use this tool for illegal purposes !\n", string(ColorReset))
}

func flagInit (test * flagStruct) {
	flag.StringVar(&test.list, "l", "", "file with all the subdir to scan.")
	flag.StringVar(&test.target, "u", "0.0.0.0", "url of the target.")
	flag.Parse()
}

func main() {
	displayBanner()
	test := flagStruct {}
	if (len(os.Args) < 2) {
		fmt.Println("Not enought parameters:\n\n\tUsage: ./discoDir -u [url of target (default 0.0.0.0)] -l [payload file]")
	}
	flagInit(&test)
	fmt.Println(test.list)
	fmt.Println(test.target)
}