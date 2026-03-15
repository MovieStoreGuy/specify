// Copyright 2026 Sean (MovieStoreGuy) Marciniak
// SPDX-License-Identifier: MIT

package specify

// Condition provides an abstraction to define conditions
// programatically that will be lazily evaluated once
// once it is checked.
type Condition[T any] interface {
	// Check evaluates the defined condition against the provided value [T],
	// in the event an error is returned, the result value should be ignored.
	Check(input T) (result bool, err error)

	And(Condition[T]) Condition[T]
	Or(Condition[T]) Condition[T]
	Not() Condition[T]
	XOr(Condition[T]) Condition[T]
}

// ConditionFunc implements the Condition[T] type
// and allows for simplified lambda abstractions.
type ConditionFunc[T any] func(input T) (bool, error)

var (
	_ Condition[any] = (*ConditionFunc[any])(nil)
)

func (cond ConditionFunc[T]) Check(input T) (bool, error) {
	return cond(input)
}

func (cond ConditionFunc[T]) And(other Condition[T]) Condition[T] {
	return ConditionFunc[T](func(input T) (bool, error) {
		if result, err := cond(input); err != nil || !result {
			return result, err
		}
		return other.Check(input)
	})
}

func (cond ConditionFunc[T]) Or(other Condition[T]) Condition[T] {
	return ConditionFunc[T](func(input T) (bool, error) {
		if result, err := cond(input); err != nil || result {
			return result, err
		}
		return other.Check(input)
	})
}

func (cond ConditionFunc[T]) Not() Condition[T] {
	return ConditionFunc[T](func(input T) (bool, error) {
		result, err := cond(input)
		result = !result
		return result, err
	})
}

func (cond ConditionFunc[T]) XOr(other Condition[T]) Condition[T] {
	return ConditionFunc[T](func(input T) (bool, error) {
		a, err := cond(input)
		if err != nil {
			return a, err
		}
		b, err := other.Check(input)
		if err != nil {
			return b, err
		}
		return a != b, nil
	})
}
