package decrypt

import (
	"github.com/plcstpierre/furtive-values/internal/data"
)

type Decrypter interface {
	Decrypt(container data.SecretContainer) (data.SensitiveContainer, error)
}
