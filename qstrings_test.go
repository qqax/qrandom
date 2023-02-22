package qrandom

import (
	"testing"
)

const MySecret = "MySecretMySecret"

func TestEncrypt(t *testing.T) {
	text := "Hello, world."
	want := "W7IVX0vS3XIO283b2w=="

	got, err := Encrypt(text, MySecret)
	if err != nil {
		t.Error(err)
		return
	}
	if got != want {
		t.Errorf("Encrypt(%q, %q) = %q, want %q", text, MySecret, got, want)
	}
}

func TestDecrypt(t *testing.T) {
	encryptedString := "W7IVX0vS3XIO283b2w=="
	want := "Hello, world."

	got, err := Decrypt(encryptedString, MySecret)
	if err != nil {
		t.Error(err)
		return
	}
	if got != want {
		t.Errorf("Encrypt(%q, %q) = %q, want %q", encryptedString, MySecret, got, want)
	}
}
