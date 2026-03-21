// Copyright 2026 Sean (MovieStoreGuy) Marciniak
// SPDX-License-Identifier: MIT

package specify

// Or lazily performs the logical or which returns one the first true result.
func Or[T any](a, b Condition[T], conditions ...Condition[T]) Condition[T] {
	return ConditionFunc[T](func(item T) (bool, error) {
		for _, cond := range append([]Condition[T]{a, b}, conditions...) {
			val, err := cond.Check(item)
			if err != nil || val {
				return val, err
			}
		}
		return false, nil
	})
}
