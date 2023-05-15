package resposity

import (
	"Gin_In_Action/util"
	"gorm.io/gorm"
	"sync"
	"time"
)

type User struct {
	Id         int64     `gorm:"colum:id"`
	Name       string    `gorm:"colum:name"`
	Avatar     string    `gorm:"colum:avatar"`
	Level      int       `gorm:"colum:level"`
	CreateTime time.Time `gorm:"colum:create_time"`
	ModifyTime time.Time `gorm:"colum:modify_time"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var (
	userDao  *UserDao
	userOnce sync.Once
)

func NewUserDaoInstance() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

func (*UserDao) QueryUserById(id int64) (*User, error) {
	var user User
	err := db.Where("id = ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		util.Logger.Error("find user by id err" + err.Error())
		return nil, err
	}

	return &user, nil
}

func (*UserDao) BenchQueryUsersByIds(ids []int64) (map[int64]*User, error) {
	var users []*User
	err := db.Where("id in (?)", ids).Find(&users).Error
	if err != nil {
		util.Logger.Error("bench find user by ids error" + err.Error())
		return nil, err
	}
	userMap := make(map[int64]*User)
	for _, user := range users {
		userMap[user.Id] = user
	}
	return userMap, nil
}
