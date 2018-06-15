package models

import (
	"fmt"
)

type User struct {
	ID         uint
	Username   string `gorm:"not null;unique;default:''"`
	Password   string `gorm:"not null;default:''"`
	Fullname   string `gorm:"not null;default:''"`
	Center     string `gorm:"not null;default:''"`
	Department string `gorm:"not null;default:''"`
	Position   string `gorm:"not null;default:''"`
	IsLogin    bool   `gorm:"not null;default:false"`
	Roles      []Role `gorm:"many2many:user_roles;"`
}

func (u *User) Create() error {
	if !u.isFieldsValid() {
		return fmt.Errorf("用户信息不能为空")
	}
	
	var users []User
	db.Where("username = ?", u.Username).Find(&users)
	if len(users) != 0 {
		return fmt.Errorf("%s, 该用户已存在", u.Username)
	}

	db.Create(u)
	
	return nil
}

func (user *User) Delete() {
}

func (user *User) isFieldsValid() bool {
	if user.Username == "" || user.Password == "" ||
		user.Fullname == "" || user.Center == "" ||
		user.Department == "" || user.Position == "" {
		return false
	} else {
		return true
	}
}

func initAdminUser() {
	var adminUser User
	if !db.Where("username=?", "admin").First(&adminUser).RecordNotFound() {
		return
	}
	
	var adminRole Role
	if db.Where("name=?", "管理员").First(&adminRole).RecordNotFound() {
		return
	}
	
	adminUser = User {
		Username: "admin",
		Password: "admin123",
		Fullname: "admin",
		Center: "山西方是科技有限公司",
		Department: "山西方是科技有限公司",
		Position: "领导",
		IsLogin: false,
		Roles: []Role {
			adminRole,
		},
	}
	
	adminUser.Create()
}
