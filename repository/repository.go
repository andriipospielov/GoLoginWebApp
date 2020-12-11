package repository

import (
	"errors"
	"fmt"
	"github.com/andriipospielov/LoginWebApp/dto"
	"github.com/andriipospielov/LoginWebApp/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Repository interface {
	New() Repository
	CloseConnection()
}

type EntityRepository interface {
	Repository
}

type AccountRepository struct {
	EntityRepository
	connection *gorm.DB
}

func (r AccountRepository) FindByCredentials(c dto.Credentials) interface{} {
	var result entity.Account

	r.connection.Where(&entity.Account{Login: c.Login}).First(&result)

	if result.Id == 0 {
		return nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.PasswordHash), []byte(c.Password)); err != nil {
		return nil
	}

	return &result
}

func (r AccountRepository) Save(a entity.Account) error {
	r.connection.Save(&a)

	if a.Id == 0 {
		return errors.New("could not save user with given parameters")
	}
	return nil
}

func (r AccountRepository) Update(a entity.Account) {
	r.connection.Save(&a)
}

func (r AccountRepository) FindAll() []entity.Account {
	var accounts []entity.Account

	r.connection.Find(&accounts)
	return accounts
}

func (r AccountRepository) New() AccountRepository {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Sprintf("DB connection failed :%s", err))
	}

	//AutoMigrate here, just for simplicity
	err = db.AutoMigrate(&entity.Account{})
	if err != nil {
		log.Fatal(fmt.Sprintf("Automigrating failed failed :%s", err))
	}

	return AccountRepository{
		connection: db,
	}
}

func (r *AccountRepository) CloseConnection() {
	db, err := r.connection.DB()

	if err != nil {
		log.Fatal(fmt.Sprintf("DB connection closing failed :%s", err))
	}

	err = db.Close()
	if err != nil {
		log.Fatal(fmt.Sprintf("DB connection closing failed :%s", err))
	}
}

func (r AccountRepository) Find(id uint64) interface{} {
	account := &entity.Account{}
	r.connection.First(account, id)
	if account.Id == 0 {
		return nil
	}
	return account
}
