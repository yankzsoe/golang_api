package services

import (
	"errors"
	"strings"
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

	tokenString, err := token.SignedString([]byte(configs.GetJWTConfigurationInstance().Key))
	if err != nil {
		return nil, err
	}

	// Create Refresh Token
	refreshToken, errs := GenerateRefreshToken(username, roleName, permissions)
	if errs != nil {
		return nil, errs
	}

	return &dtos.Token{
		Value:        tokenString,
		IssuedOn:     time.Unix(iat.Unix(), 0),
		ExpiresOn:    time.Unix(exp.Unix(), 0),
		RefreshToken: *refreshToken,
	}, nil
}

func GenerateRefreshToken(username string, roleName string, permissions []dtos.Permission) (*dtos.RefreshToken, error) {
	iat := jwt.NewNumericDate(time.Now())
	// Refresh token will active after main token is expired
	nbf := jwt.NewNumericDate(time.Now().Add(time.Millisecond * time.Duration(configs.GetJWTConfigurationInstance().Exp)))
	// Refresh token will expire twice in its lifetime from main token.
	exp := jwt.NewNumericDate(time.Now().Add(time.Millisecond * time.Duration(configs.GetJWTConfigurationInstance().Exp*2)))

	claims := &dtos.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    configs.GetJWTConfigurationInstance().Iis,
			Audience:  configs.GetJWTConfigurationInstance().Aud,
			IssuedAt:  iat,
			NotBefore: nbf,
			ExpiresAt: exp,
		},
		Role: dtos.RoleSwagger{
			Name:        roleName,
			Permissions: permissions,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(configs.GetJWTConfigurationInstance().Key))
	if err != nil {
		return nil, err
	}

	return &dtos.RefreshToken{
		Value:     tokenString,
		IssuedOn:  time.Unix(iat.Unix(), 0),
		NotBefore: time.Unix(nbf.Unix(), 0),
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

		user := dtos.UserBasic{
			Username: dto.Given_Name,
			Email:    dto.Email,
			Password: *encryptedPassword,
			RoleId:   "77691975-c695-4afd-aa40-286f2a26857d", // role user
		}

		userId, err := u.userRepository.CreateAccountBasicRole(user)
		if err != nil {
			// When an error occurs due to a duplicate name, we will try using email instead.
			if strings.Contains(err.Error(), "duplicate key") {
				user.Username = dto.Email
				userIds, errs := u.userRepository.CreateAccountBasicRole(user)
				if errs != nil {
					return 500, nil, errs
				}

				userId = userIds
			} else {
				return 500, nil, err
			}
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

func (u *AuthenticationService) RefreshToken(dto dtos.RefreshTokenRequest) (int, *dtos.Token, error) {
	if err := tools.GenerateErrorMessage(dto); err != nil {
		panic(dtos.ErrorResponse{
			ErrorCode: 400,
			Message: dtos.Response{
				Status: dtos.BaseResponse{
					Success: false,
					Message: "Failed On Validation",
				},
				Data: err,
			},
		})
	}

	username := CheckToken(dto.Token)

	userWithRoles, err := u.userRepository.FindByUsernameOrEmailWithRole(username)
	if err != nil {
		return 500, nil, err
	}

	if len(*userWithRoles) == 0 {
		return 404, nil, errors.New("user not found")
	}

	permissions := ConvertToPermissionsList(*userWithRoles)

	token, err := GenerateToken((*userWithRoles)[0].Username, (*userWithRoles)[0].RoleName, permissions)
	if err != nil {
		return 500, nil, err
	}

	return 200, token, nil
}

func CheckToken(tokenString string) string {
	jwtKey := []byte(configs.GetJWTConfigurationInstance().Key)
	token, err := jwt.ParseWithClaims(tokenString, &dtos.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKey
		}
		return jwtKey, nil
	})

	if err != nil {
		panic(dtos.ErrorResponse{
			ErrorCode: 401,
			Message: dtos.Response{
				Status: dtos.BaseResponse{
					Success: false,
					Message: err.Error(),
				},
			},
		})
	}

	claims, ok := token.Claims.(*dtos.Claims)
	if !ok || !token.Valid {
		panic(dtos.ErrorResponse{
			ErrorCode: 401,
			Message: dtos.Response{
				Status: dtos.BaseResponse{
					Success: false,
					Message: "Not Authorize",
				},
			},
		})
	}

	return claims.Username
}
