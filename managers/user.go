package managers

import (
	"errors"
	"fmt"

	"github.com/hazitgi/go_gin_server/common"
	"github.com/hazitgi/go_gin_server/database"
	"github.com/hazitgi/go_gin_server/models"
)

type UserManager struct {
	// dbClient *gorm.DB
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (userMgr *UserManager) Create(userData *common.UserCreationInput) (*models.User, error) {
	newUser := &models.User{
		FullName: userData.FullName,
		Email:    userData.Email,
	}

	fmt.Println(newUser, ">>>>>>>>>>>>>>>>>>")
	database.GetDb().Create(newUser)

	if newUser.ID == 0 {
		return nil, errors.New("failed to create user")
	}

	return newUser, nil
}
