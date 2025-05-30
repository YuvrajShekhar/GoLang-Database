package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type Address struct {
	City string
	State string
	Country string
	Pincode json.Number
}

type User struct {
	Name string
	Age json.Number
	Contact string
	Company string
	Address Address
}

func main {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil{
		fmt.Println("Error",err)
	}

	employees := []User{
		{"Jhon,", "23", "2334568632", "Myrl Tech", Address{"bangalore", "karnataka", "India", "560013"}}
	}
}