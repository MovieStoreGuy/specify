// Copyright 2026 Sean (MovieStoreGuy) Marciniak
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/MovieStoreGuy/specify"
)

type Session struct {
	userId   int
	isBanned bool
	isAdmin  bool
	age      int
}

func NewSession(opts ...func(*Session)) *Session {
	s := &Session{
		userId:   rand.Int(),
		isBanned: false,
		isAdmin:  false,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Session) UserID() int {
	return s.userId
}

func IsAdmin() specify.ConditionFunc[*Session] {
	return specify.ConditionFunc[*Session](func(s *Session) (bool, error) {
		return s.isAdmin && s.age > 18, nil
	})
}

func IsBanned() specify.Condition[*Session] {
	return specify.NewCondition(func(s *Session) (bool, error) {
		return s.isBanned, nil
	})
}

func IsValidAdmin() specify.Condition[*Session] {
	return IsAdmin().And(specify.Not(IsBanned()))
}

func IsValidAdminPimitive(s *Session) (bool, error) {
	return s.isAdmin && !s.isBanned, nil
}

func main() {
	validator := specify.NewConditionalAction(
		IsValidAdmin(),
		specify.NewAction(func(user *Session) error {
			fmt.Printf("Valid: User %x is a current admin\n", user.UserID())
			return nil
		}),
	)

	for _, user := range []*Session{
		NewSession(func(s *Session) {
			s.isBanned = true
		}),
		NewSession(func(s *Session) {
			s.isAdmin = true
			s.age = 24
		}),
		NewSession(func(s *Session) {
			s.isAdmin = true
			s.isBanned = true
		}),
		NewSession(),
	} {
		if err := validator.Do(user); err != nil {
			panic(err)
		}
	}

}
