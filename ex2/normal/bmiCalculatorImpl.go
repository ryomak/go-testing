package main

type bmiCalculator struct{}

func (b bmiCalculator)calcBMI(height,weight float64)float64{
  return weight/(height*height)
}



