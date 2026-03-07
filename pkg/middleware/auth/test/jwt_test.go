package test

import (
	"Diggpher/pkg/middleware/auth"
	"fmt"
	"testing"
)

func TestGenJwt(t *testing.T) {
	fmt.Println(auth.GenJwt())
}
