package services

import (
	"errors"
	"time"

	"golang_api/app/configs"
	"golang_api/app/dtos"
	"golang_api/app/repositories"
	"golang_api/app/tools"

	"github.com/golang-jwt/jwt/v5"
)

const key = "abcdefghij1234567890"

type AuthenticationService struct {
	userRepository repositories.UserRepository
}

func NewAuthenticationService(userRepository repositories.UserRepository) *AuthenticationService {
	return &AuthenticationService{
		userRepository: userRepository,
	}
}

func (u *AuthenticationService) Login(dto dtos.LoginRequest) (int, *dtos.Token, error) {
	userWithRoles, err := u.userRepository.FindByUsernameOrEmailWithRole(dto.Username)

	if err != nil {
		return 500, nil, err
	}

	if len(*userWithRoles) == 0 {
		return 404, nil, errors.New("user not found")
	}

	aes128 := tools.Aes128{}
	decryptedPassword, err := aes128.Decrypt((*userWithRoles)[0].Password)
	if err != nil {
		return 500, nil, err
	}

	if dto.Password != *decryptedPassword {
		return 401, nil, errors.New("invalid password")
	}

	permissions := ConvertToPermissionsList(*userWithRoles)

	token, err := GenerateToken((*userWithRoles)[0].Username, (*userWithRoles)[0].RoleName, permissions)
	if err != nil {
		return 500, nil, err
	}

	return 200, token, nil
}

func GenerateToken(username string, roleName string, permissions []dtos.Permission) (*dtos.Token, error) {
	iat := jwt.NewNumericDate(time.Now())
	exp := jwt.NewNumericDate(time.Now().Add(time.Millisecond * time.Duration(configs.GetJWTConfigurationInstance().Exp)))

	claims := &dtos.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    configs.GetJWTConfigurationInstance().Iis,
			Audience:  configs.GetJWTConfigurationInstance().Aud,
			IssuedAt:  iat,
			ExpiresAt: exp,
		},
		Role: dtos.RoleSwagger{
			Name:        roleName,
			Permissions: permissions,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	tokenString, err := token.SignedString([]byte(configs.GetJWTConfigurationInstance().Key))
	if err != nil {
		return nil, err
	}

	return &dtos.Token{
		Value:     tokenString,
		IssuedOn:  time.Unix(iat.Unix(), 0),
		ExpiresOn: time.Unix(exp.Unix(), 0),
	}, nil
}

func ConvertToPermissionsList(data []dtos.UserWithClaimsResponse) []dtos.Permission {
	result := []dtos.Permission{}
	for _, r := range data {
		result = append(result, dtos.Permission{
			Module:    r.ModuleName,
			CanCreate: r.CanCreate,
			CanRead:   r.CanRead,
			CanUpdate: r.CanUpdate,
			CanDelete: r.CanDelete,
		})
	}
	return result
}

func (u *AuthenticationService) LoginOAuthGoogle(dto tools.UserInfoResponse) (int, *dtos.Token, error) {
	aes128 := tools.Aes128{}
	userWithRoles, err := u.userRepository.FindByUsernameOrEmailWithRole(dto.Email)

	if err != nil {
		return 500, nil, err
	}

	if len(*userWithRoles) == 0 {
		// If user is'n exist password will dump
		encryptedPassword, err := aes128.Encrypt("qwerty123")
		if err != nil {
			return 500, nil, err
		}

		dto := dtos.UserBasic{
			Username: dto.Given_Name,
			Email:    dto.Email,
			Password: *encryptedPassword,
			RoleId:   "77691975-c695-4afd-aa40-286f2a26857d", // role user
		}

		userId, err := u.userRepository.CreateAccountBasicRole(dto)
		if err != nil {
			return 500, nil, err
		}

		userWithRoles, err = u.userRepository.FindByIdWithRole(*userId)
		if err != nil {
			return 500, nil, err
		}
	}

	permissions := ConvertToPermissionsList(*userWithRoles)

	token, err := GenerateToken((*userWithRoles)[0].Username, (*userWithRoles)[0].RoleName, permissions)
	if err != nil {
		return 500, nil, err
	}

	return 200, token, nil
}
