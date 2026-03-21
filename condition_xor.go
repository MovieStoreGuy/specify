// Copyright 2026 Sean (MovieStoreGuy) Marciniak
// SPDX-License-Identifier: MIT

package specify

// Xor performs the logical xor which only returns true if one condition is true
func Xor[T any](a, b Condition[T], more ...Condition[T]) Condition[T] {
	return ConditionFunc[T](func(item T) (bool, error) {
		previous := false
		for _, cond := range append([]Condition[T]{a, b}, more...) {
			current, err := cond.Check(item)
			if err != nil || (previous && current) {
				return false, err
			}
			previous = previous != current
		}
		return previous, nil
	})
}
