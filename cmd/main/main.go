package main

import (
	"fmt"
  "latihan.com/latihan/internal/example_slices"
)


func tipe() {
  // Jenis-jenis tipe data sederhana di Go
  var a bool =  true
  var b int = 5
  var c float32 = 3.14
  var d string = "Hi Hallo"

  fmt.Println("Boolean = ", a)
  fmt.Println("Integer = ", b)
  fmt.Println("Float = ", c)
  fmt.Println("String = ", d)
}


func array() {
  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"};
  harga := [4]int{10, 20, 23, 40};
  fmt.Println(cars);
  fmt.Println(harga);
  fmt.Println("Mobil = ", cars[2]);
  fmt.Println("Harga = ", harga[2]);
}

func main() {
  fmt.Println("=====================")
  array()
  fmt.Println("=====================")
  tipe()
  fmt.Println("=====================")
  eslice.Slice()
  fmt.Println("=====================")
  eslice.AccessSlice()
  fmt.Println("=====================")
  eslice.AppendSlice()
  fmt.Println("=====================")
  eslice.ChangeElementSlice()
}
