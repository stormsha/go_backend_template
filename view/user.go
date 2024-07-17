package view

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"stormsha.com/gbt/model"
	"stormsha.com/gbt/utils"
)

type Table struct {
	Name string
}

// noinspection SpellCheckingInspection(忽略单词检查，去掉波浪线，强迫症行为)
func Register(user *model.User) (*model.LoginResp, error) {
	if len(user.UserAccount) < 6 {
		return nil, errors.New("密码长度不能低于6个字符")
	}
	// 验证账号是否包含中文字符
	match, _ := regexp.MatchString(`^[a-zA-Z0-9]+$`, user.UserAccount)
	if !match {
		return nil, errors.New("账号只能由a-z,A-Z,0-9组成")
	}
	// 哈希密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.UserPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.Errorf("密码加密失败 %v", err)
		return nil, errors.New("注册失败")
	}
	user.UserPassword = string(hashedPassword)

	// 插入新用户
	tx := GormDB.Begin()
	defer tx.Rollback()
	tx.Create(user)
	if err := tx.Commit().Error; err != nil {
		logger.Errorf("%v 账号创建失败 提交事务失败 %v", tx.Error.Error())
		return nil, errors.New("注册失败")
	}

	tokenString, err := utils.GetUserToken(user)
	if err != nil {
		return nil, errors.New("注册成功，请重新登录")
	}

	res := &model.LoginResp{
		ID:          user.ID,
		UserAccount: user.UserAccount,
		UserName:    user.UserName,
		UserAvatar:  user.UserAvatar,
		Token:       tokenString,
	}
	return res, nil
}

// noinspection SpellCheckingInspection(忽略单词检查，去掉波浪线，强迫症行为)
func Login(args *model.User) (*model.LoginResp, error) {
	user := new(model.User)
	result := GormDB.Where("user_account = ?", args.UserAccount).First(&user)
	if result.Error != nil {
		logger.Errorf("登录失败 %v", result.Error)
		return nil, errors.New("登录失败")
	}

	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(args.UserPassword))
	if err != nil {
		return nil, errors.New("密码错误")
	}

	tokenString, err := utils.GetUserToken(user)
	if err != nil {
		logger.Errorf("获取 token 失败 %v", err)
		return nil, errors.New("登录失败")
	}

	res := &model.LoginResp{
		ID:          user.ID,
		UserAccount: user.UserAccount,
		UserName:    user.UserName,
		UserAvatar:  user.UserAvatar,
		Token:       tokenString,
	}
	return res, nil
}

func Detail(user *model.User) (*model.UserDetail, error) {
	result := GormDB.Where("id = ?", user.ID).First(&user)
	if result.Error != nil {
		logger.Errorf("获取用户失败 %v", result.Error)
		return nil, errors.New("用户不存在！！！")
	}
	res := &model.UserDetail{
		ID:          user.ID,
		UserAccount: user.UserAccount,
		UnionID:     user.UnionID,
		OpenID:      user.OpenID,
		UserName:    user.UserName,
		UserAvatar:  user.UserAvatar,
		UserProfile: user.UserProfile,
		UserRole:    user.UserRole,
	}
	return res, nil
}
