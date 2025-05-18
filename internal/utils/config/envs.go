package config

import "os"

type Provider struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

type Env struct {
	RabbitMQUrl string
	GoogleOauth Provider
	GithubOauth Provider
}

func LoadEnv() *Env {
	env := &Env{
		RabbitMQUrl: "amqp://anekra:1234@localhost:5672",
		GoogleOauth: Provider{
			ClientId:     os.Getenv("GOOGLE_OAUTH2_ID"),
			ClientSecret: os.Getenv("GOOGLE_OAUTH2_SECRET"),
			RedirectUrl:  os.Getenv("GOOGLE_REDIRECT_URL"),
		},
		GithubOauth: Provider{
			ClientId:     os.Getenv("GITHUB_OAUTH2_ID"),
			ClientSecret: os.Getenv("GITHUB_OAUTH2_SECRET"),
			RedirectUrl:  os.Getenv("GITHUB_REDIRECT_URL"),
		},
	}

	return env
}
