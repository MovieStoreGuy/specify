// Copyright 2026 Sean (MovieStoreGuy) Marciniak
// SPDX-License-Identifier: MIT

package specify

// Action is used to abstract methods so that
// it can be used in conjuction with conditions.
type Action[T any] interface {
	// Do will call the underlying method
	// returning any error encountered.
	Do(input T) error
}

type action[T any] func(input T) error

func (fn action[T]) Do(input T) error {
	return fn(input)
}

var (
	_ Action[any] = (*action[any])(nil)
)

// NewActions allows for wrapping a function as an `Action[T]`
func NewAction[T any](act func(input T) error) Action[T] {
	return action[T](act)
}

// NewConditionalAction wraps an `Action[T]` and will on execute the action
// if the condition is true.
func NewConditionalAction[T any](cond Condition[T], act Action[T]) Action[T] {
	return NewAction(func(input T) error {
		if ok, err := cond.Check(input); err != nil || !ok {
			return err
		}
		return act.Do(input)
	})
}
