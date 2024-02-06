package model

import "time"

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
		return err
	}
	return nil
}

func GetUserById(id int64) (*User, error) {
	user := &User{}
	_, err := engine.ID(id).Get(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByName(name string) (*User, error) {
	user := &User{}
	_, err := engine.Where("name = ?", name).Get(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) UpdatePassword(password string) error {
	u.Password = password
	_, err := engine.ID(u.Id).Cols("password").Update(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateWallet(wallet string) error {
	u.Wallet = wallet
	_, err := engine.ID(u.Id).Cols("wallet").Update(u)
	if err != nil {
		return err
	}
	return nil
}
