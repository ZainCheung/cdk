# cdk
[English Version](README.md)

激活码生成器和解析器

## 介绍

这个项目是一个激活码生成器和解析器。它可以根据自增id生成唯一的激活码，并将激活码解析回原始的自增id。这个项目适合需要在短时间内生成大量唯一码的应用。

## 入门指南

### 安装

你可以使用 `go get` 命令来获取这个项目：

```bash
go get github.com/ZainCheung/cdk
```

### 使用
在你的 Go 代码中，你可以这样使用这个项目：

```go
package main

import (
	"fmt"
	"log"
	"github.com/ZainCheung/cdk"
)

func main() {
	// 生成一个新的随机秘钥表
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
	// 使用id生成一个激活码
	code, err := c.Generate(100001)
	if err != nil {
		log.Fatalf("Generate returned an error: %v", err)
	}
	fmt.Println("Generated code:", code)
	// 使用激活码解析出id
	id, err := c.Parse(code)
	if err != nil {
		log.Fatalf("Parse returned an error: %v", err)
	}
	fmt.Println("Parsed id:", id)
}
```

在这个例子中，我们首先生成一个随机秘钥表，然后使用这个秘钥表和一个字符表创建一个新的 `Generater` 对象。
然后我们使用 `Generate` 方法根据一个自增id生成一个激活码，并使用 `Parse` 方法将激活码解析回原始的自增id。

## 性能

这个项目中的 `Generater` 是非常高效的。在基准测试中，它能够在大约 3 秒内生成 1,000,000 个激活码。这使得它适用于需要在短时间内生成大量唯一码的应用。

请注意，实际性能可能会因具体硬件和软件环境而异。
