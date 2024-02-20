package example

import "fmt"

func IfStatmentExample()  {
  x:= 20
  y:= 18
  if x > y {
    fmt.Println("x is greater than y")
  }
}

func IfElseExample()  {
  temperature := 14 
  if (temperature > 15) {
    fmt.Println("It is warm out there")
  }else {
    fmt.Println("It is cold out there")
  } 
}

func ElseIfExample()  {
  time := 22
  
  if time < 10 {
    fmt.Println("Good morning.")
  } else if time < 20 {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }
}

func NestedIfExample()  {
  num := 20
  if num >= 10 {
    fmt.Println("Num is more than 10.")
    
    if num > 15 {
      fmt.Println("Num is also more than 15.")
    }

  }else {
    fmt.Println("Num is less than 10.")
  }
}
