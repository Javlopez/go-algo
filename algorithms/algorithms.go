package algorithms

//Algorithms struct
type Algorithms struct{}

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

// i = 6
// j = 2
//if input[1](7) > input[2] 5 {
//  swap = input[2]5
//  input[2] =  input[1] 7
//  input[1] =  sawp = 5
//}
// 2,5,7
//{7, 2, 5, 1, 3, 4}
//
