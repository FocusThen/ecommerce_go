package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Error("expected has to be not empty")
	}

	if hash == "password" {
		t.Error("expected hash to be diffrent from password")
	}
}

func TestComparePasswords(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !ComparePasswords(hash, []byte("password")) {
		t.Error("expected password to match hash")
	}

	if ComparePasswords(hash, []byte("not-password")) {
		t.Error("expected password to not match hash")
	}
}
