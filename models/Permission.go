package models

import (
	"fmt"
)

type Permission struct {
	ID          uint
	Name        string `gorm:"not null;unique;default:''"`
	Description string `gorm:"not null;default:''"`
}

func (p *Permission) Create() error {
	if !p.isFieldsValid() {
		return fmt.Errorf("权限名不能为空")
	}
	
	var permissions []Permission
	db.Where("name = ?", p.Name).Find(&permissions)
	if len(permissions) != 0 {
		return fmt.Errorf("%s, 该权限已存在", p.Name)
	}

	db.Create(p)
	
	return nil
}

func (p *Permission) isFieldsValid() bool {
	if p.Name == "" {
		return false
	} 
	return true
}

func initPermissions() {
	var permissions []Permission
	db.Find(&permissions)
	if len(permissions) != 0 {
		return
	}
	
	permissions = []Permission {
		Permission {
			Name: "write_book_info",
			Description: "录入图书信息",
		},
		Permission {
			Name: "borrow_book",
			Description: "借书",
		},
		Permission {
			Name: "return_book",
			Description: "还书",
		},
		Permission {
			Name: "see_book_borrow_history",
			Description: "查看历史借阅记录",
		},
		Permission {
			Name: "edit_user",
			Description: "编辑用户",
		},
	}
	
	for _, p := range permissions {
		p.Create()
	}
}

