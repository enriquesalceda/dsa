package algorithms_test

import (
	"dsa/algorithms"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	t.Run("given the array {-10, -8, -6, -4, -2, 0, 2, 4, 6, 8, 10}", func(t *testing.T) {
		t.Run("searching for 2 it returns the index 6", func(t *testing.T) {
			index, err := algorithms.BinarySearch([]int{-10, -8, -6, -4, -2, 0, 2, 4, 6, 8, 10}, 2)
			require.NoError(t, err)
			require.Equal(t, 6, index)
		})

		t.Run("returns error 'not found' when we search for the value 20", func(t *testing.T) {
			index, err := algorithms.BinarySearch([]int{-10, -8, -6, -4, -2, 0, 2, 4, 6, 8, 10}, 20)
			require.Error(t, errors.New("not found"), err)
			require.Equal(t, -1, index)
		})
	})
}
