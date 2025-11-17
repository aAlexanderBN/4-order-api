package myUser

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name string
	Code int
}
