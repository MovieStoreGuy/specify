// Copyright 2026 Sean (MovieStoreGuy) Marciniak
// SPDX-License-Identifier: MIT

package testable

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMock(t *testing.T) {
	t.Parallel()

	boom := errors.New("boom")

	for _, tc := range []struct {
		name   string
		opts   []func(*Mock[any])
		expect *Mock[any]
	}{
		{
			name:   "no options",
			opts:   nil,
			expect: &Mock[any]{},
		},
		{
			name: "valid results",
			opts: []func(*Mock[any]){
				WithTrueResult[any](),
			},
			expect: &Mock[any]{
				result: true,
			},
		},
		{
			name: "error result",
			opts: []func(*Mock[any]){
				WithErrorResult[any](boom),
			},
			expect: &Mock[any]{
				err: boom,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := NewMock(tc.opts...)
			assert.Equal(
				t,
				tc.expect,
				actual,
				"Must match to expect value",
			)
		})
	}
}
