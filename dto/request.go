package dto

type Dto interface{}
type Request interface {
	Dto
}

type Credentials struct {
	Login    string `json:"login" binding:"required" `
	Password string `json:"password" binding:"required"`
}
