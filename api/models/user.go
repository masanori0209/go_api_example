package models

import (
	"errors"
	"strconv"
	"time"
	"github.com/astaxie/beego/orm"
)

var (
	UserList map[string]*User
)

func init() {
	orm.RegisterModel(new(User), new(Profile))
	UserList = make(map[string]*User)
}

type User struct {
	Id       string   `orm:"column(uid);pk"`
	Username string
	Password string
	Profile  *Profile `orm:"rel(one)"`
}

type Profile struct {
	Id      string   `orm:"column(uid);pk"`
	Gender  string
	Age     int
	Address string
	Email   string
	User    *User `orm:"null;reverse(one);"`
}

func AddUser(u User) string {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	UserList[u.Id] = &u
	return u.Id
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
