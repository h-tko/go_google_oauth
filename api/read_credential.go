package api

import (
	"encoding/json"
	"io/ioutil"
)

const defaultCredsFile = "./creds.json"

type readCredential interface {
	fromJSON(string) (Credentials, error)
}

type readCredentialImpl struct {
}

func (r *readCredentialImpl) fromJSON(credsFileName string) (Credentials, error) {
	file, err := ioutil.ReadFile(credsFileName)
	if err != nil {
		return Credentials{}, err
	}

	var c Credentials
	json.Unmarshal(file, &c)

	return c, nil
}
