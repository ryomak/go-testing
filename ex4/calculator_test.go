package main

import(
	"testing"
)

func TestCalcBMI(t *testing.T){
	c := newCalculator()
	tCases := []struct{
		inputH float64
		inputW float64
		expected float64
	}{
		{1.75,70,22.857142857142858},
		{1.66,52,18.870663376397157},
		{0,33,0},
	}
	for _,tc := range tCases{
		acutual := c.calcBMI(tc.inputH,tc.inputW)
		if acutual != tc.expected{
			t.Errorf("expected:%v ,acutual:%v",tc.expected,acutual)
		}
	}
}