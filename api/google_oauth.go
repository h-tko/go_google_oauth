package api

import (
	"context"
	"errors"
	"os"

	"golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"
)

const (
	authorizeEndpoint = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenEndpoint     = "https://www.googleapis.com/oauth2/v4/token"
)

type GoogleOAuth struct {
	reader readCredential
}

type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

type CallbackRequest struct {
	Code  string `query:"code"`
	State string `query:"session_state"`
}

func NewGoogleOAuth(r readCredential) *GoogleOAuth {
	return &GoogleOAuth{
		reader: r,
	}
}

func (g *GoogleOAuth) Auth() (string, error) {
	creds, err := g.reader.fromJSON(defaultCredsFile)
	if err != nil {
		return "", err
	}

	conf, err := g.oauthConfig(creds)
	if err != nil {
		return "", err
	}

	url := conf.AuthCodeURL("")

	return url, nil
}

func (g *GoogleOAuth) Callback(req *CallbackRequest) (*v2.Tokeninfo, error) {
	creds, err := g.reader.fromJSON(defaultCredsFile)
	if err != nil {
		return nil, err
	}

	conf, err := g.oauthConfig(creds)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	tok, err := conf.Exchange(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	if !tok.Valid() {
		return nil, errors.New("valid token.")
	}

	service, err := v2.New(conf.Client(ctx, tok))
	if err != nil {
		return nil, err
	}

	tokenInfo, err := service.Tokeninfo().AccessToken(tok.AccessToken).Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	return tokenInfo, nil
}

func (g *GoogleOAuth) oauthConfig(c Credentials) (*oauth2.Config, error) {
	url := os.Getenv("CALLBACK_URL")
	if url == "" {
		return nil, errors.New(".envにCALLBACK_URLが設定されていません。")
	}

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
		RedirectURL: url,
	}, nil
}
