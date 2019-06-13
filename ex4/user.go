package main

import(
	"fmt"
)

type user struct{
	name string
	height float64
	weight float64
}


func (u user)getBMI()string{
	c := newCalculator()
	return fmt.Sprintf("%s's BMI is %.1f", u.name,c.calcBMI(u.height,u.weight))
}

func (u user)getBMI2(ic ICalculator)string{
	return fmt.Sprintf("%s's BMI is %.1f", u.name,ic.calcBMI(u.height,u.weight))
}