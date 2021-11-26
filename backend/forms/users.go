package forms

type RegisterCommand struct {
	Username  string `json:"username" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginCommand struct {
	Email  string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}