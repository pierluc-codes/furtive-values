package text

import (
	"github.com/plcstpierre/furtive-values/internal/codec"
	"github.com/plcstpierre/furtive-values/internal/decrypt"
)

func DecryptSecretValue(secret string) (string, error) {
	decoder := codec.DecoderImpl{}
	secretContainer, err := decoder.Decode(secret)

	if err != nil {
		return "", err
	}

	sensitiveContainer, err := decrypt.GoogleKmsDecrypter{}.Decrypt(secretContainer)

	if err != nil {
		return "", err
	}

	return sensitiveContainer.PlainText, nil
}
