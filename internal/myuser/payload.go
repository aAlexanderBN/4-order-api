package myuser

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name      string
	Code      int
	sessionID string
	token     string
}
