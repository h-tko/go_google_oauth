package api

import (
	"encoding/json"
	"io/ioutil"
)

func readCredentialsFromJSON() (Credentials, error) {
	file, err := ioutil.ReadFile("./creds.json")
	if err != nil {
		return Credentials{}, err
	}

	var c Credentials
	json.Unmarshal(file, &c)

	return c, nil
}
