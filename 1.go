package main

import (
	"fmt"
)

func main() {
	// datatypes
	var a int = 10
	var a1 int8 = 1<<7 - 1
	var a2 int16 = 10
	var a3 int32 = 10
	var a4 int64 = 10
	var a5 uint = 10
	var a6 uint8 = 10
	var a7 uint16 = 10
	var a8 uint32 = 10
	var a9 uint64 = 10
	var a10 uintptr = 10
	var b float64 = 10.10
	var c string = "Hello"
	var d bool = true
	fmt.Println("int: ", a, "\nint8: ", a1, "\nint16: ", a2, "\nint32: ", a3, "\nint64: ", a4, "\nuint: ", a5, "\nuint8: ", a6, "\nuint16: ", a7, "\nuint32: ", a8, "\nuint64: ", a9, "\nuintptr: ", a10, "\nfloat64: ", b, "\nstring: ", c, "\nbool: ", d)

	fmt.Printf("int pointer %p bit_pattern %b\n int8 pointer %p bit_pattern %b\n int16 pointer %p bit_pattern %b\n int32 pointer %p bit_pattern %b\n int64 pointer %p bit_pattern %b\n uint pointer %p bit_pattern %b\n uint8 pointer %p bit_pattern %b\n uint16 pointer %p bit_pattern %b\n uint32 pointer %p bit_pattern %b\n uint64 pointer %p bit_pattern %b\n uintptr pointer %p bit_pattern %b\n float64 pointer %p bit_pattern %b\n string pointer %p bit_pattern %b\n bool pointer %p bit_pattern %b\n", &a, a, &a1, &a1, &a2, &a2, &a3, &a3, &a4, &a4, &a5, &a5, &a6, &a6, &a7, &a7, &a8, &a8, &a9, &a9, &a10, &a10, &b, &b, &c, &c, &d, &d) 
}