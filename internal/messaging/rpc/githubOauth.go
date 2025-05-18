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

func GithubOauth(ctx context.Context, publisherChan *amqp.Channel, delivery amqp.Delivery, code string) {
	env := config.LoadEnv()
	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", env.GithubOauth.ClientId)
	data.Set("client_secret", env.GithubOauth.ClientSecret)
	data.Set("redirect_uri", env.GithubOauth.RedirectUrl)

	tokenReq, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", strings.NewReader(data.Encode()))
	if err != nil {
		helper.OnError(err, "Failed to make tokenReq object!")
	}
	tokenReq.Header.Set("Accept", "application/json")
	tokenClient := http.Client{}
	tokenRes, err := tokenClient.Do(tokenReq)
	if err != nil {
		helper.OnError(err, "Failed to request token from github server!")
		tokenRes.Body.Close()
	}
	tokenBody, err := io.ReadAll(tokenRes.Body)
	if err != nil {
		helper.OnError(err, "Failed to read all github token body!")
		tokenRes.Body.Close()
	}

	var token types.GitHubTokenResponse
	err = json.Unmarshal(tokenBody, &token)
	if err != nil {
		helper.OnError(err, "Failed to parse github token!")
		tokenRes.Body.Close()
	}

	userInfoReq, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		helper.OnError(err, "Failed to make userInfoReq object!")
	}
	userInfoReq.Header.Set("Authorization", "Bearer "+token.AccessToken)
	userInfoClient := http.Client{}
	userInfoRes, err := userInfoClient.Do(userInfoReq)
	if err != nil {
		helper.OnError(err, "Failed to request user info from github server!")
		userInfoRes.Body.Close()
	}

	userEmailReq, err := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	if err != nil {
		helper.OnError(err, "Failed to make userEmailReq object!")
	}
	userEmailReq.Header.Set("Authorization", "Bearer "+token.AccessToken)
	userEmailClient := http.Client{}
	userEmailRes, err := userEmailClient.Do(userEmailReq)
	if err != nil {
		helper.OnError(err, "Failed to request user email from github server!")
		userEmailRes.Body.Close()
	}

	info, err := io.ReadAll(userInfoRes.Body)
	if err != nil {
		helper.OnError(err, "Failed to read github user info message body!")
	}
	var userInfo types.GitHubUserInfo
	if err := json.Unmarshal(info, &userInfo); err != nil {
		helper.OnError(err, "Failed to unmarshal github user info!")
	}

	email, err := io.ReadAll(userEmailRes.Body)
	if err != nil {
		helper.OnError(err, "Failed to read github user email message body!")
	}
	var emails []types.GitHubEmailInfo
	if err := json.Unmarshal(email, &emails); err != nil {
		helper.OnError(err, "Failed to unmarshal github user emails!")
	}

	var primaryEmail types.GitHubEmailInfo
	for _, e := range emails {
		if e.Primary && e.Verified {
			primaryEmail = e
			break
		}
	}

	userData := types.GitHubUserData{
		UserInfo:  userInfo,
		UserEmail: primaryEmail,
	}
	userDataByte, err := json.Marshal(userData)
	if err != nil {
		helper.OnError(err, "Failed to marshal github user data!")
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
			Body:          userDataByte,
		},
	}
	service.PublishMsg(publisherChan, &publishMsgConf)
}
