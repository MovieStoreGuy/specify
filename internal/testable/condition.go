// Copyright 2026 Sean (MovieStoreGuy) Marciniak
// SPDX-License-Identifier: MIT

package testable

type Mock[T any] struct {
	result bool
	err    error
	called int
}

func WithTrueResult[T any]() func(*Mock[T]) {
	return func(m *Mock[T]) {
		m.result = true
	}
}

func WithErrorResult[T any](err error) func(*Mock[T]) {
	return func(m *Mock[T]) {
		m.err = err
	}
}

func NewMock[T any](opts ...func(*Mock[T])) *Mock[T] {
	m := &Mock[T]{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (mock *Mock[T]) Check(_ T) (bool, error) {
	mock.called++
	return mock.result, mock.err
}

func (mock *Mock[T]) Called() int {
	return mock.called
}
