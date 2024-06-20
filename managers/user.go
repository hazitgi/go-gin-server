package managers

import (
	"errors"

	"github.com/hazitgi/go_gin_server/common"
	"github.com/hazitgi/go_gin_server/database"
	"github.com/hazitgi/go_gin_server/models"
)

type UserManager interface {
	Create(userData *common.UserCreationInput) (*models.User, error)
	List() ([]models.User, error)
	Get(id string) (models.User, error)
	Delete(id string) error
	Update(userId string, userData *common.UserUpdateInput) (models.User, error)
}
type userManager struct {
	// dbClient *gorm.DB
}

func NewUserManager() UserManager {
	return &userManager{}
}

func (userMgr *userManager) Create(userData *common.UserCreationInput) (*models.User, error) {
	newUser := &models.User{
		FullName: userData.FullName,
		Email:    userData.Email,
	}
	database.GetDb().Create(newUser)
	if newUser.ID == 0 {
		return nil, errors.New("failed to create user")
	}
	return newUser, nil
}

func (userMgr *userManager) List() ([]models.User, error) {
	users := []models.User{}
	database.GetDb().Find(&users)
	return users, nil
}

func (userMgr *userManager) Get(id string) (models.User, error) {
	user := models.User{}
	database.GetDb().First(&user, id)
	return user, nil
}

func (userMgr *userManager) Delete(id string) error {
	user := models.User{}
	result := database.GetDb().Delete(&user, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (userMgr *userManager) Update(userId string, userData *common.UserUpdateInput) (models.User, error) {
	user := models.User{}
	result := database.GetDb().First(&user, userId)
	if result.Error != nil {
		return user, result.Error
	}
	user.FullName = userData.FullName
	user.Email = userData.Email

	updateResult := database.GetDb().Save(&user)

	if updateResult.Error != nil {
		return user, updateResult.Error
	}
	return user, nil
}
