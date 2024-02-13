package utils

import "sort"

type Stream[T any] []T

func AsStream[T any](slice []T) *Stream[T] {
	return (*Stream[T])(&slice)
}

func (s *Stream[T]) ToSlice() []T {
	return *s
}

func (s *Stream[T]) Map(mapper func(T) T) *Stream[T] {
	result := make(Stream[T], len(*s))
	for i, v := range *s {
		result[i] = mapper(v)
	}
	return &result
}

func (s *Stream[T]) Filter(predicate func(T) bool) *Stream[T] {
	var result Stream[T]
	for _, v := range *s {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return &result
}

func (s *Stream[T]) ForEach(action func(T)) {
	for _, v := range *s {
		action(v)
	}
}

func (s *Stream[T]) Take(n int) *Stream[T] {
	var result Stream[T]
	for i, v := range *s {
		if i == n {
			break
		}
		result = append(result, v)
	}
	return &result
}

func (s *Stream[T]) Find(predicate func(T) bool) *T {
	for _, v := range *s {
		if predicate(v) {
			return &v
		}
	}
	return nil
}

func (s *Stream[T]) Reduce(predicate func(T, T) T) T {
	var result T
	for _, v := range *s {
		result = predicate(result, v)
	}
	return result
}

func (s *Stream[T]) Sort(less func(T, T) bool) *Stream[T] {
	result := make(Stream[T], len(*s))
	copy(result, *s)
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return &result
}
