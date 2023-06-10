package services

import (
	"golang_api/app/dtos"
	"golang_api/app/models"
	"golang_api/app/repositories"
	"golang_api/tools"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (u *UserService) CreateUser(user *models.UserModel) (*models.UserModel, error) {
	// Enkripsi password
	aes128 := tools.Aes128{}
	encryptedPassword, err := aes128.Encrypt(user.Password)
	if err != nil {
		return nil, err
	}

	// Simpan user ke database
	user.Password = *encryptedPassword
	return u.userRepository.Create(user)
}

func (u *UserService) GetUserByID(userID string) (*models.UserModel, error) {
	return u.userRepository.FindByID(userID)
}

func (u *UserService) GetAllUser(param dtos.CommonParam) (*[]models.UserModel, error) {
	return u.userRepository.FindAll(param)
}

func (u *UserService) UpdateUser(userId string, user dtos.CreateOrUpdateUserRequest) error {
	aes128 := tools.Aes128{}
	// Encrypt password before put to database
	encryptedPassword, err := aes128.Encrypt(user.ConfirmPassword)
	if err != nil {
		return err
	}

	user.ConfirmPassword = *encryptedPassword
	return u.userRepository.Update(userId, user)
}

func (u *UserService) DeleteUser(userId string) error {
	return u.userRepository.Delete(userId)
}
