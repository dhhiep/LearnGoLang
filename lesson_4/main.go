package main

import "fmt"

// func func_name (params) return_type { // content }

func sayHello() {
	fmt.Println("Hello World!!!")
}

func sayHello2(name string) {
	fmt.Println("Hello", name)
}

func sayHello3(name string) string {
	result := fmt.Sprintf("Hello %s", name)
	return result
}

// Multiple return values
func recInfo(w, h int) (int, int, int) {
	area := w * h
	return w , h, area
}

// Named return values
func recInfo2(w, h int) (width int, height int, area int, isSquare bool) {
	width = w
	height = h
	isSquare = w == h
	area = w * h
	return
}


func main(){
	fmt.Println("============ Passing argument into functions  ==============")
	sayHello()
	sayHello2("Hiep")
	fmt.Println(sayHello3("Hiep"))

	fmt.Println("============ Multiple return values ==============")
	w, h, area := recInfo(2, 3)
	fmt.Println("Width:", w)
	fmt.Println("Height:", h)
	fmt.Println("Area:", area)

	fmt.Println("============ Named return values ==============")
	w2, h2, area2, isSquare2 := recInfo2(3, 3)
	fmt.Println("Width:", w2)
	fmt.Println("Height:", h2)
	fmt.Println("Area:", area2)
	fmt.Println("isSquare2:", isSquare2)
}
