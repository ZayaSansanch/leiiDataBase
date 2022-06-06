package com

import (
	"os"
)

// func NewDataBase(mesto string, name string) {
func NewDataBase(name string) {
	//	file, err := os.Create(mesto + "/" + name)
	//	if err != nil {
	//		fmt.Println("Unable to create file:", err)
	//		os.Exit(1)
	//	}
	//	defer file.Close()
	//	fmt.Println(file.Name())

	//	fmt.Println("Database created")

	// f, err := os.Create("created.file")
	// // if err != nil {
	// // 	panic(err)
	// // }
	// f.Close()

	os.Mkdir("data/"+name, 0777)
}
