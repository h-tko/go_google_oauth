package api

import (
	v2 "google.golang.org/api/oauth2/v2"
)

type Auth interface {
	Auth() (string, error)
	Callback(*CallbackRequest) (*v2.Tokeninfo, error)
}

func NewAuthAPI() Auth {
	return NewGoogleOAuth(&readCredentialImpl{})
}
