package auth

import (
	"testing"
)

func TestGenerateAdminToken(t *testing.T) {
	token, err := GenerateAdminToken(1)
	if err != nil {
		t.Fatalf("GenerateAdminToken failed: %v", err)
	}
	if token == "" {
		t.Fatal("empty token")
	}
	claims, err := parseToken(token)
	if err != nil {
		t.Fatalf("parseToken failed: %v", err)
	}
	if claims.UserID != 1 || claims.Type != "admin" {
		t.Fatalf("unexpected claims: %+v", claims)
	}
}
