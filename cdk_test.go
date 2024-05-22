package cdk

import (
	"testing"
	"time"
)

// Secret 16组秘钥 每组8位 十进制
var Secret = [][]int32{
	{78, 0, 43, 30, 88, 72, 53, 68},
	{59, 73, 75, 80, 63, 88, 16, 44},
	{62, 84, 75, 24, 1, 68, 61, 44},
	{21, 92, 48, 76, 49, 91, 48, 82},
	{22, 37, 34, 25, 35, 93, 75, 81},
	{77, 96, 87, 29, 56, 67, 43, 47},
	{61, 71, 85, 99, 26, 9, 96, 15},
	{56, 86, 77, 67, 2, 75, 67, 24},
	{74, 11, 21, 81, 91, 16, 74, 85},
	{3, 50, 37, 15, 38, 94, 27, 51},
	{38, 32, 45, 64, 13, 85, 6, 65},
	{59, 33, 41, 52, 96, 92, 32, 79},
	{44, 1, 7, 92, 61, 76, 82, 53},
	{60, 36, 93, 45, 13, 87, 43, 2},
	{97, 83, 87, 51, 87, 24, 96, 79},
	{56, 48, 90, 56, 37, 83, 65, 60},
}

// CharTable 字符表
var CharTable = []string{
	"A", "B", "C", "D", "E",
	"F", "G", "H", "J", "K",
	"L", "M", "N", "P", "Q",
	"R", "S", "T", "U", "V",
	"W", "X", "Y", "Z", "2",
	"3", "4", "5", "6", "7",
	"8", "9",
}

func TestCdk_Generate(t *testing.T) {
	c := New(Secret, CharTable)
	generate, err := c.Generate(100001)
	if err != nil {
		return
	}
	t.Log(generate)
}

func TestCdk_BatchGenerate(t *testing.T) {
	c := New(Secret, CharTable)
	batchGenerate, err := c.BatchGenerate(100001, 10)
	if err != nil {
		return
	}
	t.Log(batchGenerate)
}

func TestCdk_Parse(t *testing.T) {
	c := New(Secret, CharTable)
	parse, err := c.Parse("A3B4C5D6E7F8G9HJ")
	if err != nil {
		return
	}
	t.Log(parse)
}

func TestCdk_GenerateAndParse(t *testing.T) {
	c := New(Secret, CharTable)
	generate, err := c.Generate(100001)
	if err != nil {
		return
	}
	t.Log(generate)
	parse, err := c.Parse(generate)
	if err != nil {
		return
	}
	t.Log(parse)
}

func TestCdk_GenerateRandomSecret(t *testing.T) {
	s, err := GenerateRandomSecret()
	if err != nil {
		return
	}
	t.Log(s)
}

func TestCdk_BatchGeneratePerformance(t *testing.T) {
	c := New(Secret, CharTable)
	count := 1000000
	start := time.Now() // 获取当前时间
	_, err := c.BatchGenerate(100001, uint(count))
	if err != nil {
		t.Errorf("BatchGenerate returned an error: %v", err)
	}
	duration := time.Since(start) // 计算执行时间
	t.Logf("BatchGenerate took %v to generate %v codes", duration, count)
}

func TestCdk_GeneratePerformance(t *testing.T) {
	c := New(Secret, CharTable)
	count := 100000
	start := time.Now() // 获取当前时间
	for i := 0; i < count; i++ {
		_, err := c.Generate(100001 + i)
		if err != nil {
			t.Errorf("Generate returned an error: %v", err)
		}
	}
	duration := time.Since(start) // 计算执行时间
	t.Logf("Generate took %v to generate %v codes", duration, count)
}
