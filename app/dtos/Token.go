package dtos

import "time"

type Token struct {
	Value     string    `json:"value"`
	IssuedOn  time.Time `json:"issuedOn"`
	ExpiresOn time.Time `json:"expiresOn"`
}
