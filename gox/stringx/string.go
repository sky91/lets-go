package stringx

import "github.com/sky91/lets-go/gox"

func NewString(s string) *string {
	return &s
}

func Nil2Empty(s *string) string {
	return gox.Nil2Zero(s)
}
