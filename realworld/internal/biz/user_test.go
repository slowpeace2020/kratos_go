package biz

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	s := hashPassword("abcd")
	//debug 输出
	spew.Dump(s)
	//panic(1)
}

func TestVerifyPassword(t *testing.T) {
	a := assert.New(t)
	a.True(verifyPassword("$2a$10$LgMWjwYft6NMpvH/7Bbmq.dbfQWLmym6xICwpxL7weBARoF7Z.RNq", "abcd"))
	a.False(verifyPassword("$2a$10$LgMWjwYft6NMpvH/7Bbmq.dbfQWLmym6xICwpxL7weBARoF7Z.RNq", "abcd1"))
}
