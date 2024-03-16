package model

import (
	"mygo/internal/pkg/common"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type User struct {
	ID         int         `xorm:"pk autoincr 'id'"`
	Name       string      `xorm:"varchar notnull unique  'name'"`
	Password   string      `xorm:"varchar(32) notnull 'password'"`
	Role       common.Role `xorm:"type:role 'role'"`
	Age        int         `xorm:"int 'age'"`
	Wallet     int         `xorm:"varchar 'wallet'"`
	Passphrase string      `xorm:"varchar 'passphrase'"`
	CreatedAt  time.Time   `xorm:"created 'created_at'"`
	UpdatedAt  time.Time   `xorm:"updated 'updated_at'"`
}

func (u *User) ToResponse() *common.UserResponse {
	return &common.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Age:       u.Age,
		Role:      u.Role,
		CreatedAt: u.CreatedAt.Unix(),
		UpdatedAt: u.UpdatedAt.Unix(),
	}
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

func GetAllUsers() ([]User, error) {
	users := make([]User, 0)
	err := engine.Find(&users)
	if err != nil {
		return nil, common.ErrorOperateDatabase
	}
	return users, nil
}

func GetUserById(id int) (*User, error) {
	user := &User{}
	isFind, err := engine.ID(id).Get(user)
	if err != nil {
		return nil, common.ErrorOperateDatabase
	}
	if !isFind {
		return nil, common.ErrorUnknownUserId
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
	err := error(nil)
	u.Wallet, err = strconv.Atoi(wallet)
	if err != nil {
		return common.ErrorOperateDatabase
	}
	_, err = engine.ID(u.ID).Cols("wallet").Update(u)
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
