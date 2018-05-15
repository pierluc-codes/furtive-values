package text

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractEncryptedValues(t *testing.T) {
	testNoMatch(t, "nomatch")
	testNoMatch(t, "start(furtivevalue1:valueend")
	testNoMatch(t, "startfurtivevalue1:valueend")
	testNoMatch(t, "start(furtivevalue:value)end")
	testNoMatch(t, "start(furtivevaluevalue)end")

	testValue(t, "start(furtivevalue1:value)end", "(furtivevalue1:value)")
	testCollection(t, "start(furtivevalue1:value1)end(furtivevalue1:value2)", []string{"(furtivevalue1:value1)", "(furtivevalue1:value2)"})
}

func testNoMatch(t *testing.T, value string) {
	testCollection(t, value, []string{})
}

func testValue(t *testing.T, value string, expected string) {
	testCollection(t, value, []string{expected})
}

func testCollection(t *testing.T, value string, expected []string) {
	assert.Equal(t, expected, extractEncryptedValues(value), "Found unexpected match")
}

func TestProcessStream(t *testing.T) {
	input := `line1
	line2
	line3a(furtivevalue1:eyJiIjoiZ2NwIiwicCI6InNvbGlkLW11c2UtMjAzOTAxIiwibCI6Imdsb2JhbCIsInIiOiJzdGFnaW5nLWtleXJpbmciLCJrIjoia2V5LW51bWJlci1vbmUiLCJjIjoiQ2lRQWtheDU0UGxXQS8zMG5LSnVHaXIvODVTS2lsaS9nZUlzVUJzNkpMcmhEa1lzMERZU1R3RC9qcjFtS2FNdG1HdFhpTVMzVzd2N3ZrUVBHYTVlbitLY3pQSGt5a2hocG9aRHFWYm4vWUNxalc2QWRsMThvbFpyVm9hRmtubWUzVzYzYWlWN2tYVCtveE5uYXpwZ0t4dmlhWGpMS0cwPSJ9)line3b
	line4`

	expected := `line1
	line2
	line3a5yg5fzTS9F7Dr7B8a19zXMGwRy3nzOSYCdnwmYline3b
	line4
` // ByteBuffer is adding a final line return

	buffer := new(bytes.Buffer)

	inputReader := strings.NewReader(input)

	ProcessStream(inputReader, buffer)

	actual := buffer.String()

	assert.Equal(t, expected, actual)
}
