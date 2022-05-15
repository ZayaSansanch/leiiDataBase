package com

import (
	"fmt"
)

// func name(name string) {
// 	fmt.Println("What name?")
// 	fmt.Scan(&name)
// }

func ComNewDataBase() {
	var name string
	var mesto string

	fmt.Println("What name?")
	fmt.Scan(&name)

	fmt.Println("Where?")
	fmt.Scan(&mesto)

	NewDataBase(mesto, name)
}
