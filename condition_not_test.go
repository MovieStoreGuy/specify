// Copyright 2026 Sean (MovieStoreGuy) Marciniak
// SPDX-License-Identifier: MIT

package specify

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MovieStoreGuy/specify/internal/testable"
)

func TestConditionNot(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		cond   Condition[any]
		expect bool
		errVal string
	}{
		{
			name:   "true result",
			cond:   testable.NewMock[any](),
			expect: true,
			errVal: "",
		},
		{
			name: "false result",
			cond: testable.NewMock[any](
				testable.WithTrueResult[any](),
			),
			expect: false,
			errVal: "",
		},
		{
			name: "error result",
			cond: testable.NewMock(
				testable.WithErrorResult[any](errors.New("boom")),
			),
			expect: false,
			errVal: "boom",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual, err := Not(tc.cond).Check(nil)
			assert.Equal(t, tc.expect, actual, "Must match the expected result")
			if tc.errVal != "" {
				assert.EqualError(t, err, tc.errVal, "Must match the expected error")
			} else {
				assert.NoError(t, err, "Must not error when condition check")
			}
		})
	}
}
