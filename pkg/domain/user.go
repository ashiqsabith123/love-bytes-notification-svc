package domain

import "gorm.io/gorm"

type Notifications struct {
	gorm.Model
	SenderID   uint
	RecieverID uint
	Name       string
	Image      string
	Type       string
	Status     string
}
