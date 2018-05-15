package codec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	assert.Equal(t, true, true, "true")

	container, err := DecoderImpl{}.Decode("(furtivevalue1:eyJiIjoiZ2NwIiwicCI6InNvbGlkLW11c2UtMjAzOTAxIiwibCI6Imdsb2JhbCIsInIiOiJzdGFnaW5nLWtleXJpbmciLCJrIjoia2V5LW51bWJlci1vbmUiLCJjIjoiQ2lRQWtheDU0UGxXQS8zMG5LSnVHaXIvODVTS2lsaS9nZUlzVUJzNkpMcmhEa1lzMERZU1R3RC9qcjFtS2FNdG1HdFhpTVMzVzd2N3ZrUVBHYTVlbitLY3pQSGt5a2hocG9aRHFWYm4vWUNxalc2QWRsMThvbFpyVm9hRmtubWUzVzYzYWlWN2tYVCtveE5uYXpwZ0t4dmlhWGpMS0cwPSJ9)")

	assert.Nil(t, err)
	assert.Equal(t, "gcp", container.Backend)
	assert.Equal(t, "solid-muse-203901", container.Project)
	assert.Equal(t, "staging-keyring", container.KeyRing)
	assert.Equal(t, "key-number-one", container.Key)
	assert.Equal(t, "global", container.Location)
	assert.Equal(t, "CiQAkax54PlWA/30nKJuGir/85SKili/geIsUBs6JLrhDkYs0DYSTwD/jr1mKaMtmGtXiMS3W7v7vkQPGa5en+KczPHkykhhpoZDqVbn/YCqjW6Adl18olZrVoaFknme3W63aiV7kXT+oxNnazpgKxviaXjLKG0=", container.CipherText)
}

func TestInvalidJson(t *testing.T) {
	assert.Equal(t, true, true, "true")

	_, err := DecoderImpl{}.Decode("(furtivevalue1:eyJiIjoiZ2NwIix9)")

	assert.NotNil(t, err)
}
