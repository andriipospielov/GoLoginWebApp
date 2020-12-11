package entity

import (
	"time"
)

type RootEntity interface{}

type Account struct {
	Id           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Login        string    `json:"login" binding:"required" gorm:"type:varchar(128);index:idx_name,unique"`
	PasswordHash string    `json:"-" gorm:"type:varchar(128)"`
	Password     string    `json:"password" binding:"required" gorm:"-"`
	FirstName    string    `json:"firstname" binding:"required" gorm:"type:varchar(128)"`
	LastName     string    `json:"lastname" binding:"required" gorm:"type:varchar(128)"`
	Age          int8      `json:"age" binding:"gte=1,lte=130" gorm:"type:varchar(128)"`
	Email        string    `json:"email" binding:"required,email" gorm:"type:varchar(128)"`
	CreatedAt    time.Time `json:"-" gorm:"CURRENT_TIMESTAMP" `
	UpdatedAt    time.Time `json:"-" gorm:"CURRENT_TIMESTAMP" `
}

//func (acccount *Account) BeforeCreate (tx *gorm.DB) (err error) {
//	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(acccount.Password), 5)
//
//	if err != nil {
//		return err
//	}
//
//	//tx.Statement.SetColumn("password_hash", string(HashedPassword))
//	acccount.PasswordHash = string(HashedPassword)
//
//	return
//}
