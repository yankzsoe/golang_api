package dtos

type RefreshTokenRequest struct {
	Token string `json:"token" validate:"jwt"`
}
