package datavase

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	name  string
	votes int
	users []*User `gorm:"many2many:user_role;"`
}

type User struct {
	gorm.Model
	id    string
	name  string
	roles []*User
}
