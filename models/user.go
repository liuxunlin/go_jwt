package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int    `json:id`
	Name       string `json:name`
	Email      string `json:email`
	Age        int    `json:age`
	Created_at int    `json:created_at`
	Updated_at int    `json:updated_at`
}

func (m *User) TableName() string {
	return TableName("user")
}

// func init() {
// 	orm.RegisterModel(new(User))
// }

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
