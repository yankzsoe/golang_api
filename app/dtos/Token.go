package dtos

import "time"

type Token struct {
	Value        string       `json:"token"`
	IssuedOn     time.Time    `json:"issuedOn"`
	ExpiresOn    time.Time    `json:"expiresOn"`
	RefreshToken RefreshToken `json:"refreshToken"`
}

type RefreshToken struct {
	Value     string    `json:"token"`
	IssuedOn  time.Time `json:"issuedOn"`
	NotBefore time.Time `json:"notBefore"`
	ExpiresOn time.Time `json:"expiresOn"`
}
