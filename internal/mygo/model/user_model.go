package model

import (
	"mygo/internal/pkg/common"
	"time"

	log "github.com/sirupsen/logrus"
)

type User struct {
	ID         int         `xorm:"pk autoincr 'id'"`
	Name       string      `xorm:"varchar(200) notnull unique  'name'"`
	Password   string      `xorm:"varchar(32) notnull 'password'"`
	Role       common.Role `xorm:"enum('Old','Volunteer','Admin') default 'Old' 'role'"`
	Age        int         `xorm:"int 'age'"`
	Wallet     string      `xorm:"varchar(200) 'wallet'"`
	Passphrase string      `xorm:"varchar(200) 'passphrase'"`
	CreatedAt  time.Time   `xorm:"created 'created_at'"`
	UpdatedAt  time.Time   `xorm:"updated 'updated_at'"`
}

func CreateUser(name, password string, role common.Role) (int, error) {
	user := User{
		Name:     name,
		Password: password,
		Role:     role,
	}
	_, err := engine.Insert(&user)
	if err != nil {
		log.Debug("Pwd len:", len(password))
		log.Error("CreateUser error: ", err)
		return 0, common.ErrorOperateDatabase
	}
	return user.ID, nil
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
	_, err := engine.ID(u.ID).Update(u)
	if err != nil {
		return common.ErrorOperateDatabase
	}
	return nil
}

func (u *User) UpdatePassword(password string) error {
	u.Password = password
	_, err := engine.ID(u.ID).Cols("password").Update(u)
	if err != nil {
		return common.ErrorOperateDatabase
	}
	return nil
}

func (u *User) UpdateWallet(wallet string) error {
	u.Wallet = wallet
	_, err := engine.ID(u.ID).Cols("wallet").Update(u)
	if err != nil {
		return common.ErrorOperateDatabase
	}
	return nil
}

func (u *User) UpdatePassphrase(passphrase string) error {
	u.Passphrase = passphrase
	_, err := engine.ID(u.ID).Cols("passphrase").Update(u)
	if err != nil {
		return common.ErrorOperateDatabase
	}
	return nil
}
