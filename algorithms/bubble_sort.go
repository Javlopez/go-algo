package algorithms

//BubbleSort method
func (a *Algorithms) BubbleSort(input []int) []int {
	for i := len(input); i > 0; i-- {
		for j := 1; j < i; j++ {
			if input[j-1] > input[j] {
				swap := input[j]
				input[j] = input[j-1]
				input[j-1] = swap
			}
		}
	}
	return input
}
