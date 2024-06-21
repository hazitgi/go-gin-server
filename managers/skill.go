package managers

import (
	"errors"
	"fmt"

	"github.com/hazitgi/go_gin_server/common"
	"github.com/hazitgi/go_gin_server/database"
	"github.com/hazitgi/go_gin_server/models"
)

type SkillManager interface {
	// Skill
	Create(inputData *common.SkillCreationInput) (*models.Skill, error)
	List() ([]models.Skill, error)
	Get(id string) (*models.Skill, error)
	Update(id string, inputData *common.SkillUpdateInput) (*models.Skill, error)
	Delete(id string) error
	// Skill Group
	CreateGroup(inputData *common.SkillGroupCreationInput) (*models.SkillGroup, error)
	ListGroup() ([]models.SkillGroup, error)
	GetGroup(id string) (*models.SkillGroup, error)
	UpdateGroup(id string, inputData *common.SkillGroupUpdateInput) (*models.SkillGroup, error)
	DeleteGroup(id string) error
}

type skillManager struct {
	// DatabaseDriver
	// dbClient
}

func NewSkillManager() SkillManager {
	return &skillManager{}
}

func (skillMgr *skillManager) Create(inputData *common.SkillCreationInput) (*models.Skill, error) {
	newSkillObj := &models.Skill{Name: inputData.Name, SkillGroupID: inputData.Group}
	database.GetDb().Create(newSkillObj)
	if newSkillObj.ID == 0 {
		return nil, errors.New("skill creation failed")
	}

	return newSkillObj, nil
}

func (skillMgr *skillManager) List() ([]models.Skill, error) {
	skillObj := []models.Skill{}

	result := database.GetDb().Preload("SkillGroup").Find(&skillObj)
	if result.Error != nil {
		return nil, result.Error
	}
	return skillObj, nil
}

func (skillMgr *skillManager) Get(id string) (*models.Skill, error) {
	skillObj := models.NewSkill()

	result := database.GetDb().Preload("SkillGroup").First(skillObj, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if skillObj.ID == 0 {
		return nil, errors.New("no skill found")
	}

	return skillObj, nil
}

func (skillMgr *skillManager) Update(id string, inputData *common.SkillUpdateInput) (*models.Skill, error) {
	skillObj := &models.Skill{}

	result := database.GetDb().First(skillObj, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if skillObj.ID == 0 {
		return nil, errors.New("item does not exist")
	}

	updateResult := database.GetDb().Model(skillObj).Updates(models.Skill{Name: inputData.Name})

	if updateResult.Error != nil {
		return skillObj, updateResult.Error
	}
	return skillObj, nil
}

func (skillMgr *skillManager) Delete(id string) error {
	skillObj := models.Skill{}

	result := database.GetDb().First(&skillObj, id)
	if result.Error != nil {
		return result.Error
	}

	if skillObj.ID == 0 {
		return errors.New("item does not exist")
	}
	result = database.GetDb().Delete(&skillObj)
	if result.Error != nil {
		return result.Error
	}
	// TODO: handle errors
	return nil
}

func (skillMgr *skillManager) CreateGroup(inputData *common.SkillGroupCreationInput) (*models.SkillGroup, error) {
	newSkillGroupObj := &models.SkillGroup{Name: inputData.Name}
	result := database.GetDb().Create(newSkillGroupObj)
	if result.Error != nil {
		return nil, result.Error
	}

	if newSkillGroupObj.ID == 0 {
		return nil, errors.New("skill group creation failed")
	}

	return newSkillGroupObj, nil
}

func (skillMgr *skillManager) ListGroup() ([]models.SkillGroup, error) {
	skillGroups := []models.SkillGroup{}

	// database.GetDb().Find(&skillGroupObj)
	result := database.GetDb().Model(&skillGroups).Preload("Skills").Find(&skillGroups)
	if result.Error != nil {
		return nil, result.Error
	}
	// TODO: handle errors
	return skillGroups, nil
}

func (skillMgr *skillManager) GetGroup(id string) (*models.SkillGroup, error) {
	skillGroupObj := models.NewSkillGroup()

	database.GetDb().Model(skillGroupObj).Preload("Skills").Find(skillGroupObj)

	if skillGroupObj.ID == 0 {
		return nil, errors.New("item does not exist")
	}

	return skillGroupObj, nil
}

func (skillMgr *skillManager) UpdateGroup(id string, inputData *common.SkillGroupUpdateInput) (*models.SkillGroup, error) {
	skillGroupObj := models.NewSkillGroup()

	database.GetDb().First(skillGroupObj, id)

	if skillGroupObj.ID == 0 {
		return nil, errors.New("item does not exist")
	}

	skillMapping, err := skillMgr.getSkillList(inputData.Skills)
	if err != nil {
		return nil, err
	}

	database.GetDb().Model(skillGroupObj).Updates(models.Skill{Name: inputData.Name})
	database.GetDb().Model(skillGroupObj).Association("Skills").Replace(skillMapping)
	// TODO: handle errors
	return skillGroupObj, nil
}

func (skillMgr *skillManager) getSkillList(inputList []int) ([]*models.Skill, error) {
	skills := []*models.Skill{}

	for _, id := range inputList {
		skill := models.NewSkill()
		database.GetDb().First(skill, id)
		if skill.ID == 0 {
			return nil, fmt.Errorf("skill with id %v not exists", id)
		}
		skills = append(skills, skill)
	}

	return skills, nil
}

func (skillMgr *skillManager) DeleteGroup(id string) error {
	skillGroupObj := models.NewSkillGroup()

	database.GetDb().First(skillGroupObj, id)

	if skillGroupObj.ID == 0 {
		return errors.New("item does not exist")
	}
	database.GetDb().Delete(skillGroupObj)
	// TODO: handle errors
	return nil
}
