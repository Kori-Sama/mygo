package utils

import "sort"

type Stream[T any] struct {
	data []T
	opts []func(*Stream[T])
}

func AsStream[T any](slice []T) *Stream[T] {
	return &Stream[T]{data: slice}
}

func (s *Stream[T]) ToSlice() []T {
	s.consume()
	return s.data
}

func (s *Stream[T]) Map(mapper func(T) T) *Stream[T] {
	opt := func(s *Stream[T]) {
		for i, v := range s.data {
			s.data[i] = mapper(v)
		}
	}
	s.opts = append(s.opts, opt)
	return s
}

func (s *Stream[T]) Filter(predicate func(T) bool) *Stream[T] {
	opt := func(s *Stream[T]) {
		var result []T
		for _, v := range s.data {
			if predicate(v) {
				result = append(result, v)
			}
		}
		s.data = result
	}
	s.opts = append(s.opts, opt)
	return s
}

func (s *Stream[T]) Take(n int) *Stream[T] {
	opt := func(s *Stream[T]) {
		if n < len(s.data) {
			s.data = s.data[:n]
		}
	}
	s.opts = append(s.opts, opt)
	return s
}

func (s *Stream[T]) Sort(predicate func(T, T) bool) *Stream[T] {
	opt := func(s *Stream[T]) {
		sort.Slice(s.data, func(i, j int) bool {
			return predicate(s.data[i], s.data[j])
		})
	}
	s.opts = append(s.opts, opt)
	return s
}

func (s *Stream[T]) ForEach(action func(T)) {
	s.consume()
	for _, v := range s.data {
		action(v)
	}
}

func (s *Stream[T]) Reduce(reducer func(T, T) T) T {
	s.consume()
	var result T
	for _, v := range s.data {
		result = reducer(result, v)
	}
	return result
}

func (s *Stream[T]) Find(predicate func(T) bool) *T {
	s.consume()
	for _, v := range s.data {
		if predicate(v) {
			return &v
		}
	}
	return nil
}

func (s *Stream[T]) consume() {
	if len(s.opts) > 0 {
		for _, opt := range s.opts {
			opt(s)
		}
	}
}
