package decrypt

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/plcstpierre/furtive-values/internal/data"
	"golang.org/x/oauth2/google"
	cloudkms "google.golang.org/api/cloudkms/v1"
)

type GoogleKmsDecrypter struct {
}

const timeout = 2000 * time.Millisecond

func (GoogleKmsDecrypter) Decrypt(container data.SecretContainer) (data.SensitiveContainer, error) {
	backend := strings.TrimSpace(container.Backend)

	if strings.ToLower(backend) != "gcp" {
		return data.SensitiveContainer{}, errors.New("Invalid backend " + backend)
	}

	cipherText := strings.TrimSpace(container.CipherText)

	if len(cipherText) <= 0 {
		return data.SensitiveContainer{}, errors.New("cipherText is empty")
	}

	keyPath, err := createPath(&container)
	if err != nil {
		return data.SensitiveContainer{}, err
	}

	decryptionRequest := &cloudkms.DecryptRequest{
		Ciphertext: cipherText,
	}

	decryptResponse, err := decrypt(keyPath, decryptionRequest)
	if err != nil {
		return data.SensitiveContainer{}, err
	}

	rawPlaintext, err := base64.StdEncoding.DecodeString((*decryptResponse).Plaintext)
	if err != nil {
		return data.SensitiveContainer{}, err
	}

	plainText := strings.TrimSpace(string(rawPlaintext))

	result := data.SensitiveContainer{
		PlainText: plainText,
	}

	return result, nil
}

func createPath(container *data.SecretContainer) (string, error) {
	projectID := strings.TrimSpace(container.Project)
	locationID := strings.TrimSpace(container.Location)
	keyRingID := strings.TrimSpace(container.KeyRing)
	cryptoKeyID := strings.TrimSpace(container.Key)

	if len(projectID) <= 0 {
		return "", errors.New("projectId is missing")
	}

	if len(locationID) <= 0 {
		return "", errors.New("projectId is missing")
	}

	if len(keyRingID) <= 0 {
		return "", errors.New("projectId is missing")
	}

	if len(cryptoKeyID) <= 0 {
		return "", errors.New("projectId is missing")
	}

	path := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s", projectID, locationID, keyRingID, cryptoKeyID)

	return path, nil
}

func decrypt(keyPath string, request *cloudkms.DecryptRequest) (*cloudkms.DecryptResponse, error) {
	decryptionContext, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client, err := google.DefaultClient(decryptionContext, cloudkms.CloudPlatformScope)
	if err != nil {
		return nil, err
	}

	cloudkmsService, err := cloudkms.New(client)
	if err != nil {
		return nil, err
	}

	response, err := cloudkmsService.Projects.Locations.KeyRings.CryptoKeys.Decrypt(keyPath, request).Do()
	if err != nil {
		return nil, err
	}

	return response, nil
}
