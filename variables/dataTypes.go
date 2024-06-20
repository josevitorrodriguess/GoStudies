package variables

import (
	"errors"
	"fmt"
)

// Examples for data types
func types() {
	// Numbers int8, int16, int32, int64 numbers of bytes
	// If you use "int", Go will use the architecture of your computer
	var number int = 1000000

	// uint unsigned int
	var number2 uint32 = 1000

	// alias
	// int32 = rune
	var number3 rune = 12323

	// byte = uint8
	var number4 byte = 123

	//float32, float64
	//string (in go don't have char)
	/*in Go if you use '' he print the number value off ascci table for this char for exemple:  char := "B" in terminal the output is 66*/

	var err error = errors.New("inter error")

	fmt.Println(number, number2, number3, number4, err)
}
