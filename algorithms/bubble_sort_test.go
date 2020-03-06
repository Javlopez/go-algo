package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {

	for _, tc := range []struct {
		Title    string
		Input    []int
		Expected []int
	}{
		{Title: "First evaluation", Input: []int{7, 2, 5, 1, 3, 4}, Expected: []int{1, 2, 3, 4, 5, 7}},
		{Title: "With zero evaluation", Input: []int{7, 2, 5, 0, 3, 4}, Expected: []int{0, 2, 3, 4, 5, 7}},
		{Title: "With 2 zero evaluation", Input: []int{7, 2, 0, 0, 3, 4}, Expected: []int{0, 0, 2, 3, 4, 7}},
	} {
		t.Run(tc.Title, func(t *testing.T) {
			a := &Algorithms{}
			assert.Equal(t, tc.Expected, a.BubbleSort(tc.Input))
		})
	}
}
