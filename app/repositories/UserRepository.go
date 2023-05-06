package repositories

import (
	"net/http"
	"time"

	config "golang_api/app/configs"
	"golang_api/app/dtos"
	"golang_api/app/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: config.GetDB(),
	}
}

func (repo *UserRepository) Create(user *models.UserModel) (*models.UserModel, error) {
	result := repo.DB.Create(user)
	if result.Error != nil {
		// return nil, errors.New("failed to create user")
		panic(dtos.ErrorResponse{
			ErrorCode: http.StatusInternalServerError,
			Message: dtos.Response{
				Status: dtos.BaseResponse{
					Success: false,
					Message: "failed to create user",
				},
			},
		})
	}
	return user, nil
}

func (repo *UserRepository) FindAll(param dtos.CommonParam) (*[]models.UserModel, error) {
	var user []models.UserModel
	result := repo.DB.Where("username LIKE ?", "%"+param.Where+"%").Limit(param.Limit).Offset(param.Offset).Find(&user)
	if result.Error != nil {
		// return nil, errors.New("failed to find user")
		panic(dtos.ErrorResponse{
			ErrorCode: http.StatusInternalServerError,
			Message: dtos.Response{
				Status: dtos.BaseResponse{
					Success: false,
					Message: "failed to find user",
				},
			},
		})
	}
	return &user, nil
}

func (repo *UserRepository) FindByID(id string) (*models.UserModel, error) {
	var user models.UserModel
	result := repo.DB.Where("id = ?", id).Find(&user)
	if result.Error != nil {
		// return nil, errors.New("failed to find user")
		panic(dtos.ErrorResponse{
			ErrorCode: http.StatusInternalServerError,
			Message: dtos.Response{
				Status: dtos.BaseResponse{
					Success: false,
					Message: "failed to find user",
				},
			},
		})
	}
	return &user, nil
}

func (repo *UserRepository) Update(userId string, user dtos.CreateOrUpdateUserRequest) error {
	tNow := time.Now()
	if err := repo.DB.Model(models.UserModel{}).Where("id = ?", userId).Updates(models.UserModel{
		Username:    user.Username,
		Nickname:    user.Nickname,
		Email:       user.Email,
		Password:    user.ConfirmPassword,
		UpdatedDate: &tNow,
	}).Error; err != nil {
		// return errors.New("failed to update user")
		panic(dtos.ErrorResponse{
			ErrorCode: http.StatusInternalServerError,
			Message: dtos.Response{
				Status: dtos.BaseResponse{
					Success: false,
					Message: "failed to update user",
				},
			},
		})
	}
	return nil
}

func (repo *UserRepository) Delete(userId string) error {
	user := models.UserModel{}
	if err := repo.DB.Clauses(clause.Returning{}).Delete(&user, "id", userId).Error; err != nil {
		// return errors.New("failed to delete user")
		panic(dtos.ErrorResponse{
			ErrorCode: http.StatusInternalServerError,
			Message: dtos.Response{
				Status: dtos.BaseResponse{
					Success: false,
					Message: "failed to delete user",
				},
			},
		})
	}

	if user.Id == "" {
		// return errors.New("user not found")
		panic(dtos.ErrorResponse{
			ErrorCode: http.StatusInternalServerError,
			Message: dtos.Response{
				Status: dtos.BaseResponse{
					Success: false,
					Message: "user not found",
				},
			},
		})
	}

	return nil
}

func (repo *UserRepository) FindByUsernameOrEmail(username string) (*models.UserModel, error) {
	var user models.UserModel
	result := repo.DB.Where("username = ? OR email = ?", username, username).Find(&user)
	if result.Error != nil {
		// return nil, errors.New("failed to find user")
		panic(dtos.ErrorResponse{
			ErrorCode: http.StatusInternalServerError,
			Message: dtos.Response{
				Status: dtos.BaseResponse{
					Success: false,
					Message: "failed to find user",
				},
			},
		})
	}
	return &user, nil
}

func (repo *UserRepository) FindByUsernameOrEmailWithRole(username string) (*[]dtos.UserWithClaimsResponse, error) {
	queryResult := []dtos.UserWithClaimsResponse{}
	rows, err := repo.DB.Raw("SELECT u.username, u.\"password\", r.role_name, m.module_name, rm.can_create, rm.can_read, rm.can_update, rm.can_delete"+
		" FROM users u"+
		" JOIN \"role\" r ON r.role_id = u.role_id"+
		" JOIN role_module rm ON rm.role_id = r.role_id"+
		" JOIN \"module\" m ON m.module_id = rm.module_id"+
		" WHERE u.username = ? OR u.email = ?", username, username).Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data dtos.UserWithClaimsResponse
		err := repo.DB.ScanRows(rows, &data)
		if err != nil {
			// return nil, err
			panic(dtos.ErrorResponse{
				ErrorCode: http.StatusInternalServerError,
				Message: dtos.Response{
					Status: dtos.BaseResponse{
						Success: false,
						Message: err.Error(),
					},
				},
			})
		}
		queryResult = append(queryResult, data)
	}

	return &queryResult, nil
}
