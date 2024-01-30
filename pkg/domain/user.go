package domain

import "gorm.io/gorm"

type Notifications struct {
	gorm.Model
	SenderID   uint
	ReceiverID uint
	Name       string
	Image      string
	Type       string
	Status     string
	CommonID   uint
}

type FcmTokens struct{
	gorm.Model
	UserID uint
	Token string
}
