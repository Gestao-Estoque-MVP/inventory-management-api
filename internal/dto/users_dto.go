package dto

type UserCreateDTO struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Document    string `json:"document" binding:"required"`
	MobilePhone string `json:"mobile_phone" binding:"required"`
}
