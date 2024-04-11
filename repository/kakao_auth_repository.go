package repository

import (
	"encoding/json"
	"github.com/cho8833/CC-Calendar/pkg/kakao_auth_wrapper/config"
	"github.com/cho8833/CC-Calendar/pkg/kakao_auth_wrapper/dto"
	"github.com/cho8833/CC-Calendar/pkg/kakao_auth_wrapper/util"
	"log"
	"net/http"
)

type KakaoClient interface {
	Do(req *http.Request) (*http.Response, error)
}
type KaKaoAuthRepository struct {
	client KakaoClient
	config *config.KakaoAuthConfig
}

func NewKakaoAuthRepository(config *config.KakaoAuthConfig, client *KakaoClient) *KaKaoAuthRepository {
	return &KaKaoAuthRepository{client: *client, config: config}
}

func (repository *KaKaoAuthRepository) GetUserInfo(accessToken string) (*dto.KakaoUserInfoRes, error) {
	// Make Http Request
	req, err := http.NewRequest("GET", "https://kapi.kakao.com/v2/user/me", nil)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	req.Header.Add("Authorization", "Bearer "+accessToken)
	q := req.URL.Query()
	q.Add("secure_resource", "true")
	req.URL.RawQuery = q.Encode()

	// Request
	res, err := repository.client.Do(req)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	// Handle Response
	bodyByte, err := util.HttpResponseHandler(res)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	// Parse Body
	data := dto.KakaoUserInfoRes{}
	err = json.Unmarshal(bodyByte, &data)
	if err != nil {
		log.Printf(err.Error())
		panic(err)
	}

	return &data, nil
}

func (repository *KaKaoAuthRepository) RequestToken(code string) (*dto.TokenRes, error) {

	// Make Http Request
	req, err := http.NewRequest("POST", "https://kauth.kakao.com/oauth/token", nil)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	req.Header.Add("Content-type", "application/x-www-form-urlencoded;charset=utf-8")
	q := req.URL.Query()
	q.Add("grant_type", "authorization_code")
	q.Add("client_id", repository.config.ClientId)
	q.Add("redirect_url", repository.config.RedirectUrl)
	q.Add("code", code)
	req.URL.RawQuery = q.Encode()

	// Request
	res, err := repository.client.Do(req)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	// Handle Response
	bodyByte, err := util.HttpResponseHandler(res)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	// Parse Body
	data := dto.TokenRes{}
	err = json.Unmarshal(bodyByte, &data)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	return &data, nil
}
