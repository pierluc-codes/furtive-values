package text

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

const encryptedValuePattern = `\(furtivevalue1:((\w|=)+)\)`

var encryptedValuePatternRegex = regexp.MustCompile(encryptedValuePattern)

func extractEncryptedValues(line string) []string {
	matching := encryptedValuePatternRegex.FindAllString(line, -1)

	if matching == nil {
		return []string{}
	}

	return matching
}

func processLine(line string, output io.Writer) error {
	var resultingLine = line

	for _, encryptedValue := range extractEncryptedValues(line) {
		decryptedValue, err := DecryptSecretValue(encryptedValue)

		if err != nil {
			return err
		}

		resultingLine = strings.Replace(line, encryptedValue, decryptedValue, -1)
	}

	fmt.Fprintln(output, resultingLine)

	return nil
}

func ProcessStream(input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		err := processLine(line, output)

		if err != nil {
			return err
		}

		if err := scanner.Err(); err != nil {
			return err
		}
	}

	return nil
}
