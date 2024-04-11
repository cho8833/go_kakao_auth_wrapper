package config

type KakaoAuthConfig struct {
	RedirectUrl string
	ClientId    string
}

func InitKakaoAuthConfig(redirectUrl string, clientId string) *KakaoAuthConfig {
	return &KakaoAuthConfig{RedirectUrl: redirectUrl, ClientId: clientId}
}
