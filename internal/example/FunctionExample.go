package example

import "fmt"


// with param
func myFunction(fname string)  {
  fmt.Println("Hello", fname, "Tinja")
}

func multiParam(fname string, age int) {
  fmt.Println("Hello", age, "year old", fname, "Tinja")
}

func returnValue(x int, y int) int {
  return x + y
}

// Or

func returnValueNamed(x int, y int) (result int){
  result = x + y
  return
}

func multipleReturnValue(x int, y string) (result int, txt1 string){
  result = x + x
  txt1 = y + "world!"
  return
}

func FunctionExample(){
  myFunction("Adam")
  multiParam("Adam", 69)
  
  a := returnValue(4, 5)
  x, y := multipleReturnValue(23, "Hello")

  fmt.Println(a)
  fmt.Println(x)
  fmt.Println(y)
}
