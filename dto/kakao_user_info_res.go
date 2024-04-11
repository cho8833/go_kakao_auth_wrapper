package dto

import "time"

type KakaoUserInfoRes struct {
	id           *int64
	HasSignedUp  *bool
	ConnectedAt  *time.Time
	SynchedAt    *time.Time
	Properties   *Property
	KakaoAccount *KakaoAccount `json:"kakao_account"`
}

type KakaoAccount struct {
	ProfileNeedsAgreement         *bool    `json:"profile_needs_agreement"`
	ProfileNicknameNeedsAgreement *bool    `json:"profile_nickname_needs_agreement"`
	ProfileImageNeedsAgreement    *bool    `json:"profile_image_needs_agreement"`
	Profile                       *Profile `json:"profile"`
	NameNeedsAgreement            *bool    `json:"name_needs_agreement"`
	Name                          *string  `json:"name"`
	EmailNeedsAgreement           *bool    `json:"email_needs_agreement"`
	IsEmailValid                  *bool
	IsEmailVerified               *bool
	Email                         *string `json:"email"`
	AgeRangeNeedsAgreement        *bool
	AgeRage                       *string
	BirthYearNeedsAgreement       *bool
	BirthYear                     *string
	BirthdayNeedsAgreement        *bool
	Birthday                      *string
	BirthdayType                  *string
	GenderNeedsAgreement          *bool
	Gender                        *string
	PhoneNumberNeedsAgreement     *bool   `json:"phone_number_needs_agreement"`
	PhoneNumber                   *string `json:"phone_number"`
	CiNeedsAgreement              *bool
	Ci                            *string
	CiAuthenticatedAt             *time.Time
}

type Profile struct {
	Nickname          *string
	ThumbnailImageUrl *string `json:"thumbnail_image_url"`
	ProfileImageUrl   *string
	IsDefaultImage    *bool
	IsDefaultNickname *bool
}

type Property struct {
	Nickname       *string
	ThumbnailImage *string `json:"thumbnail_image"`
	ProfileImage   *string `json:"profile_image"`
}
