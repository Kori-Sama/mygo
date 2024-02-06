package model

import (
	utils "mygo/internal/pkg"
	"time"
)

type User struct {
	Id       int64
	Name     string
	Age      int
	Password string    `xorm:"varchar(200) notnull"`
	Wallet   string    `xorm:"varchar(200)"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}

func CreateUser(name, password string, age int) error {
	user := User{
		Name:     name,
		Password: password,
		Age:      age,
	}
	_, err := engine.Insert(&user)
	if err != nil {
		return utils.ErrorOperateDatabase
	}
	return nil
}

func GetUserById(id int64) (*User, error) {
	user := &User{}
	_, err := engine.ID(id).Get(user)
	if err != nil {
		return nil, utils.ErrorOperateDatabase
	}
	return user, nil
}

func GetUserByName(name string) (*User, error) {
	user := &User{}
	isFind, err := engine.Where("name = ?", name).Get(user)
	if err != nil {
		return nil, utils.ErrorOperateDatabase
	}
	if !isFind {
		return nil, utils.ErrorUnknownUsername
	}
	return user, nil
}

func (u *User) UpdatePassword(password string) error {
	u.Password = password
	_, err := engine.ID(u.Id).Cols("password").Update(u)
	if err != nil {
		return utils.ErrorOperateDatabase
	}
	return nil
}

func (u *User) UpdateWallet(wallet string) error {
	u.Wallet = wallet
	_, err := engine.ID(u.Id).Cols("wallet").Update(u)
	if err != nil {
		return utils.ErrorOperateDatabase
	}
	return nil
}
