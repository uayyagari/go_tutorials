package main

import (
	"fmt"
)

func main() {
	var intNumber int8 = 127
	fmt.Println("intNumber:", intNumber)

	var floatNumber float32 = 3.14
	fmt.Println("floatNumber:", floatNumber)
	// Can't add intNumber directly to floatNumber. Need to cast one of them to the other type.
	fmt.Println("floatNumber:", floatNumber+float32(intNumber))
	// But can add 1 to floatNumber directly,
	fmt.Println("floatNumber:", floatNumber+1)
	//
	var floatNumber2 float32 = 2.2
	// This is a way to declare and assign a variable in one line without specifying the type.
	// The type is inferred from the value. and the scope of the variable is limited to the block.
	res := floatNumber2 / 1.1
	// This prints out the type of literal computation result.
	fmt.Printf("literal float division %T\n", res)
}
