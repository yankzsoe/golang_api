package repositories

import (
	"net/http"
	"strings"
	"time"

	"golang_api/app/dtos"
	"golang_api/app/models"
	config "golang_api/configs"
	"golang_api/tools"

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

func (repo *UserRepository) FindAll(param dtos.CommonParam) *[]dtos.GetUserResponse {
	results := []dtos.GetUserResponse{}

	rows, err := repo.DB.Raw("SELECT us.\"id\", us.username, us.nickname, us.email, us.created_date, us.updated_date, ro.role_id, ro.role_name"+
		" FROM \"users\" us"+
		" JOIN \"role\" ro ON us.role_id = ro.role_id"+
		" WHERE deleted_date IS NULL AND LOWER(us.username) LIKE ?", "%"+strings.ToLower(param.Where)+"%").Rows()
	if err != nil {
		tools.ThrowException(http.StatusInternalServerError, err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		scanRows := dtos.GetUserResponse{}
		err := repo.DB.ScanRows(rows, &scanRows)
		if err != nil {
			tools.ThrowException(http.StatusInternalServerError, err.Error())
		}

		results = append(results, dtos.GetUserResponse{
			Id:          scanRows.Id,
			Username:    scanRows.Username,
			Nickname:    scanRows.Nickname,
			Email:       scanRows.Email,
			CreatedDate: scanRows.CreatedDate,
			UpdatedDate: scanRows.UpdatedDate,
			RoleId:      scanRows.RoleId,
			RoleName:    scanRows.RoleName,
		})
	}

	return &results
}

func (repo *UserRepository) FindByID(id string) *dtos.GetUserResponse {
	var results dtos.GetUserResponse
	rows, err := repo.DB.Raw("SELECT us.\"id\", us.username, us.nickname, us.email, us.created_date, us.updated_date, ro.role_id, ro.role_name"+
		" FROM \"users\" us"+
		" JOIN \"role\" ro ON us.role_id = ro.role_id"+
		" WHERE deleted_date IS NULL AND us.\"id\" = ?", id).Rows()

	if err != nil {
		tools.ThrowException(http.StatusInternalServerError, err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		err := repo.DB.ScanRows(rows, &results)
		if err != nil {
			tools.ThrowException(http.StatusInternalServerError, err.Error())
		}
	}

	return &results
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
		panic(dtos.ErrorResponse{
			ErrorCode: http.StatusNotFound,
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
	result := repo.DB.Where("deleted_date IS NULL AND (username = ? OR email = ?)", username, username).Find(&user)
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
	rows, err := repo.DB.Raw("SELECT u.username, u.\"password\", r.role_code, r.role_name, m.module_code, m.module_name, rm.can_create, rm.can_read, rm.can_update, rm.can_delete"+
		" FROM users u"+
		" JOIN \"role\" r ON r.role_id = u.role_id"+
		" JOIN role_module rm ON rm.role_id = r.role_id"+
		" JOIN \"module\" m ON m.module_id = rm.module_id"+
		" WHERE u.deleted_date IS NULL AND (u.username = ? OR u.email = ?) AND rm.is_active = true", username, username).Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data dtos.UserWithClaimsResponse
		err := repo.DB.ScanRows(rows, &data)
		if err != nil {
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

func (repo *UserRepository) FindByIdWithRole(id string) (*[]dtos.UserWithClaimsResponse, error) {
	queryResult := []dtos.UserWithClaimsResponse{}
	rows, err := repo.DB.Raw("SELECT u.username, u.\"password\", r.role_name, m.module_name, rm.can_create, rm.can_read, rm.can_update, rm.can_delete"+
		" FROM users u"+
		" JOIN \"role\" r ON r.role_id = u.role_id"+
		" JOIN role_module rm ON rm.role_id = r.role_id"+
		" JOIN \"module\" m ON m.module_id = rm.module_id"+
		" WHERE u.deleted_date IS NULL AND u.Id = ?", id).Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data dtos.UserWithClaimsResponse
		err := repo.DB.ScanRows(rows, &data)
		if err != nil {
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

func (repo *UserRepository) CreateAccountBasicRole(userBasic dtos.UserBasic) (*string, error) {
	user := models.UserModel{
		Username:    userBasic.Username,
		Email:       userBasic.Email,
		Password:    userBasic.Password,
		CreatedDate: time.Now(),
		RoleId:      userBasic.RoleId,
	}

	result := repo.DB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user.Id, nil
}
