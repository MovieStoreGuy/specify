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
}

// ConditionAnd is an extension to condition to allow for natural
// reading condition statements when using `ConditionFunc`.
type ConditionAnd[T any] interface {
	Condition[T]

	// And is an alias for `specify.And(this, b, others...)`
	And(b Condition[T], others ...Condition[T]) Condition[T]
}

// ConditionOr is an extension to condition to allow for natural
// reading condition statements when using `ConditionFunc`.
type ConditionOr[T any] interface {
	Condition[T]

	// Or is an alias for `specify.Or(this, b, others ...)`
	Or(b Condition[T], others ...Condition[T]) Condition[T]
}

// ConditionXor is an extension to condition to allow for natural
// reading condition statements when using `ConditionFunc`.
type ConditionXor[T any] interface {
	Condition[T]

	// Xor is an alias for `specify.Xor(this, b)`
	Xor(b Condition[T]) Condition[T]
}

// ConditionFunc implements the Condition[T] type
// and allows for simplified lambda abstractions.
type ConditionFunc[T any] func(input T) (bool, error)

var (
	_ Condition[any]    = (*ConditionFunc[any])(nil)
	_ ConditionAnd[any] = (*ConditionFunc[any])(nil)
	_ ConditionOr[any]  = (*ConditionFunc[any])(nil)
	_ ConditionXor[any] = (*ConditionFunc[any])(nil)
)

func NewCondition[T any](check func(input T) (bool, error)) Condition[T] {
	return ConditionFunc[T](check)
}

func (cond ConditionFunc[T]) Check(input T) (bool, error) {
	return cond(input)
}

func (cond ConditionFunc[T]) And(b Condition[T], others ...Condition[T]) Condition[T] {
	return And(cond, b, others...)
}

func (cond ConditionFunc[T]) Or(b Condition[T], others ...Condition[T]) Condition[T] {
	return Or(cond, b, others...)
}

func (cond ConditionFunc[T]) Xor(b Condition[T]) Condition[T] {
	return Xor(cond, b)
}
