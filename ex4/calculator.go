package main


type ICalculator interface{
	calcBMI(height ,weight float64)float64
}

type calculator struct{}

func newCalculator()calculator{
	return calculator{}
}

func (calculator)calcBMI(height ,weight float64)float64{
	if height <=0 || weight <= 0{
		return 0.0
	} 
	return weight/(height *height)
}
