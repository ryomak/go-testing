package main

import(
	"fmt"
)

func main(){
	user := user{
		name:"user1",
		height:1.66,
		weight:52.0,
	}
	fmt.Println(user.getBMI())
	fmt.Println(user.getBMI2(newCalculator()))
}