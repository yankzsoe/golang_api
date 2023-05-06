package dtos

import "github.com/golang-jwt/jwt/v5"

type Permission struct {
	Module    string `json:"module"`
	CanCreate bool   `json:"canCreate"`
	CanRead   bool   `json:"canRead"`
	CanUpdate bool   `json:"canUpdate"`
	CanDelete bool   `json:"canDelete"`
}

type RoleSwagger struct {
	Name        string       `json:"name"`
	Permissions []Permission `json:"actions"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
	Role RoleSwagger
}
