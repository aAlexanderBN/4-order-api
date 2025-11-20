package myuser

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name      string
	Phone     string
	Code      int
	SessionID string
	Token     string
}
