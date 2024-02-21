package example

import "fmt"

type Person struct {
  name string
  age int
  job string
  salary int
}

func AccessStructMember() {
  var pers1 Person
  var pers2 Person

  // Pers1 Specification
  pers1.name = "Adam"
  pers1.age = 24
  pers1.job = "Software Enginner"
  pers1.salary = 4500

  // Pers2 Specification
  pers2.name = "Yaniq"
  pers2.age = 25
  pers2.job = "Marketing"
  pers2.salary = 4000

  // Accsess and print pers1 info

  fmt.Println("Name : ", pers1.name)
  fmt.Println("Age : ", pers1.age)
  fmt.Println("Job : ", pers1.job)
  fmt.Println("Salary : ", pers1.salary)

  // Access ang print pers2 info

  fmt.Println("Name : ", pers2.name)
  fmt.Println("Age : ", pers2.age)
  fmt.Println("Job : ", pers2.job)
  fmt.Println("Salary : ", pers2.salary)
}


// Pass Struct as Function Arguments

func printPerson(pers Person){
  fmt.Println("Name : ", pers.name)
  fmt.Println("Age : ", pers.age)
  fmt.Println("Job : ", pers.job)
  fmt.Println("Salary : ", pers.salary)
}

func PassStructAsArguments(){
  var pers1 Person
  var pers2 Person

  // Pers1 Specification
  pers1.name = "Adam"
  pers1.age = 24
  pers1.job = "Software Enginner"
  pers1.salary = 4500

  // Pers2 Specification
  pers2.name = "Yaniq"
  pers2.age = 25
  pers2.job = "Marketing"
  pers2.salary = 4000

  printPerson(pers1)

  fmt.Println("++++++++++++++++++++++++++")

  printPerson(pers2)
}
