package dao

import (
	"backend/internal/model"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
)

type UserDao struct {
	mysqldb *gorm.DB
	mongo   *mongo.Database
}

func NewUserDao(mysqldb *gorm.DB, mongodb *mongo.Database) *UserDao {
	return &UserDao{
		mysqldb: mysqldb,
		mongo:   mongodb,
	}
}

func (u *UserDao) CreateUser(user *model.User) error {
	return u.mysqldb.Create(&user).Error
}

func (u *UserDao) FindUserByUserid(userid uint) (*model.User, error) {
	var user model.User
	err := u.mysqldb.Where("id = ?", userid).First(&user).Error
	return &user, err
}

func (u *UserDao) FindUserByAccount(account string) (*model.User, error) {
	var user model.User
	err := u.mysqldb.Where("account = ?", account).First(&user).Error
	return &user, err
}
func (u *UserDao) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.mysqldb.Where("email = ?", email).First(&user).Error
	return &user, err
}
