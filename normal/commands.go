package normal

import (
	"fmt"

	"github.com/ZayaSansanch/leiiDataBase/com"
)

// func podScanCommand(command string) {
// 	fmt.Print("leiiDataBase > ")
// 	fmt.Scan(&command)
// 	fmt.Println("")

// 	return
// }

func ScanCommand() {
	var command string
	var while bool = true
	// var commands = map[string]string{
	// 	"h":   ".h",
	// 	"new": ".newDataBase",
	// }

	for while {
		command = ""

		// podScanCommand(command)

		fmt.Print("leiiDataBase > ")
		fmt.Scan(&command)
		// fmt.Println("")

		if command == ".h" {
			// fmt.Println("In this time this command is not valid: '.h'")
			com.ComH()
		} else if command == ".newDataBase" {
			com.ComNewDataBase()
		} else if command == ".quit" {
			while = false
			fmt.Println("Thanck you, by used leiiDataBase!")
		} else {
			fmt.Println("In this time this command is not valid: ", command)
		}
	}
}
