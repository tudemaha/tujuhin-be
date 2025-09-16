package dto

type UserRegister struct {
	Name           string `json:"name" validate:"required,max=50"`
	Username       string `json:"username" validate:"required,max=50"`
	Password       string `json:"password" validate:"required,min=8,max=50"`
	PasswordRepeat string `json:"password_repeat" validate:"required,min=8,max=50,eqfield=Password"`
}

type UserLogin struct {
	Username string `json:"username" validate:"required,max=50"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

type GenerateAccess struct {
	RefreshToken string `json:"refresh_token" validate:"required,min=100"`
}
