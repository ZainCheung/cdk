// Package cdk is a generator and parser for redemption codes
package cdk

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// CdkI Cdk interface 兑换码接口
type CdkI interface {
	Generate(input int) (string, error)
	Parse(code string) (int, error)
	BatchGenerate(input int, count uint) ([]string, error)
}

// Cdk Cdk struct 兑换码
type Cdk struct {
	secret    [][]int32
	charTable []string
}

// New create a Cdk interface 创建兑换码接口
func New(secret [][]int32, charTable []string) CdkI {
	// check secret length
	if len(secret) != 16 {
		panic("secret length must be 16")
	}
	// check secret element length
	for _, s := range secret {
		if len(s) != 8 {
			panic("secret element length must be 8")
		}
	}
	// check charTable length
	if len(charTable) != 32 {
		panic("charTable length must be 32")
	}
	return &Cdk{
		secret:    secret,
		charTable: charTable,
	}
}

// ExampleSecret 16组秘钥 请确保不要泄露
var ExampleSecret = [][]int32{
	{3, 6, 7, 1, 22, 17, 23, 8},
	{9, 10, 11, 12, 13, 14, 15, 16},
	{18, 19, 20, 21, 24, 25, 26, 27},
	{28, 29, 30, 31, 32, 33, 34, 35},
	{36, 37, 38, 39, 40, 41, 42, 43},
	{44, 45, 46, 47, 48, 49, 50, 51},
	{52, 53, 54, 55, 56, 57, 58, 59},
	{60, 61, 62, 63, 64, 65, 66, 67},
	{3, 6, 7, 1, 22, 17, 23, 8},
	{9, 10, 11, 12, 13, 14, 15, 16},
	{18, 19, 20, 21, 24, 25, 26, 27},
	{28, 29, 30, 31, 32, 33, 34, 35},
	{36, 37, 38, 39, 40, 41, 42, 43},
	{44, 45, 46, 47, 48, 49, 50, 51},
	{52, 53, 54, 55, 56, 57, 58, 59},
	{60, 61, 62, 63, 64, 65, 66, 67},
}

// ExampleCharTable 字符表
var ExampleCharTable = []string{
	"A", "B", "C", "D", "E",
	"F", "G", "H", "J", "K",
	"L", "M", "N", "P", "Q",
	"R", "S", "T", "U", "V",
	"W", "X", "Y", "Z", "2",
	"3", "4", "5", "6", "7",
	"8", "9",
}

// Generate
// creates an activation code based on the input increment id.
//
// The function first converts the input increment id into a 32-bit binary string, and then divides this string into 8 groups, each with 4 bits.
//
// Then, the function generates a random number from 0 to 15 as the freshness value, and converts this freshness value into a 4-bit binary string.
//
// Then, the function takes out the secret key corresponding to the freshness value from the secret key table, and uses this secret key and each group of the increment id to perform weighted summation to get the signature.
//
// Finally, the function concatenates the signature, freshness value, and increment id into a 50-bit binary string, and then converts this string into a 10-character activation code.
//
// If an error occurs during the process, the function will return an empty string and error information.
//
// Parameters:
// input: Increment id, used to generate the activation code.
//
// Return values:
// string: The generated activation code.
// error: If an error occurs during the generation of the activation code, this value will contain error information.
//
// 根据输入的自增id生成一个激活码。
//
// 该函数首先将输入的自增id转换为32位的二进制字符串，然后将这个字符串分为8组，每组4位。
//
// 接着，函数生成一个0到15的随机数作为新鲜值，并将这个新鲜值转换为4位的二进制字符串。
//
// 然后，函数从秘钥表中取出对应新鲜值的秘钥，并用这个秘钥和自增id的每一组进行加权求和，得到签名。
//
// 最后，函数将签名、新鲜值和自增id拼接成一个50位的二进制字符串，然后将这个字符串转换为10个字符的激活码。
//
// 如果在过程中出现错误，函数将返回一个空字符串和错误信息。
//
// 参数:
// input: 自增id，用于生成激活码。
//
// 返回值:
// string: 生成的激活码。
// error: 如果在生成激活码的过程中出现错误，这个值将包含错误信息。
func (c *Cdk) Generate(incrementID int) (string, error) {
	return c.generateCode(incrementID)
}

