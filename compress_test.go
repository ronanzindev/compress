package compress

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	testData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t.Run("should return only even number", func(t *testing.T) {
		comp := New(testData)
		result := comp.Filter(func(elem int) bool {
			return elem%2 == 0
		}).Collect()
		assert.Equal(t, []int{2, 4, 6, 8, 10}, result)
	})
	t.Run("reducer should return the sum", func(t *testing.T) {
		comp := New(testData)
		result := comp.Reduce(0, func(i1, i2 int) int {
			return i1 + i2
		})
		fmt.Println(result)
		assert.Equal(t, 55, result)
	})

}
