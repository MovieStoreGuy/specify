// Copyright 2026 Sean (MovieStoreGuy) Marciniak
// SPDX-License-Identifier: MIT

package specify

import (
	"errors"
	"testing"

	"github.com/MovieStoreGuy/specify/internal/testable"
	"github.com/stretchr/testify/assert"
)

func TestConditionOr(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		A, B   Condition[any]
		expect bool
		errVal string
	}{
		{
			name:   "none set",
			A:      testable.NewMock[any](),
			B:      testable.NewMock[any](),
			expect: false,
			errVal: "",
		},
		{
			name: "A set",
			A: testable.NewMock(
				testable.WithTrueResult[any](),
			),
			B:      testable.NewMock[any](),
			expect: true,
			errVal: "",
		},
		{
			name: "B set",
			A:    testable.NewMock[any](),
			B: testable.NewMock[any](
				testable.WithTrueResult[any](),
			),
			expect: true,
			errVal: "",
		},
		{
			name: "A & B set",
			A: testable.NewMock[any](
				testable.WithTrueResult[any](),
			),
			B: testable.NewMock[any](
				testable.WithTrueResult[any](),
			),
			expect: true,
			errVal: "",
		},
		{
			name: "A error",
			A: testable.NewMock(
				testable.WithErrorResult[any](errors.New("boom")),
			),
			B:      testable.NewMock[any](),
			expect: false,
			errVal: "boom",
		},
		{
			name: "B error",
			A:    testable.NewMock[any](),
			B: testable.NewMock[any](
				testable.WithErrorResult[any](errors.New("boom")),
			),
			expect: false,
			errVal: "boom",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual, err := Or(tc.A, tc.B).Check(nil)
			assert.Equal(t, tc.expect, actual, "Must match the expected result")
			if tc.errVal != "" {
				assert.EqualError(t, err, tc.errVal, "Must match the expected error")
			} else {
				assert.NoError(t, err, "Must not return any error value")
			}
		})
	}
}
