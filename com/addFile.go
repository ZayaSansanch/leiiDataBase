package com

import (
	"fmt"
	"os"
)

func NewDataBase(mesto string, name string) {
	file, err := os.Create(mesto + "/" + name)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Println(file.Name())

	fmt.Println("Database created")
}
