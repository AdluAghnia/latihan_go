package main

import (
  "fmt"
)

func slice()  {
  // Create Slice Using the []datatype{values} format
  myslice1 := []int{};
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))
  fmt.Println(myslice1);
  

  myslice2 := []string{"Go", "Pyhton", "PeHaPe", "C++"};
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2);

  // Create a Slice from array
  arr := [6]int{10,12,14,11,18,19}
  slice := arr[2:4]

  fmt.Printf("slice = %v\n", slice)
  fmt.Printf("length = %d\n", len(slice))
  fmt.Printf("slice = %v\n", cap(slice))

  // Using the make() function
  // make(type, size, cappacity)
  make_slice := make([]int, 5, 12)
  
  fmt.Printf("make_slice = %v\n", make_slice)
  fmt.Printf("cappacity = %d\n", cap(make_slice))
  fmt.Printf("length = %d\n",len(make_slice))
  make_slice2 := make([]int, 5)
  
  fmt.Printf("make_slice = %v\n", make_slice2)
  fmt.Printf("cappacity = %d\n", cap(make_slice2))
  fmt.Printf("length = %d\n",len(make_slice2))


}

func access_slice(){
  harga := []int{20,40,30}

  fmt.Println(harga[0])
  fmt.Println(harga[2])
}

func change_element_slice(){
  harga := []int{10, 20, 30}
  harga[1] = 50
  fmt.Println(harga[1])

}

func append_slice(){
  myslice := []int{1,2,3,4,5,6}


  fmt.Printf("myslice =  %v\n", myslice)
  fmt.Printf("myslice length =  %d\n", len(myslice))
  fmt.Printf("myslice cappacity =  %v\n", cap(myslice))

  myslice = append(myslice, 20, 21)

  fmt.Printf("myslice =  %v\n", myslice)
  fmt.Printf("myslice length =  %d\n", len(myslice))
  fmt.Printf("myslice cappacity =  %v\n", cap(myslice))

}
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
  slice()

  fmt.Println("=====================")
  access_slice()

  fmt.Println("=====================")
  change_element_slice()

  fmt.Println("=====================")
  append_slice()
}
