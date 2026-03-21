// Copyright 2026 Sean (MovieStoreGuy) Marciniak
// SPDX-License-Identifier: MIT

package specify

type TestableCondition struct {
	result bool
	err    error
	called int
}
