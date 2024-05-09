package linkedlist_test

import (
	"dsa/lists/linkedlist"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppend(t *testing.T) {
	t.Run("When empty list", func(t *testing.T) {
		t.Run("Creates and inserts item at the first item", func(t *testing.T) {
			list := linkedlist.NewEmpty[int]()

			list.Append(1)

			require.Equal(t, 1, list.Size())
		})

		t.Run("When there is only one item in the list", func(t *testing.T) {
			t.Run("Creates and inserts item into the last node of the list as the second item", func(t *testing.T) {
				list := linkedlist.NewEmpty[int]()

				list.Append(1).Append(2)

				require.Equal(t, 2, list.Size())
				require.Equal(t, 1, list.First().Item)
				require.Equal(t, 2, list.First().Next().Item)
			})
		})

		t.Run("When there are two items in the list", func(t *testing.T) {
			t.Run("Creates and inserts item into the last node of the list as the third item", func(t *testing.T) {
				list := linkedlist.NewEmpty[int]()

				list.Append(1).Append(2).Append(3).Append(4)

				require.Equal(t, 4, list.Size())
				require.Equal(t, 1, list.First().Item)
				require.Equal(t, 2, list.First().Next().Item)
				require.Equal(t, 3, list.First().Next().Next().Item)
				require.Equal(t, 4, list.First().Next().Next().Next().Item)
				require.Equal(t, []int{1, 2, 3, 4}, list.Items())
			})
		})
	})
}

func TestFirst(t *testing.T) {
	t.Run("Returns the first node in the list", func(t *testing.T) {
		list := linkedlist.NewWithItems(1)

		require.Equal(t, &linkedlist.Node[int]{Item: 1}, list.First())
	})
}

func TestItems(t *testing.T) {
	t.Run("Returns a slice of all the items in the list", func(t *testing.T) {
		list := linkedlist.NewWithItems(1, 2, 3, 4)

		items := list.Items()

		require.Equal(t, []int{1, 2, 3, 4}, items)
	})
}

func TestSize(t *testing.T) {
	t.Run("Returns list size", func(t *testing.T) {
		list := linkedlist.NewEmpty[int]()
		expectedSize := 100

		for i := 0; i < expectedSize; i++ {
			list.Append(i)
		}

		require.Equal(t, expectedSize, list.Size())
	})
}

func TestIndexOf(t *testing.T) {
	t.Run("Returns the node position containing item in the list", func(t *testing.T) {
		list := linkedlist.NewWithItems(1, 2, 3, 4)

		found, index := list.IndexOf(2)

		require.True(t, found)
		require.Equal(t, 2, index)
	})
}

func TestNewWithItems(t *testing.T) {
	t.Run("Creates a new linked list with the given items", func(t *testing.T) {
		list := linkedlist.NewWithItems(1, 2, 3, 4)

		items := list.Items()

		require.Equal(t, []int{1, 2, 3, 4}, items)
	})
}

func TestInsert(t *testing.T) {
	t.Run("Creates and inserts item in the ith node of the list", func(t *testing.T) {
		t.Run("at the beginning", func(t *testing.T) {
			list := linkedlist.NewWithItems(0, 2, 3, 4, 5)

			list.InsertAt(1, 1)

			require.Equal(t, []int{1, 2, 3, 4, 5}, list.Items())
		})

		t.Run("in the middle", func(t *testing.T) {
			list := linkedlist.NewWithItems(1, 2, 4, 5, 6)

			list.InsertAt(3, 3)

			require.Equal(t, []int{1, 2, 3, 4, 5, 6}, list.Items())
		})
	})

	t.Run("Panics when desired index does not exist", func(t *testing.T) {
		list := linkedlist.NewWithItems(1, 2, 3)

		require.Panics(
			t,
			func() {
				list.InsertAt(5, 5)
			},
			"linked queue does not have this index",
		)
	})

	t.Run("Panics when desired index is zero", func(t *testing.T) {
		list := linkedlist.NewWithItems(1, 2, 3)

		require.Panics(
			t,
			func() {
				list.InsertAt(0, 1)
			},
			"linked queue index starts from one (1)",
		)
	})
}