func (c *Cdk) generateCode(incrementID int) (string, error) {
	// 生成10位的激活码
	// 1. 输入自增id
	// 2. 生成4位的新鲜值
	// 3. 从秘钥表中取出对应的秘钥
	// 4. 秘钥和自增id组合加权求和，得到签名
	// 5. 签名+新鲜值+自增id 组合成50位的二进制
	// 6. 二进制转为10组 5位的二进制 然后转为10进制 从字符表中取出对应的字符
	// 7. 返回激活码
	// 原始id
	// 转为32位二进制 前面补0
	id := fmt.Sprintf("%032b", incrementID)
	// 将自增id（32位）每4位分为一组，共8组，都转为10进制
	var inputArr []int32
	for i := 0; i < 8; i++ {
		decimalInt, err := strconv.ParseInt(id[i*4:i*4+4], 2, 32)
		if err != nil {
			fmt.Println("ParseInt error:", err)
			return "", fmt.Errorf("ParseInt error: %w", err)
		}
		inputArr = append(inputArr, int32(decimalInt))
	}
	// 随机0到15 转为 4位的二进制的新鲜值
	rand.NewSource(time.Now().UnixNano())
	randomFresh := rand.Intn(16)
	freshBinary := fmt.Sprintf("%04b", randomFresh)
	// 取出秘钥
	secretKey := c.secret[randomFresh]
	// 把每一组数加权求和，得到的结果就是签名
	var sign int32
	var newId string
	for i := 0; i < 8; i++ {
		sign += inputArr[i] * secretKey[i]
		newId += xorBinaryStrings(fmt.Sprintf("%04b", inputArr[i]), fmt.Sprintf("%04b", secretKey[i]))
	}
	// 生成最终签名 转为14位的二进制 最大值16383
	sign14 := fmt.Sprintf("%014b", sign)

	// 签名+fresh值+自增id
	// 生成最终的50位二进制
	finalBinaryCode := sign14 + freshBinary + newId

	// 二进制转为10组 5位的二进制 然后转为10进制 从字符表中取出对应的字符
	var result string
	for i := 0; i < 10; i++ {
		decimalInt, err := strconv.ParseInt(finalBinaryCode[i*5:i*5+5], 2, 32)
		if err != nil {
			fmt.Println("ParseInt error:", err)
			return "", fmt.Errorf("ParseInt error: %w", err)
		}
		result += c.charTable[decimalInt]
	}
	return result, nil
}

