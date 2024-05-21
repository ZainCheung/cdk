# cdk
[中文版](README.zh.md)

Generator and parser for redemption codes

## Introduction

This project is a generator and parser for redemption codes. It can generate unique redemption codes based on an increment id, and parse the redemption codes back into the original increment id. This project is suitable for applications that require the generation of large numbers of unique codes in a short period of time.


## Getting Started

### Installation

You can get the project by using the `go get` command:

```bash
go get github.com/ZainCheung/cdk
```

### Usage
In your Go code, you can use this project as follows:

```go
package main

import (
	"fmt"
	"log"
	"github.com/ZainCheung/cdk"
)

func main() {
	// get a new random secret table
	randomSecret, err := cdk.GenerateRandomSecret()
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
	c := cdk.New(randomSecret, CharTable)
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
```

In this example, we first generate a random secret table, then create a new `Generater` object with the secret table 
and a character table. We then use the `Generate` method to generate a redemption code based on an increment id, 
and the `Parse` method to parse the redemption code back into the original increment id.

## Performance

The `Generater` in this project is highly efficient. In benchmark tests, it was able to generate 100,000 redemption 
codes in approximately 1.35 seconds. This makes it suitable for applications that require the generation of large 
numbers of unique codes in a short period of time.

Please note that actual performance may vary depending on the specific hardware and software environment.