package eslice

import "fmt"

func Slice()  {
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
func AccessSlice(){
  harga := []int{20,40,30}

  fmt.Println(harga[0])
  fmt.Println(harga[2])
}

func ChangeElementSlice(){
  harga := []int{10, 20, 30}
  harga[1] = 50
  fmt.Println(harga[1])

}

func AppendSlice(){
  myslice := []int{1,2,3,4,5,6}


  fmt.Printf("myslice =  %v\n", myslice)
  fmt.Printf("myslice length =  %d\n", len(myslice))
  fmt.Printf("myslice cappacity =  %v\n", cap(myslice))

  myslice = append(myslice, 20, 21)

  fmt.Printf("myslice =  %v\n", myslice)
  fmt.Printf("myslice length =  %d\n", len(myslice))
  fmt.Printf("myslice cappacity =  %v\n", cap(myslice))

}

