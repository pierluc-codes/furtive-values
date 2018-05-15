package data

type SecretContainer struct {
	Backend    string `json:"b"`
	Project    string `json:"p"`
	Location   string `json:"l"`
	KeyRing    string `json:"r"`
	Key        string `json:"k"`
	CipherText string `json:"c"`
}
