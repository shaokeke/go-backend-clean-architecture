package domain

import (
	"context"
)

type User struct {
	ID       uint   `json:"id" gorm:"primary_key;auto_increment"`
	Name     string `json:"name" gorm:"type:varchar(128);not null"`
	Email    string `json:"email" gorm:"type:varchar(128);not null"`
	Password string `json:"password" gorm:"type:varchar(128);not null"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id uint) (User, error)
}
