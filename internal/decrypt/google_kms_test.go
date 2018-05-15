package decrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/plcstpierre/furtive-values/internal/data"
)

func TestBasic(t *testing.T) {
	container := data.SecretContainer{
		Backend:    "gcp",
		Project:    "solid-muse-203901",
		Location:   "global",
		KeyRing:    "staging-keyring",
		Key:        "key-number-one",
		CipherText: "CiQAkax54PlWA/30nKJuGir/85SKili/geIsUBs6JLrhDkYs0DYSTwD/jr1mKaMtmGtXiMS3W7v7vkQPGa5en+KczPHkykhhpoZDqVbn/YCqjW6Adl18olZrVoaFknme3W63aiV7kXT+oxNnazpgKxviaXjLKG0=",
	}

	value, err := GoogleKmsDecrypter{}.Decrypt(container)

	assert.Nil(t, err)
	assert.Equal(t, "5yg5fzTS9F7Dr7B8a19zXMGwRy3nzOSYCdnwmY", value.PlainText)
}

func TestInvalidKey(t *testing.T) {
	container := data.SecretContainer{
		Backend:    "gcp",
		Project:    "solid-muse-203901",
		Location:   "global",
		KeyRing:    "staging-keyring",
		Key:        "key-number-invalid",
		CipherText: "CiQAkax54PlWA/30nKJuGir/85SKili/geIsUBs6JLrhDkYs0DYSTwD/jr1mKaMtmGtXiMS3W7v7vkQPGa5en+KczPHkykhhpoZDqVbn/YCqjW6Adl18olZrVoaFknme3W63aiV7kXT+oxNnazpgKxviaXjLKG0=",
	}

	value, err := GoogleKmsDecrypter{}.Decrypt(container)

	assert.NotNil(t, err)
	assert.Equal(t, "", value.PlainText)
}

func TestInvalidCipher(t *testing.T) {
	container := data.SecretContainer{
		Backend:    "gcp",
		Project:    "solid-muse-203901",
		Location:   "global",
		KeyRing:    "staging-keyring",
		Key:        "key-number-one",
		CipherText: "invalidCipherText",
	}

	value, err := GoogleKmsDecrypter{}.Decrypt(container)

	assert.NotNil(t, err)
	assert.Equal(t, "", value.PlainText)
}
