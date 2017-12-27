package api

type Auth interface {
	Auth() (string, error)
}

func NewAuthAPI() Auth {
	return &GoogleOAuth{}
}
