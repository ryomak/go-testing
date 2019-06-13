package main

import(
  "fmt"
)

type Person struct{
  name string
  height float64
  weight float64
}
func (person *Person)showBMI(calclator Calclator){
  fmt.Println(calclator.calcBMI(person.height,person.weight))
}

func main(){
  me := Person{
    name:"user1",
    height:1.75,
    weight:64.5,
  }
  c := bmiCalculator{}
  me.showBMI(c)
}
