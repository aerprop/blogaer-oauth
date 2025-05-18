package types

import amqp "github.com/rabbitmq/amqp091-go"

type Channels struct {
	PublisherChan *amqp.Channel
	ConsumerChan  *amqp.Channel
}

type Message struct {
	Code string `json:"code"`
}

type GoogleOAuthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	IDToken      string `json:"id_token"`
}

type GitHubTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type GitHubUserInfo struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatar_url"`
	Name      string `json:"name"`
}

type GitHubEmailInfo struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
	Visibility *string `json:"visibility"`
}

type GitHubUserData struct {
	UserInfo  GitHubUserInfo   `json:"userInfo"`
	UserEmail GitHubEmailInfo  `json:"userEmail"`
}