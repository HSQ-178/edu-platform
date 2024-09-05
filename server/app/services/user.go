package services

import (
	"edu-profit/app/models"
	"edu-profit/database"
	"edu-profit/utils"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"regexp"
	"strconv"
)

type UserService interface {
	Register(u *models.UserRegisterReq) error
	Login(u *models.UserLoginReq) (models.UserLoginResp, error)
	List(u *models.UserReq) (models.UserListResp, error)
	Update(u *models.User) error
	findUserByCredential(credential string, userData *models.UserResp) error
	matchRegexp(pattern, value string) bool
}

type UserServiceImpl struct {
}

const (
	StatusNormal int = 1 + iota
	StatusFrozen
	StatusDeleted
)

const (
	RegexpForPhone    = "/^(?:(?:\\+|00)86)?1(?:(?:3[\\d])|(?:4[5-79])|(?:5[0-35-9])|(?:6[5-7])|(?:7[0-8])|(?:8[\\d])|(?:9[1589]))\\d{8}$/"
	RegexpForEmail    = "/^(([^<>()[\\]\\\\.,;:\\s@\"]+(\\.[^<>()[\\]\\\\.,;:\\s@\"]+)*)|(\".+\"))@((\\[[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\])|(([a-zA-Z\\-0-9]+\\.)+[a-zA-Z]{2,}))$/"
	RegexpForUsername = "/^[\\w-]{4,16}$/"
)

func (UserServiceImpl) Register(req *models.UserRegisterReq) error {

	var userData models.User

	// 检查用户名重复
	row := database.GetMySQL().Model(&models.User{}).Where("username = ? AND status != ?", req.Username, StatusDeleted).First(&userData).RowsAffected
	if row > 0 {
		return errors.New("用户名已存在")
	}

	// 检查昵称重复
	row = database.GetMySQL().Model(&models.User{}).Where("nickname = ? AND status != ?", req.Username, StatusDeleted).First(&userData).RowsAffected
	if row > 0 {
		return errors.New("昵称已存在")
	}

	user := &models.User{
		ID:       (&utils.Snowflake{}).NextVal(),
		Username: req.Username,
		Password: utils.MD5(req.Password),
		Nickname: req.Nickname,
		Status:   StatusNormal,
	}

	return database.GetMySQL().Create(user).Error

}

func (UserServiceImpl) Login(req *models.UserLoginReq) (models.UserLoginResp, error) {

	var userData models.UserResp

	switch req.Type {
	case 1: // 用户名/手机号/邮箱 + 密码
		if err := findUserByCredential(req.Username, &userData); err != nil {
			return models.UserLoginResp{}, err
		}

		if userData.Password != utils.MD5(req.Password) {
			return models.UserLoginResp{}, errors.New("用户名或密码错误")
		}
		break

	case 2: // 手机号 + 验证码
		return models.UserLoginResp{}, nil

	case 3: // 邮箱 + 验证码
		return models.UserLoginResp{}, nil

	default:
		return models.UserLoginResp{}, errors.New("缺少参数")
	}

	if userData.Status == StatusFrozen {
		return models.UserLoginResp{}, errors.New("用户已冻结")
	}

	if userData.Status == StatusDeleted {
		return models.UserLoginResp{}, errors.New("用户已注销")
	}

	token, err := utils.GenerateToken(userData.ID, userData.Username)
	if err != nil {
		return models.UserLoginResp{}, errors.New("生成token失败: " + err.Error())
	}

	loginResp := models.UserLoginResp{
		Token: token,
		User:  userData,
	}

	return loginResp, nil
}

func (UserServiceImpl) List(req *models.UserReq) (models.UserListResp, error) {

	var userListResp models.UserListResp

	if req.IDStr != "" {
		req.ID, _ = strconv.ParseInt(req.IDStr, 10, 64)
	}

	db := database.GetMySQL().Model(&models.User{}).Order(req.OrderBy + " " + req.Sorted)

	filters := []QueryOption{
		WithID(req.ID),
		WithRoleID(req.RoleID),
		WithUsername(req.Username),
		WithNickname(req.Nickname),
		WithEmail(req.Email),
		WithPhone(req.Phone),
		WithStatus(req.Status),
		WithDateRange(req.DateRange),
		WithPagination(req.Pagination),
	}

	ApplyFilters(db, filters...)

	if err := db.Count(&userListResp.Total).Error; err != nil {
		return userListResp, errors.New("查询失败")
	}

	if err := db.Preload(clause.Associations).Find(&userListResp.Records).Error; err != nil {
		return userListResp, errors.New("查询失败")
	}

	return userListResp, nil
}

func (UserServiceImpl) Update(u *models.User) error {
	if u.Password != "" {
		u.Password = utils.MD5(u.Password)
	}

	if u.Status == StatusDeleted {
		var user models.User

		err := database.GetMySQL().Model(&models.User{}).First(&user, user.ID).Error
		if err != nil {
			return errors.New("用户不存在")
		}

		u.Username = user.Username + "_del"
		u.Nickname = user.Nickname + "_del"
	}

	err := database.GetMySQL().Model(&models.User{}).Updates(u).Error
	if err != nil {
		return errors.New("更新失败")
	}

	return nil
}

// 将正则表达式匹配和数据库查询逻辑提取到一个单独的函数
func findUserByCredential(credential string, userData *models.UserResp) error {

	var err error

	switch {
	case matchRegexp(RegexpForUsername, credential): // 用户名
		err = database.GetMySQL().Model(&models.User{}).Where("username = ?", credential).Preload(clause.Associations).First(userData).Error
	case matchRegexp(RegexpForPhone, credential): // 手机号
		err = database.GetMySQL().Model(&models.User{}).Where("phone = ?", credential).Preload(clause.Associations).First(userData).Error
	case matchRegexp(RegexpForEmail, credential): // 邮箱
		err = database.GetMySQL().Model(&models.User{}).Where("email = ?", credential).Preload(clause.Associations).First(userData).Error
	default:
		return errors.New("无效的用户名、手机号或邮箱")
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("用户不存在")
	}

	return err
}

// 匹配正则表达式的辅助函数
func matchRegexp(pattern, value string) bool {
	matched, _ := regexp.MatchString(pattern, value)
	return matched
}
