package example

import "fmt"

func ArithmeticExample(){
  // Addition
  fmt.Println("======== Addition ========")
  fmt.Printf("1 + 2 = %d \n", 1 + 2)

  // Substraction
  fmt.Println("======== Substraction ========")
  fmt.Printf("5 - 2 = %d \n", 5 - 2)

  //Multiplication
  fmt.Println("======== Multiplication ========")
  fmt.Printf("3 * 2 = %d \n", 3 * 2)

  // Division
  fmt.Println("======== Division ========")
  fmt.Printf("10 / 2 = %d \n", 10 + 2)

  // Modulus
  fmt.Println("======== Modulus ========")
  fmt.Printf("10 %% 2 = %d \n", 10 % 2)

  // Increment
  fmt.Println("======== Increment ========")
  x := 10
  fmt.Println("x = 10")  
  x++
  fmt.Printf("x++ = %d \n", x)
  
  // Decrement
  fmt.Println("======== Decrement ========")
  x = 10
  x--
  fmt.Printf("x-- = %d \n", x)
}

func AssignmentOperatorsExample()  {
  fmt.Println("x = 5")
  x := 5
  fmt.Println("======== += ========")
  x += 3
  fmt.Printf("x+=3 : %d \n", x)
  
  x = 5
  fmt.Println("======== += ========")
  x-=3
  fmt.Printf("x-=3 : %d \n", x)
  
  x = 5
  fmt.Println("======== += ========")
  x *= 3
  fmt.Printf("x*=3 : %d \n", x)
  
  x = 5
  fmt.Println("======== += ========")
  x/=3
  fmt.Printf("x/=3 : %d \n", x)
  
  x = 5
  fmt.Println("======== += ========")
  x%=3
  fmt.Printf("x%%=3 : %d \n", x)
 
}

func ComparationExample()  {
  // Equal to " == "
  value := 5 == 5
  fmt.Printf("5 == 5 : %v \n", value)
  // Not Equal " != "
  value = 5 != 5
  fmt.Printf("5 != 5 : %v \n", value)
  // Greater Than " > "
  value = 5 > 5
  fmt.Printf("5 > 5 : %v \n", value)
  // Less Than " < "
  value = 5 < 5
  fmt.Printf("5 < 5 : %v \n", value)
  // Less Than or equal to " <= "
  value = 5 <= 5
  fmt.Printf("5 <= 5 : %v \n", value)
 // Greater Than or equal to " >= "
  value = 5 >= 5
  fmt.Printf("5 >= 5 : %v \n", value)
}

func LogicalOperationExample()  {
  // Logical And
  x := 4
  val := x < 5 &&  x < 10
  fmt.Printf("x = %d\n", x)
  fmt.Printf(" x < 5 &&  x < 10 = %v\n", val)
  // Logical Or
  x = 4
  val = x < 5 ||  x < 10
  fmt.Printf("x = %d\n", x)
  fmt.Printf(" x < 5 ||  x < 10 = %v\n", val)
  
  // Logical Not
  x = 4
  val = !(x < 5 && x < 10)
  fmt.Printf("x = %d\n", x)
  fmt.Printf(" !(x < 5 && x < 10) = %v\n ", val)

}
