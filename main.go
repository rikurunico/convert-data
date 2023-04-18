package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	splitData := strings.Split(data, ",")
	if len(splitData) == 3 {
		name := splitData[0]
		age, err := strconv.Atoi(splitData[1])
		if err != nil {
			panic("invalid age value") 
		}
		address := splitData[2]

		return User{
			Name:    name,
			Age:     age,
			Address: address,
		}
	}

	return User{}
}

func main() {
	data := "Edo,20,Jakarta"
	fmt.Printf("{ name: %q, age: %d, address: %q }\n", ConvertData(data).Name, ConvertData(data).Age, ConvertData(data).Address)
}
