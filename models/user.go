package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	UserList map[string]*User
)

type User struct {
	Id         int64     `json:name`
	Name       string    `json:name`
	Age        int       `json:age`
	Email      string    `json:email`
	Created_At time.Time `json:created_at`
	Updated_At time.Time `json:updated_at`
}

func (u *User) TableName() string {
	return TableName("user")
}

//根据用户ID查询用户
func FindUserById(uid int64) (User, error) {

	o := orm.NewOrm()
	u := User{Id: uid}

	err := o.Read(&u)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		return u, err
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return u, err
	} else {
		fmt.Println(u.Id, u.Name)
		return u, err
	}
}
