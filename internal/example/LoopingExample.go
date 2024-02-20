package example

import "fmt"

func Forloop() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
}

func ForLoopBreak() {
   for i:=0; i < 5; i++ {
    if i == 3 {
      break
    }
   fmt.Println(i)
  } 
}

func NestedLoop() {
  adj := [2]string{"big", "tasty"}
  fruits := [3]string{"apple", "banana", "orange"}

  for i:=0; i<len(adj);i++{
    for j:=0;j<len(fruits);j++{
      fmt.Println(adj[i],fruits[j])
    }
  }
}

func RangeLoop()  {
  fruits := [3]string{"banana", "apple", "orange"}

  for idx, val := range fruits {
    fmt.Printf("%v\t%v\n", idx, val)
  }

  fmt.Println("=============================")
  fmt.Print("With Out Indexing \n\n")

  for _, val := range fruits{
    fmt.Printf("%v\n", val)
  }
}
