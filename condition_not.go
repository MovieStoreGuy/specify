// Copyright 2026 Sean (MovieStoreGuy) Marciniak
// SPDX-License-Identifier: MIT

package specify

// Not performs the logical not and inverts the result.
func Not[T any](cond Condition[T]) Condition[T] {
	return ConditionFunc[T](func(item T) (bool, error) {
		val, err := cond.Check(item)
		return err == nil && !val, err
	})
}
