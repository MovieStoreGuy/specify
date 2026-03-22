// Copyright 2026 Sean (MovieStoreGuy) Marciniak
// SPDX-License-Identifier: MIT

package specify

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAction(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		action Action[any]
		expect string
	}{
		{
			name: "successful action",
			action: NewAction(func(_ any) error {
				return nil
			}),
			expect: "",
		},
		{
			name: "failed action",
			action: NewAction(func(_ any) error {
				return errors.New("faulty")
			}),
			expect: "faulty",
		},
		{
			name: "unexecuted action",
			action: NewConditionalAction(
				NewCondition(func(_ any) (bool, error) {
					return false, nil
				}),
				NewAction(func(_ any) error {
					return errors.New("unexpected call")
				}),
			),
			expect: "",
		},
		{
			name: "condition error",
			action: NewConditionalAction(
				NewCondition(func(_ any) (bool, error) {
					return true, errors.New("condition failed")
				}),
				NewAction(func(_ any) error {
					return errors.New("unexpected call")
				}),
			),
			expect: "condition failed",
		},
		{
			name: "expected action",
			action: NewConditionalAction(
				NewCondition(func(_ any) (bool, error) {
					return true, nil
				}),
				NewAction(func(_ any) error {
					return errors.New("intentional fail")
				}),
			),
			expect: "intentional fail",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := tc.action.Do(nil)
			if tc.expect != "" {
				assert.EqualError(t, err, tc.expect, "Must match the expected error")
			} else {
				assert.NoError(t, err, "Must not return an error")
			}
		})
	}
}
