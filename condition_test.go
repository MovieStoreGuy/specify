// Copyright 2026 Sean (MovieStoreGuy) Marciniak
// SPDX-License-Identifier: MIT

package specify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func StaticValue[T any](tb testing.TB, result bool, called int) Condition[T] {
	tb.Helper()

	count := 0
	tb.Cleanup(func() {
		assert.Equal(tb, called, count, "Must match the expected called count")
	})

	return ConditionFunc[T](func(_ T) (bool, error) {
		count++
		return result, nil
	})
}

func TestConditionFunc_Not(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		cond   func(t testing.TB) Condition[struct{}]
		expect bool
	}{
		{
			name: "static true",
			cond: func(t testing.TB) Condition[struct{}] {
				return StaticValue[struct{}](t, true, 1)
			},
			expect: false,
		},
		{
			name: "static false",
			cond: func(t testing.TB) Condition[struct{}] {
				return StaticValue[struct{}](t, false, 1)
			},
			expect: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.cond(t).Not().Check(struct{}{})
			assert.NoError(t, err, "Must not error when validating invert")
			assert.Equal(t, tc.expect, actual, "Must match the expected value")
		})
	}
}

func TestConditionFunc_And(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		cond   func(t testing.TB) Condition[struct{}]
		expect bool
	}{
		{
			name: "True Table(true, true)",
			cond: func(t testing.TB) Condition[struct{}] {
				return StaticValue[struct{}](t, true, 1).
					And(StaticValue[struct{}](t, true, 1))
			},
			expect: true,
		},
		{
			name: "False Table(true, false)",
			cond: func(t testing.TB) Condition[struct{}] {
				return StaticValue[struct{}](t, true, 1).
					And(StaticValue[struct{}](t, false, 1))
			},
			expect: false,
		},
		{
			name: "False Table(false, true)(lazy evaluation)",
			cond: func(t testing.TB) Condition[struct{}] {
				return StaticValue[struct{}](t, false, 1).
					And(StaticValue[struct{}](t, true, 0))
			},
			expect: false,
		},
		{
			name: "False Table(false, false) (early exit)",
			cond: func(t testing.TB) Condition[struct{}] {
				return StaticValue[struct{}](t, false, 1).
					And(StaticValue[struct{}](t, false, 0))
			},
			expect: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.cond(t).Check(struct{}{})
			assert.NoError(t, err, "Must not error when evaluating condition")
			assert.Equal(t, tc.expect, actual, "Must match the expected result")
		})
	}
}

func TestConditionFunc_Or(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		cond   func(t testing.TB) Condition[struct{}]
		expect bool
	}{
		{
			name: "True Table(true, true)(lazy evaluation)",
			cond: func(t testing.TB) Condition[struct{}] {
				return StaticValue[struct{}](t, true, 1).
					Or(StaticValue[struct{}](t, true, 0))
			},
			expect: true,
		},
		{
			name: "True Table(True, False)(lazy evaluation)",
			cond: func(t testing.TB) Condition[struct{}] {
				return StaticValue[struct{}](t, true, 1).
					Or(StaticValue[struct{}](t, false, 0))
			},
			expect: true,
		},
		{
			name: "True Table(False, True)",
			cond: func(t testing.TB) Condition[struct{}] {
				return StaticValue[struct{}](t, false, 1).
					Or(StaticValue[struct{}](t, true, 1))
			},
			expect: true,
		},
		{
			name: "False Table(False, False)",
			cond: func(t testing.TB) Condition[struct{}] {
				return StaticValue[struct{}](t, false, 1).
					Or(StaticValue[struct{}](t, false, 1))
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.cond(t).Check(struct{}{})
			assert.NoError(t, err, "Must not error when performing check")
			assert.Equal(t, tc.expect, actual, "Must match the expected result")
		})
	}
}

func TestConditionFunc_XOr(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		cond   func(t testing.TB) Condition[struct{}]
		expect bool
	}{
		{
			name: "False Table(True,True)",
			cond: func(t testing.TB) Condition[struct{}] {
				return StaticValue[struct{}](t, true, 1).
					XOr(StaticValue[struct{}](t, true, 1))
			},
			expect: false,
		},
		{
			name: "True Table(True,False)",
			cond: func(t testing.TB) Condition[struct{}] {
				return StaticValue[struct{}](t, true, 1).
					XOr(StaticValue[struct{}](t, false, 1))
			},
			expect: true,
		},
		{
			name: "True Table(False,True)",
			cond: func(t testing.TB) Condition[struct{}] {
				return StaticValue[struct{}](t, false, 1).
					XOr(StaticValue[struct{}](t, true, 1))
			},
			expect: true,
		},
		{
			name: "False Table(False,False)",
			cond: func(t testing.TB) Condition[struct{}] {
				return StaticValue[struct{}](t, false, 1).
					XOr(StaticValue[struct{}](t, false, 1))
			},
			expect: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.cond(t).Check(struct{}{})
			assert.NoError(t, err, "Must not return an error when evaluating")
			assert.Equal(t, tc.expect, actual, "Must match the expected result")
		})
	}
}
