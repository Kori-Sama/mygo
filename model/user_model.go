package model

import (
	"mygo/pkg/common"

	log "github.com/sirupsen/logrus"
)

type User struct {
	Id         int    `xorm:"pk autoincr notnull unique 'id'"`
	Name       string `xorm:"varchar(200) notnull unique"`
	Age        int    `xorm:"int"`
	Password   string `xorm:"varchar(32) notnull"`
	Wallet     string `xorm:"varchar(200)"`
	Passphrase string `xorm:"varchar(200)"`
}

func CreateUser(name, password string) (int, error) {
	user := User{
		Name:     name,
		Password: password,
	}
	_, err := engine.Insert(&user)
	if err != nil {
		log.Debug("Pwd len:", len(password))
		log.Error("CreateUser error: ", err)
		return 0, common.ErrorOperateDatabase
	}
	return user.Id, nil
}

func GetUserById(id int) (*User, error) {
	user := &User{}
	_, err := engine.ID(id).Get(user)
	if err != nil {
		return nil, common.ErrorOperateDatabase
	}
	return user, nil
}

func GetUserByName(name string) (*User, error) {
	user := &User{}
	isFind, err := engine.Where("name = ?", name).Get(user)
	if err != nil {
		return nil, common.ErrorOperateDatabase
	}
	if !isFind {
		return nil, common.ErrorUnknownUsername
	}
	return user, nil
}

func (u *User) UpdateUser() error {
	_, err := engine.ID(u.Id).Update(u)
	if err != nil {
		return common.ErrorOperateDatabase
	}
	return nil
}

func (u *User) UpdatePassword(password string) error {
	u.Password = password
	_, err := engine.ID(u.Id).Cols("password").Update(u)
	if err != nil {
		return common.ErrorOperateDatabase
	}
	return nil
}

func (u *User) UpdateWallet(wallet string) error {
	u.Wallet = wallet
	_, err := engine.ID(u.Id).Cols("wallet").Update(u)
	if err != nil {
		return common.ErrorOperateDatabase
	}
	return nil
}

func (u *User) UpdatePassphrase(passphrase string) error {
	u.Passphrase = passphrase
	_, err := engine.ID(u.Id).Cols("passphrase").Update(u)
	if err != nil {
		return common.ErrorOperateDatabase
	}
	return nil
}
