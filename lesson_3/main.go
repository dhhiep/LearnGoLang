package main

import (
  "fmt"
  "math"
  "math/bits"
)


func main(){
  // Bool
  var myBool bool = true // false
  fmt.Println(myBool)

  var mySecondBool = false
  fmt.Println(mySecondBool)

  // String
  var myString string = "String"
  fmt.Println(myString)

  // int
  var myInt int = 123
  fmt.Println(myInt)

  // int 8, 16, 32, 64
  // 1. Range
  // Range int8: -128 .. 127
  fmt.Println(math.MinInt8)
  fmt.Println(math.MaxInt8)

  // Range int16: -32768 .. 32767
  fmt.Println(math.MinInt16)
  fmt.Println(math.MaxInt16)

  // Range int32: -2147483648 .. 2147483647
  fmt.Println(math.MinInt32)
  fmt.Println(math.MaxInt32)

  // Range int64: -9223372036854775808 .. 9223372036854775807
  fmt.Println(math.MinInt64)
  fmt.Println(math.MaxInt64)

  var defaultInt int = 9223372036854775807 // Phụ thuộc vào system đang chạy là 32-bit hay 64-bit
  fmt.Println(defaultInt)
  fmt.Printf("%T\n", defaultInt)

  // 2. Bits
  fmt.Println("===================")
  fmt.Println(bits.OnesCount8(math.MaxUint8))
  fmt.Println(bits.OnesCount16(math.MaxUint16))
  fmt.Println(bits.OnesCount32(math.MaxUint32))
  fmt.Println(bits.OnesCount64(math.MaxUint64))

  // Byte alias for uint8

  var myUint uint = 10
  fmt.Println(myUint)

  var myByte byte = 11
  fmt.Println(myByte)
  fmt.Printf("%T\n", myByte)

  var myByte2 byte = 255 // 0 .. 255, error append when value is 256 ./main.go:61:22: constant 256 overflows byte
  fmt.Println(myByte2)

  var a byte = 'A'
  fmt.Println(a) // => 65

  var a2 = 65
  fmt.Printf("%c\n", a2)

  // float
  var myFloat float64 = 10.01
  fmt.Println(myFloat)

  // complex z = a + bi
  var z1 complex64 = 10 + 1i
  fmt.Println(z1)

  var z2 complex64 = 10 + 3i
  fmt.Println(z2)

  var z = z1 + z2
  fmt.Println(z) // => 20 + 4i

  // Rune
  var myString2 string = "Nhẫn"
  fmt.Println(myString2)

  for i:=0; i < len(myString2); i++ {
    fmt.Printf("%c", myString2[i])
  } // => Nháº«n

  fmt.Println("")

  // Rune type
  // Không quan tâm code point cần bao byte
  // Alias của int 32
  // Trong Go, Rune type thường dùng để thể hiện Unicode code point
  runes := []rune(myString2)

  for i:=0; i < len(runes); i++ {
    fmt.Printf("%c", runes[i])
  } // => Nhẫn

  fmt.Println("")

  fmt.Println("===================")
  var myRune rune = 'A'
  fmt.Printf("%T\n", myRune)

  fmt.Println("===================")
  var myRune2 rune = '♥'
  fmt.Printf("%c - %U - %d\n", myRune2, myRune2, myRune2) // %U is Unicode / %d is HTML Entry decimal

  // ZERO VALUE
  // Initialize variable without assign value
  // 0 for numberic types
  // false for boolean type
  // "" for strings

  // Type conversions
  var a1 int = 1
  var b1 float64 = 2.1

  // fmt.Println(a1 + b1) #=> invalid operation: a1 + b1 (mismatched types int and float64)
  fmt.Println(float64(a1) + b1)

  // CONSTANT
  var PI = 3.14
  fmt.Println(PI)
  PI = 3.15
  fmt.Println(PI)

  PI2 := 3.16
  fmt.Println(PI2)
  PI2 = 3.17
  fmt.Println(PI2)

  const PI3 = 3.14
  fmt.Println(PI3)
  // PI3 = 3.15 // => cannot assign to PI3

  // UNSIGNED INT POINTER
}