// Parse
//
// parses the activation code and returns the original increment id.
//
// The function first takes out 10 characters from the activation code, each character is converted into a 5-bit binary.
//
// Then, 10 5-bit binaries are concatenated into a 50-bit binary.
//
// Then, 14-bit signature, 4-bit fresh value, and 32-bit increment id are taken out from the 50-bit binary.
//
// Then, the secret key corresponding to the fresh value is taken out.
//
// Then, the increment id is decrypted with the secret key to get the original id.
//
// Finally, the original id and the signature are checked.
//
// If the check passes, return the original id. Otherwise, return an error message.
//
// Parameters:
//
// code: Activation code, used to parse the original increment id.
//
// Return values:
//
// int: The original increment id.
// error: If an error occurs during the parsing of the activation code, this value will contain error information.
//
// 解析激活码，返回原始的自增id。
//
// 该函数首先从激活码中取出10个字符，每个字符转为5位的二进制。
//
// 然后，10个5位的二进制拼接为50位的二进制。
//
// 接着，从50位的二进制中取出14位的签名，4位的fresh值，32位的自增id。
//
// 然后，从fresh值中取出对应的秘钥。
//
// 接着，用秘钥对自增id进行解密，得到原始id。
//
// 最后，用原始id和签名进行校验。
//
// 如果校验通过，返回原始id。否则，返回错误信息。
//
// 参数:
// code: 激活码，用于解析得到原始的自增id。
//
// 返回值:
//
// int: 原始的自增id。
// error: 如果在解析激活码的过程中出现错误，这个值将包含错误信息。
func (c *Cdk) Parse(code string) (int, error) {
	// 1. 从激活码中取出10个字符，每个字符转为5位的二进制
	// 2. 10个5位的二进制拼接为50位的二进制
	// 3. 从50位的二进制中取出14位的签名，4位的fresh值，32位的自增id
	// 4. 从fresh值中取出对应的秘钥
	// 5. 用秘钥对自增id进行解密，得到原始id
	// 6. 用原始id和签名进行校验
	// 7. 返回原始id
	// 从字符表中取出对应的字符
	binaryString, err := c.convertToBinary(code)
	if err != nil {
		return 0, err
	}
	// 从50位的二进制中取出14位的签名，4位的fresh值，32位的自增id
	signatureBinary := binaryString[:14]
	freshnessBinary := binaryString[14:18]
	incrementIDBinary := binaryString[18:]
	signNum, err := strconv.ParseInt(signatureBinary, 2, 32)
	if err != nil {
		fmt.Println("ParseInt error:", err)
		return 0, fmt.Errorf("ParseInt error: %w", err)
	}
	// 从fresh值中取出对应的秘钥
	freshInt, err := strconv.ParseInt(freshnessBinary, 2, 32)
	if err != nil {
		fmt.Println("ParseInt error:", err)
		return 0, fmt.Errorf("ParseInt error: %w", err)
	}
	secretKey := c.secret[freshInt]
	// 用秘钥对自增id进行异或运算，得到原始id
	var originalId string
	for i := 0; i < 8; i++ {
		originalId += xorBinaryStrings(incrementIDBinary[i*4:i*4+4], fmt.Sprintf("%04b", secretKey[i]))
	}
	lastId, err := strconv.ParseInt(originalId, 2, 32)
	if err != nil {
		fmt.Println("ParseInt error:", err)
		return 0, fmt.Errorf("ParseInt error: %w", err)
	}
	// 用原始id和签名进行校验
	var inputArr []int32
	for i := 0; i < 8; i++ {
		decimalInt, err := strconv.ParseInt(originalId[i*4:i*4+4], 2, 32)
		if err != nil {
			fmt.Println("ParseInt error:", err)
			return 0, fmt.Errorf("ParseInt error: %w", err)
		}
		inputArr = append(inputArr, int32(decimalInt))
	}
	var signInt int32
	for i := 0; i < 8; i++ {
		signInt += inputArr[i] * secretKey[i]
	}
	if signInt != int32(signNum) {
		return 0, fmt.Errorf("invalid code")
	}
	return int(lastId), nil
}

// BatchGenerate generates multiple activation codes based on the input increment id.
func (c *Cdk) BatchGenerate(startIncrementID int, numCodes uint) ([]string, error) {
	var result []string
	for i := 0; i < int(numCodes); i++ {
		code, err := c.generateCode(startIncrementID + i)
		if err != nil {
			return nil, err
		}
		result = append(result, code)
	}
	return result, nil
}

// convertToBinary converts a string to a binary string
func (c *Cdk) convertToBinary(s string) (string, error) {
	var binaryString string
	for i := 0; i < len(s); i++ {
		charIndex := -1
		for j, c := range c.charTable {
			if c == string(s[i]) {
				charIndex = j
				break
			}
		}
		if charIndex == -1 {
			return "", fmt.Errorf("invalid character in code: %v", s[i])
		}
		binaryString += fmt.Sprintf("%05b", charIndex)
	}
	return binaryString, nil
}

// GenerateRandomSecret generates a random secret key table.
//
// The function generates a 16x8 secret key table, each element is a random number between 0 and 99.
//
// The function prints the secret key table in the format of a Go slice, and you can copy and paste it into your code.
func GenerateRandomSecret() ([][]int32, error) {
	var randomSecretKeyTable [][]int32
	for i := 0; i < 16; i++ {
		var s []int32
		for j := 0; j < 8; j++ {
			s = append(s, int32(rand.Intn(100)))
		}
		randomSecretKeyTable = append(randomSecretKeyTable, s)
	}
	fmt.Println("Here is your random secret key table, please copy and paste it to your code:")
	// format print randomSecretKeyTable
	fmt.Println("var ExampleSecret = [][]int32{")
	for _, s := range randomSecretKeyTable {
		fmt.Printf("\t{%v, %v, %v, %v, %v, %v, %v, %v},\n", s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7])
	}
	fmt.Println("}")
	return randomSecretKeyTable, nil
}

// XOR operation between two binary strings
func xorBinaryStrings(s1, s2 string) string {
	var result string
	for i := 0; i < len(s1); i++ {
		b1 := s1[i] - '0'
		b2 := s2[i%len(s2)] - '0'
		result += strconv.Itoa(int(b1 ^ b2))
	}
	return result
}
