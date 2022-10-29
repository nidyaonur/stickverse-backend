package utils

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestWeightedRandomCapacity(t *testing.T) {
	t.Run("weighted check", func(t *testing.T) {
		options := []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		optionResMap := map[uint]uint{
			0:  0,
			1:  0,
			2:  0,
			3:  0,
			4:  0,
			5:  0,
			6:  0,
			7:  0,
			8:  0,
			9:  0,
			10: 0,
		}
		maxCapacity := uint(10)
		for i := 0; i < 100000; i++ {
			capacity := WeightedRandomCapacity(options, maxCapacity)
			optionResMap[capacity]++
		}
		fmt.Println(optionResMap)
		assert.Equal(t, optionResMap[10], uint(0))
	})
}
