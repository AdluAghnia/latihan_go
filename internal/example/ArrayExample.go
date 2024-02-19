package example

import "fmt"


func ArrayExample() {
  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"};
  harga := [4]int{10, 20, 23, 40};
  fmt.Println(cars);
  fmt.Println(harga);
  fmt.Println("Mobil = ", cars[2]);
  fmt.Println("Harga = ", harga[2]);
}
