package auth

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token := GenerateToken("tester")
	spew.Dump(token)
	panic("tk")
}
