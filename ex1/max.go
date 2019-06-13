package main

func max(inputs []int) int {
	maxVal := 0
	for _, v := range inputs {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func maxModify(inputs []int) int {
	maxVal := inputs[0]
	for i := 1; i < len(inputs); i++ {
		if inputs[i] > maxVal {
			maxVal = inputs[i]
		}
	}
	return maxVal
}
