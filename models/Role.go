package models

import (
	"fmt"
)

type Role struct {
	ID uint
	Name string `gorm:"not null;unique;default:''"`
	Description string `gorm:"not null;default:''"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

func (r *Role) isFieldsValid() bool {
	if r.Name == "" {
		return false
	} 
	return true
}

func (r *Role) Create() error {
	if !r.isFieldsValid() {
		return fmt.Errorf("角色名不能为空")
	}
	
	var roles []Role
	db.Where("name = ?", r.Name).Find(&roles)
	if len(roles) != 0 {
		return fmt.Errorf("%s, 该角色已存在", r.Name)
	}

	db.Create(r)
	
	return nil
}

func (r *Role) Update() error {
	if !r.isFieldsValid() {
		return fmt.Errorf("角色名不能为空")
	}
	
	var roles []Role
	db.Where("name = ?", r.Name).Find(&roles)
	if len(roles) == 0 {
		return fmt.Errorf("%s, 该角色尚不存在", r.Name)
	}
	
	db.Save(r)
	
	return nil
}

func initRoles() {
	var roles []Role
	db.Find(&roles)
	if len(roles) != 0 {
		return
	}
	
	var permissions []Permission
	db.Find(&permissions)
	if len(permissions) == 0 {
		return
	}
	
	role := Role {
		Name: "管理员",
		Description: "管理员",
		Permissions: permissions,
	}
	role.Create()
	
	role = Role {
		Name: "访客",
		Description: "访客",
	}
	role.Create()
	
	var commonUserPermissions []Permission
	for _, p := range permissions {
		if p.Name == "borrow_book" {
			commonUserPermissions = append(commonUserPermissions, p)
			break
		}
	}
	role = Role {
		Name: "普通用户",
		Description: "普通用户",
		Permissions: commonUserPermissions,
	}
	role.Create()
}

