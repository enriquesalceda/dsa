package linkedlist_test

import (
	"dsa/lists/linkedlist"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAppend(t *testing.T) {
	t.Run("When empty list", func(t *testing.T) {
		t.Run("Creates and inserts item at the first item", func(t *testing.T) {
			list := linkedlist.NewList[int]()

			list.Append(linkedlist.Node[int]{Item: 1})

			require.Equal(t, 1, list.Size())
		})
	})

}
