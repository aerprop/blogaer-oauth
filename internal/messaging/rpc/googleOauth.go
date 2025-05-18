package rpc

import (
	"blogaer-oauth/internal/service"
	"blogaer-oauth/internal/utils/config"
	"blogaer-oauth/internal/utils/helper"
	"blogaer-oauth/internal/utils/types"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
)

func GoogleOauth(ctx context.Context, publisherChan *amqp.Channel, delivery amqp.Delivery, code string) {
	env := config.LoadEnv()

	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", env.GoogleOauth.ClientId)
	data.Set("client_secret", env.GoogleOauth.ClientSecret)
	data.Set("redirect_uri", env.GoogleOauth.RedirectUrl)
	data.Set("grant_type", "authorization_code")

	tokenReq, err := http.NewRequest(
		"POST",
		"https://oauth2.googleapis.com/token",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		helper.OnError(err, "Failed to make request object!")
	}
	tokenReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	tokenClient := http.Client{}
	tokenRes, err := tokenClient.Do(tokenReq)
	if err != nil {
		helper.OnError(err, "Failed to request token from google server!")
		tokenRes.Body.Close()
	}
	tokenBody, err := io.ReadAll(tokenRes.Body)
	if err != nil {
		helper.OnError(err, "Failed to read all google token body!")
		tokenRes.Body.Close()
	}

	var token types.GoogleOAuthTokenResponse
	if err := json.Unmarshal(tokenBody, &token); err != nil {
		helper.OnError(err, "Failed to parse google token!")
	}

	userInfoReq, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo", nil)
	if err != nil {
		helper.OnError(err, "Failed to make request object!")
	}
	userInfoReq.Header.Set("Authorization", "Bearer "+token.AccessToken)
	userInfoClient := http.Client{}
	userInfoRes, err := userInfoClient.Do(userInfoReq)
	if err != nil {
		helper.OnError(err, "Failed to request user info from google server!")
		userInfoRes.Body.Close()
	}

	userInfo, err := io.ReadAll(userInfoRes.Body)
	if err != nil {
		helper.OnError(err, "Failed to read message body!")
	}

	publishMsgConf := service.PublishMsgParams{
		Ctx:       ctx,
		Exchange:  "",
		Key:       delivery.ReplyTo,
		Mandatory: false,
		Immediate: false,
		Msg: amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: delivery.CorrelationId,
			Body:          userInfo,
		},
	}
	service.PublishMsg(publisherChan, &publishMsgConf)
}
