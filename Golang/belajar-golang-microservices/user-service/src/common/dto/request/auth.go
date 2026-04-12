package dto

type Register struct {
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"full_name" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}


type SendOTP struct {
	Email string `json:"email" validate:"required,email"`
	Otp string `json:"otp" validate:"required,len=6"`

}