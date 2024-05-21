package cdk

import (
	"fmt"
	"log"
)

func main() {
	// get a new random secret table
	randomSecret, err := GenerateRandomSecret()
	if err != nil {
		return
	}
	var CharTable = []string{
		"A", "B", "C", "D", "E",
		"F", "G", "H", "J", "K",
		"L", "M", "N", "P", "Q",
		"R", "S", "T", "U", "V",
		"W", "X", "Y", "Z", "2",
		"3", "4", "5", "6", "7",
		"8", "9",
	}
	c := New(randomSecret, CharTable)
	// get a code from an id
	code, err := c.Generate(100001)
	if err != nil {
		log.Fatalf("Generate returned an error: %v", err)
	}
	fmt.Println("Generated code:", code)
	// get an id from a code
	id, err := c.Parse(code)
	if err != nil {
		log.Fatalf("Parse returned an error: %v", err)
	}
	fmt.Println("Parsed id:", id)
}
