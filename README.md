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
import "github.com/ZainCheung/cdk"

func main() {
    c := cdk.New(Secret, CharTable)
    code, err := c.Generate(100001)
    if err != nil {
        log.Fatalf("Generate returned an error: %v", err)
    }
    fmt.Println("Generated code:", code)
}
```

In this example, we first create a new cdk instance, then call the Generate function to generate an activation code.

## Performance

The `Generater` in this project is highly efficient. In benchmark tests, it was able to generate 100,000 redemption codes in approximately 1.15 seconds. This makes it suitable for applications that require the generation of large numbers of unique codes in a short period of time.

Please note that actual performance may vary depending on the specific hardware and software environment.