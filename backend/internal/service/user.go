package service

import (
	"backend/internal/dao"
	"backend/internal/db"
	"backend/internal/dto"
	"backend/internal/model"
	"backend/pkg/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserDao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{
		UserDao: dao.NewUserDao(db.MysqlDatabase, db.MongoDatabase),
	}
}

func (u *UserService) Register(req dto.RegisterReq) error {
	_, err := u.UserDao.FindUserByAccount(req.Account)
	if err == nil {
		return errors.New("账号已存在")
	}
	_, err = u.UserDao.FindUserByEmail(req.Email)
	if err == nil {
		return errors.New("邮箱已存在")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &model.User{
		Account:  req.Account,
		Email:    req.Email,
		Username: req.Username,
		Password: string(password),
	}
	err = u.UserDao.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) Login(req dto.LoginReq) (string, error) {
	user, err := u.UserDao.FindUserByAccount(req.Account)
	if err != nil {
		user, err = u.UserDao.FindUserByEmail(req.Account)
		if err != nil {
			return "", err
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}
	token, err := utils.GenerateToken(uint(user.Id), user.RemuseId)
	if err != nil {
		return "", err
	}
	return token, nil
}
