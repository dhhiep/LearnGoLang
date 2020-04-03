package main

import "fmt"

func main(){
  var number int

  number = 10

  fmt.Println(number)

  var number1 int = 11

  fmt.Println(number1)

  // type inference

  var email = "Hiệp test"
  fmt.Println(email)

  // Khai bao nhieu bien
  // 1. Voi cung kieu dữ lieu

  var a, b int

  a = 1
  b = 2

  fmt.Println(a)
  fmt.Println(b)

  var a1, b1 int = 10, 11

  fmt.Println(a1)
  fmt.Println(b1)

  var a2, b2 = 20, 21

  fmt.Println(a2)
  fmt.Println(b2)

  // 2. Với khác kiểu dữ liện

  var (
    name string
    address string
    age int
  )

  name = "Hiep"
  address = "Lê Thị Bạch Cát"
  age = 29

  fmt.Println(name)
  fmt.Println(address)
  fmt.Println(age)

  var (
    name1 string = "Hiep 1"
    address1 string = "Lê Thị Bạch Cát 1"
    age1 int = 30
  )

  fmt.Println(name1)
  fmt.Println(address1)
  fmt.Println(age1)


  var (
    name2 = "Hiep 2"
    address2 = "Lê Thị Bạch Cát 2"
    age2 = 30
  )

  fmt.Println(name2)
  fmt.Println(address2)
  fmt.Println(age2)

  var name3, address3, age3 = "Hiep 3", "Lê Thị Bạch Cát 3", 30

  fmt.Println(name3)
  fmt.Println(address3)
  fmt.Println(age3)

  // Short hand
  short_hand := "Short hand"
  fmt.Println(short_hand)

  name4, address4, age4 := "Hiep 4", "Lê Thị Bạch Cát 4", 40

  fmt.Println(name4)
  fmt.Println(address4)
  fmt.Println(age4)
}
