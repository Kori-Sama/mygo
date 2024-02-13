package test

import (
	"mygo/pkg/utils"
	"reflect"
	"testing"
)

func TestStream(t *testing.T) {
	t.Run("chain", func(t *testing.T) {
		arr := []int{3, -1, 1, 5, 2, 6, 8, 2, 1, 4, 6, -12, 22, -2, 7}

		got := utils.AsStream[int](arr).Map(func(v int) int {
			return v * 2
		}).Filter(func(v int) bool {
			return v > 0
		}).Sort(func(i, j int) bool {
			return i < j
		}).Take(5).ToSlice()

		want := []int{2, 2, 4, 4, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Map", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		result := utils.AsStream[int](arr).Map(func(v int) int {
			return v * 2
		})
		want := []int{2, 4, 6, 8, 10}
		got := result.ToSlice()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Filter", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		result := utils.AsStream[int](arr).Filter(func(v int) bool {
			return v%2 == 0
		})
		want := []int{2, 4}
		got := result.ToSlice()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)

		}
	})

	t.Run("Reduce", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		result := utils.AsStream[int](arr).Reduce(func(acc int, v int) int {
			return acc + v
		})
		want := 15
		if result != want {
			t.Errorf("got %v, want %v", result, want)
		}
	})

	t.Run("Find", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		result := utils.AsStream[int](arr).Find(func(v int) bool {
			return v == 3
		})
		want := 3
		if *result != want {
			t.Errorf("got %v, want %v", *result, want)
		}
	})

	t.Run("Take", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		got := utils.AsStream[int](arr).Take(3).ToSlice()
		want := []int{1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Sort", func(t *testing.T) {
		arr := []int{3, 1, 2, 5, 4}
		got := utils.AsStream[int](arr).Sort(func(a, b int) bool { return a < b }).ToSlice()
		want := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
