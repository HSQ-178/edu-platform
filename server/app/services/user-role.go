package services

import (
	"edu-profit/app/models"
	"edu-profit/database"
	"errors"
	"gorm.io/gorm/clause"
)

type UserRoleService interface {
	Create(req *[]models.UserRole) error
	List(req *models.UserRoleReq) (models.UserRoleListResp, error)
	Update(req *[]models.UserRoleReq) error
	Delete(req *[]models.UserRole) error
}

type UserRoleServiceImpl struct{}

func (UserRoleServiceImpl) Create(req *[]models.UserRole) error {

	tx := database.GetMySQL().Begin()

	for _, user := range *req {
		if err := tx.Create(&user).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (UserRoleServiceImpl) List(req *models.UserRoleReq) (models.UserRoleListResp, error) {

	var resp models.UserRoleListResp

	db := database.GetMySQL().Model(&models.UserRole{}).Order(req.OrderBy + " " + req.Sorted)

	filters := []QueryOption{
		WithID32(req.ID),
		WithRoleName(req.RoleName),
		WithDateRange(req.DateRange),
		WithPagination(req.Pagination),
	}

	ApplyFilters(db, filters...)

	if err := db.Count(&resp.Total).Error; err != nil {
		return resp, errors.New("查询失败")
	}

	if err := db.Preload(clause.Associations).Find(&resp.Records).Error; err != nil {
		return resp, errors.New("查询失败")
	}

	return resp, nil
}

func (UserRoleServiceImpl) Update(req *[]models.UserRoleReq) error {
	tx := database.GetMySQL().Begin()

	for _, r := range *req {
		if err := database.GetMySQL().Model(&models.User{}).Updates(r).Error; err != nil {
			tx.Rollback()
			return errors.New("部分数据存在异常，操作失败")
		}
	}

	tx.Commit()
	return nil
}

func (UserRoleServiceImpl) Delete(req *[]models.UserRole) error {

	tx := database.GetMySQL().Begin()

	for _, r := range *req {

		var userRole models.UserRole

		err := database.GetMySQL().Model(&models.UserRole{}).First(&userRole, userRole.ID).Error
		if err != nil {
			return errors.New("数据不存在")
		}

		r.RoleName = r.RoleName + "_del"

		if err := database.GetMySQL().Model(&models.UserRole{}).Updates(r).Error; err != nil {
			tx.Rollback()
			return errors.New("部分数据存在异常，操作失败")
		}

		if err := database.GetMySQL().Delete(&r).Error; err != nil {
			tx.Rollback()
			return errors.New("部分数据存在异常，操作失败")
		}
	}

	tx.Commit()
	return nil
}
