package linkedlist_test

import (
	"dsa/lists/linkedlist"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAppend(t *testing.T) {
	t.Run("When empty list", func(t *testing.T) {
		t.Run("Creates and inserts item at the first item", func(t *testing.T) {
			list := linkedlist.New[int]()

			list.Append(1)

			require.Equal(t, 1, list.Size())
		})

		t.Run("When there is only one item in the list", func(t *testing.T) {
			t.Run("Creates and inserts item into the last node of the list as the second item", func(t *testing.T) {
				list := linkedlist.New[int]()

				list.Append(1).Append(2)

				require.Equal(t, 2, list.Size())
				require.Equal(t, 1, list.First().Item)
				require.Equal(t, 2, list.First().Next().Item)
			})
		})

		t.Run("When there are two items in the list", func(t *testing.T) {
			t.Run("Creates and inserts item into the last node of the list as the third item", func(t *testing.T) {
				list := linkedlist.New[int]()

				list.Append(1).Append(2).Append(3).Append(4)

				require.Equal(t, 4, list.Size())
				require.Equal(t, 1, list.First().Item)
				require.Equal(t, 2, list.First().Next().Item)
				require.Equal(t, 3, list.First().Next().Next().Item)
				require.Equal(t, 4, list.First().Next().Next().Next().Item)
			})
		})
	})
}

func TestFirst(t *testing.T) {
	t.Run("Returns the first node in the list", func(t *testing.T) {
		list := linkedlist.New[int]()

		list.Append(1)

		require.Equal(t, &linkedlist.Node[int]{Item: 1}, list.First())
	})
}

func TestItems(t *testing.T) {
	t.Run("Returns a slice of all the items in the list", func(t *testing.T) {
		list := linkedlist.New[int]()
		list.Append(1).Append(2).Append(3).Append(4)

		items := list.Items()

		require.Equal(t, []int{1, 2, 3, 4}, items)
	})
}

func TestSize(t *testing.T) {
	t.Run("Returns list size", func(t *testing.T) {
		list := linkedlist.New[int]()
		expectedSize := 100

		for i := 0; i < expectedSize; i++ {
			list.Append(i)
		}

		require.Equal(t, expectedSize, list.Size())
	})
}

func TestIndexOf(t *testing.T) {
	t.Run("Returns the node position containing item in the list", func(t *testing.T) {
		list := linkedlist.New[int]()
		list.Append(1).Append(2).Append(3).Append(4)

		found, index := list.IndexOf(2)

		require.True(t, found)
		require.Equal(t, 2, index)
	})
}
