package dto

type LoginDTO struct {
	UserID   uint
	UserName string `json:"username"`
	Token    string `json:"token"`
}

type LoginReq struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
