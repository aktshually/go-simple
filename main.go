package main

import "fmt"

func main() {
	db := Database{
		Path: "./tests",
	}
	db.Connect(&Config{
		CreateIfDoesNotExist: true,
		Pattern:              "camelCase",
	})
	fmt.Println(db.Schemas)
}
