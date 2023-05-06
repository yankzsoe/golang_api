package configs

import (
	"sync"

	"golang_api/app/tools"
)

var (
	instance *JwtConfiguration
	onceJwt  sync.Once
)

type JwtConfiguration struct {
	Key string   `json:"key"`
	Iis string   `json:"iis"`
	Aud []string `json:"aud"`
	Exp int64    `json:"exp"`
}

func GetJWTConfigurationInstance() *JwtConfiguration {
	onceJwt.Do(func() {
		reader := tools.ConfigReader{}
		jtwConfig := JwtConfiguration{}
		reader.ReadFileConfiguration("/app/configs/jwt.json", &jtwConfig)
		instance = &jtwConfig
	})
	return instance
}
