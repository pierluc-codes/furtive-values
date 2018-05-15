package codec

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	. "github.com/plcstpierre/furtive-values/internal/data"
)

type Decoder interface {
	Decode(value string) (SecretContainer, error)
}

type DecoderImpl struct {
}

func (DecoderImpl) Decode(value string) (SecretContainer, error) {
	if !strings.HasPrefix(value, "(furtivevalue1:") {
		return SecretContainer{}, errors.New("Invalid prefix")
	}

	if !strings.HasSuffix(value, ")") {
		return SecretContainer{}, errors.New("Invalid suffix")
	}

	splitedValue := strings.Split(value, ":")

	if len(splitedValue) != 2 {
		return SecretContainer{}, errors.New("Invalid payload")
	}

	payloadBase64 := strings.TrimRight(splitedValue[1], ")")

	payloadBytes, err := base64.StdEncoding.DecodeString(payloadBase64)

	if err != nil {
		return SecretContainer{}, err
	}

	var container SecretContainer

	err = json.Unmarshal(payloadBytes, &container)

	if err != nil {
		return SecretContainer{}, err
	}

	return container, nil
}
