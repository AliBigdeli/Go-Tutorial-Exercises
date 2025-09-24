package main

import "fmt"

func main() {
	// Simple print function
	fmt.Println("Hello, World!")

	// Integer numbers
	var intNum int = 10                      // int size depends on architecture (32-bit or 64-bit)
	var int8Num int8 = 120                   // range: -128 to 127
	var int16Num int16 = 32000               // range: -32,768 to 32,767
	var int32Num int32 = 2000000000          // range: -2,147,483,648 to 2,147,483,647
	var int64Num int64 = 9000000000000000000 // range: -9,223,372,036,854,775,808 to +9,223,372,036,854,775,807

	var uintNum uint = 10                       // unsigned int (only positive)
	var uint8Num uint8 = 255                    // range: 0 to 255
	var uint16Num uint16 = 65535                // range: 0 to 65,535
	var uint32Num uint32 = 4294967295           // range: 0 to 4,294,967,295
	var uint64Num uint64 = 18446744073709551615 // range: 0 to 18,446,744,073,709,551,615

	fmt.Println("Int:", intNum)
	fmt.Println("Int8:", int8Num, "Int16:", int16Num, "Int32:", int32Num, "Int64:", int64Num)
	fmt.Println("Unsigned Int:", uintNum)
	fmt.Println("Uint8:", uint8Num, "Uint16:", uint16Num, "Uint32:", uint32Num, "Uint64:", uint64Num)

	// Floating-point numbers
	var float32Num float32 = 3.14  // 32-bit float: ~ ±3.4e38
	var float64Num float64 = 19.90 // 64-bit float: ~ ±1.7e308
	fmt.Println("Floats:", float32Num, float64Num)

	// Math Operations
	// + addition, - subtraction, * multiplication, / division, % modulo
	// Go does not allow operations between mixed types directly
	// Example: multiplying float32 and int requires explicit conversion
	var result float32 = float32Num * float32(intNum)
	fmt.Println("Result of multiplication:", result)

	// Simple string
	var str1 string = "Hello"
	var str2 string = "World"

	// Concatenation
	var resultString string = str1 + " " + str2
	fmt.Println("Concatenated:", resultString)

	// String length
	fmt.Println("Length of str1:", len(str1))

	// Formatting strings
	var name string = "Alice"
	var age int = 30
	var formattedString string = fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	fmt.Println("Formatted:", formattedString)
}
