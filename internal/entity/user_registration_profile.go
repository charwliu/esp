package entity

import (
	"time"
)

type UserRegistrationProfile struct {
	Id                  int `gorm:"primary_key;autoIncrement"`
	ActivationKey       string
	ActivationExpires   time.Time
	ResetKey            string
	ResetExpires        time.Time
	VerificationKey     string
	VerificationExpires time.Time
	UserEmail           string
}

func (UserRegistrationProfile) TableName() string {
	return "user_registration_profile"
}
