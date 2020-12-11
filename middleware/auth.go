package middleware

import (
	"github.com/andriipospielov/LoginWebApp/dto"
	"github.com/andriipospielov/LoginWebApp/entity"
	"github.com/andriipospielov/LoginWebApp/repository"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var (
	identityKey       = "id"
	accountRepository = repository.AccountRepository{}.New()
)

func NewJwtAuthenticator() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "Forbidden",
		Key:         []byte("321321321321"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			//println(data)
			if v, ok := data.(*entity.Account); ok {
				return jwt.MapClaims{
					identityKey: v.Id,
					"Login":     v.Login,
					"FirstName": v.FirstName,
					"LastName":  v.LastName,
					"Age":       v.Age,
					"Email":     v.Email,
					"CreatedAt": v.CreatedAt,
					"UpdatedAt": v.UpdatedAt,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return accountRepository.Find(uint64(claims[identityKey].(float64)))
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals dto.Credentials
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if result := accountRepository.FindByCredentials(loginVals); result != nil {
				return result, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {

			if v, ok := data.(*entity.Account); ok {
				c.Set("account", v)
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup: "header: Authorization, query: token, cookie: jwt",

		TokenHeadName: "Bearer",

		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}
