package main

import (
	"fmt"
	"unicode/utf8"
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

	// Strings
	var str string = "Hello, World!"
	// Be careful with the length of the string. It is not the number of characters in the string for characters that are not ASCII.
	fmt.Println("Length of str:", len(str)) // prints out the number of bytes in the string.
	var str2 string = "ł"
	fmt.Println("str2:", str2)
	fmt.Println("Length of str2:", len(str2)) // prints out 2 because the character is not ASCII.
	fmt.Println("Length of str2 using utf8.RuneCountInString:", utf8.RuneCountInString(str2))

	var myRune rune = 'ł'
	fmt.Println("myRune:", myRune) // prints out the unicode code point of the character.

	var myBool bool = true // or false
	fmt.Println("myBool:", myBool)

	// Default values of types
	var intNum3 rune
	fmt.Println("intNum3:", intNum3) // prints out 0 because it is the default value of the type.
	var intNum4 int
	fmt.Println("intNum4:", intNum4) // prints out 0 because it is the default value of the type.
	var str1 string
	fmt.Println("str1:", str1) // prints out "" because it is the default value of the type.
	var myBool1 bool
	fmt.Println("myBool1:", myBool1) // prints out false because it is the default value of the type.

	// Constants
	// Use them if you don't want the value to change.
	const myConst int = 42
	fmt.Println("myConst:", myConst)
	// This will cause an error because constants can't be reassigned.
	// myConst = 43

	var1, var2 := 1, 2
	fmt.Println("var1:", var1)
	fmt.Println("var2:", var2)

	// Try to use the types when declaring variables. Makes readablity better.

}
