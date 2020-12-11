package controller

import (
	"github.com/andriipospielov/LoginWebApp/entity"
	"github.com/andriipospielov/LoginWebApp/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

type Controller interface {
	New() Controller
}

type CrudController interface {
	Controller
	Index() []entity.RootEntity
	Create(ctx *gin.Context) (entity.RootEntity, error)
}

type AccountController struct {
	CrudController
	AccountRepository repository.AccountRepository
}

func (c AccountController) New() *AccountController {
	return &AccountController{AccountRepository: repository.AccountRepository{}.New()}
}

func (c AccountController) Index(ctx *gin.Context) {
	accounts := c.AccountRepository.FindAll()
	ctx.JSON(http.StatusOK, gin.H{
		"accounts": accounts,
	})

}

func (c AccountController) Create(ctx *gin.Context) {
	var account entity.Account
	err := ctx.BindJSON(&account)

	account, err = hashPasswordForAccount(account)

	if err != nil {
		ctx.JSON(500, gin.H{})
		log.Fatal(err.Error())
		return
	}

	err = c.AccountRepository.Save(account)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, account)
}

func (c AccountController) Update(ctx *gin.Context) {
	v, err := strconv.ParseUint(ctx.Param("id"), 0, 0)

	v2, exists := ctx.Get("account")
	authorizedAcc, converts := v2.(*entity.Account)
	if !(exists && converts && authorizedAcc.Id == v) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "access denied"})
		return
	}

	var account entity.Account
	err = ctx.BindJSON(&account)

	if err != nil {
		ctx.JSON(500, gin.H{})
		log.Fatal(err)
	}
	account.Id = v

	account, err = hashPasswordForAccount(account)
	if err != nil {
		ctx.JSON(500, gin.H{})
		log.Fatal(err)
	}

	c.AccountRepository.Update(account)

	ctx.JSON(http.StatusAccepted, account)
}
func hashPasswordForAccount(account entity.Account) (entity.Account, error) {
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), 5)

	if err != nil {
		log.Fatal(err)
		return account, err
	}

	account.PasswordHash = string(HashedPassword)
	return account, nil
}
