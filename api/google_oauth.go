package api

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"
)

const (
	authorizeEndpoint = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenEndpoint     = "https://www.googleapis.com/oauth2/v4/token"
)

type GoogleOAuth struct {
}

type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

type CallbackRequest struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

func (g *GoogleOAuth) Auth() (string, error) {
	creds, err := readCredentialsFromJSON()
	if err != nil {
		return "", err
	}

	conf := g.oauthConfig(creds)

	url := conf.AuthCodeURL("")

	return url, nil
}

func (g *GoogleOAuth) Callback(req *CallbackRequest) error {
	creds, err := readCredentialsFromJSON()
	if err != nil {
		return err
	}

	conf := g.oauthConfig(creds)
	ctx := context.Background()

	tok, err := conf.Exchange(ctx, req.Code)
	if err != nil {
		return err
	}

	if !tok.Valid() {
		return errors.New("valid token.")
	}

	service, _ := v2.New(conf.Client(ctx, tok))
	tokenInfo, _ := service.Tokeninfo().AccessToken(tok.AccessToken).Context(ctx).Do()
	fmt.Printf("%#V", tokenInfo)

	return nil
}

func (g *GoogleOAuth) oauthConfig(c Credentials) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.Cid,
		ClientSecret: c.Csecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
		Scopes: []string{
			"openid",
			"email",
			"profile",
		},
		RedirectURL: "http://localhost:8899/callback",
	}
}
