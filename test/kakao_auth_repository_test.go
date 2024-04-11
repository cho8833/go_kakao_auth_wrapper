package test

import (
	"bytes"
	"github.com/cho8833/CC-Calendar/pkg/kakao_auth_wrapper/config"
	"github.com/cho8833/CC-Calendar/pkg/kakao_auth_wrapper/repository"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

const successGetUserInfoBody string = "{\n  \"id\": 123456789,\n  \"kakao_account\": {\n    \"profile_needs_agreement\": false,\n    \"profile\": {\n      \"nickname\": \"홍길동\",\n      \"thumbnail_image_url\": \"http://yyy.kakao.com/.../img_110x110.jpg\",\n      \"profile_image_url\": \"http://yyy.kakao.com/dn/.../img_640x640.jpg\",\n      \"is_default_image\": false,\n      \"is_default_nickname\": false\n    },\n    \"email_needs_agreement\": false,\n    \"is_email_valid\": true,\n    \"is_email_verified\": true,\n    \"email\": \"sample@sample.com\",\n    \"name_needs_agreement\": false,\n    \"name\": \"홍길동\",\n    \"age_range_needs_agreement\": false,\n    \"age_range\": \"20~29\",\n    \"birthday_needs_agreement\": false,\n    \"birthday\": \"1130\",\n    \"gender_needs_agreement\": false,\n    \"gender\": \"female\"\n  },\n  \"properties\": {\n    \"nickname\": \"홍길동카톡\",\n    \"thumbnail_image\": \"http://xxx.kakao.co.kr/.../aaa.jpg\",\n    \"profile_image\": \"http://xxx.kakao.co.kr/.../bbb.jpg\",\n    \"custom_field1\": \"23\",\n    \"custom_field2\": \"여\"\n  }\n}"
const FailGetKakaoUser string = "\"{\\n  \\\"code\\\": -101,\\n  \\\"msg\\\": \\\"NotRegisteredUserException\\\"\\n}\""
const successReqTokenBody = "{\n    \"access_token\": \"B9OXI8MDF5KYafZBzTnwxHYGAG9IOPca-dUKKiVSAAABjsrcjW0h5oEAb4_jFQ\",\n    \"token_type\": \"bearer\",\n    \"refresh_token\": \"aPp0jOC8is5OwSKx3up3GzjiSgtcNxrqX9gKKiVSAAABjsrcjWoh5oEAb4_jFQ\",\n    \"expires_in\": 21599,\n    \"scope\": \"profile_nickname\",\n    \"refresh_token_expires_in\": 5183999\n}"

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (c *MockClient) Do(req *http.Request) (*http.Response, error) {
	return c.DoFunc(req)
}

func Test_getUserInfo(t *testing.T) {

	t.Run("parse body if success", func(t *testing.T) {
		// given
		var mockClient repository.KakaoClient = &MockClient{DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(successGetUserInfoBody))}, nil
		}}
		kakaoRepository := repository.NewKakaoAuthRepository(&config.KakaoAuthConfig{}, &mockClient)

		// when
		_, err := kakaoRepository.GetUserInfo("test")

		// then
		assert.NoError(t, err)
	})

	t.Run("parse error if response error", func(t *testing.T) {
		// given
		var mockClient repository.KakaoClient = &MockClient{DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 401, Body: io.NopCloser(bytes.NewBufferString(FailGetKakaoUser))}, nil
		}}
		kakaoRepository := repository.NewKakaoAuthRepository(&config.KakaoAuthConfig{}, &mockClient)

		// when
		_, err := kakaoRepository.GetUserInfo("test")

		// then
		assert.Error(t, err)
	})
}

func Test_requestToken(t *testing.T) {

	t.Run("parse token res if success requesting", func(t *testing.T) {
		// given
		var mockClient repository.KakaoClient = &MockClient{DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(successReqTokenBody))}, nil
		}}
		kakaoRepository := repository.NewKakaoAuthRepository(&config.KakaoAuthConfig{}, &mockClient)

		// when
		_, err := kakaoRepository.RequestToken("test")

		// then
		assert.NoError(t, err)

	})
}
