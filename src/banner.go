package main

import "fmt"

func displayBanner() {
	banner := "					                                                              \n                                                                                  \n                                                       ,---,                      \n      ,---,  ,--,                                    .'  .' `\\    ,--,            \n    ,---.'|,--.'|                           ,---.  ,---.'     \\ ,--.'|    __  ,-. \n    |   | :|  |,      .--.--.              '   ,'\\ |   |  .`\\  ||  |,   ,' ,'/ /| \n    |   | |`--'_     /  /    '     ,---.  /   /   |:   : |  '  |`--'_   '  | |' | \n  ,--.__| |,' ,'|   |  :  /`./    /     \\.   ; ,. :|   ' '  ;  :,' ,'|  |  |   ,' \n /   ,'   |'  | |   |  :  ;_     /    / ''   | |: :'   | ;  .  |'  | |  '  :  /   \n.   '  /  ||  | :    \\  \\    `. .    ' / '   | .; :|   | :  |  '|  | :  |  | '    \n'   ; |:  |'  : |__   `----.   \\'   ; :__|   :    |'   : | /  ; '  : |__;  : |    \n|   | '/  '|  | '.'| /  /`--'  /'   | '.'|\\   \\  / |   | '` ,/  |  | '.'|  , ;    \n|   :    :|;  :    ;'--'.     / |   :    : `----'  ;   :  .'    ;  :    ;---'     \n \\   \\  /  |  ,   /   `--'---'   \\   \\  /          |   ,.'      |  ,   /          \n  `----'    ---`-'                `----'           '---'         ---`-'           \n                                                                                  "
	fmt.Println(string(ColorGreen), banner, string(ColorReset))
	fmt.Println(string(ColorBlue), "[Github := @txmfac]", string(ColorReset))
	fmt.Println(string(ColorRed), "Don't use this tool for illegal purposes !\n", string(ColorReset))
}
