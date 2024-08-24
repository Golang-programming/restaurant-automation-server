package dto

type SendOTPInput struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type ValidateOTPInput struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	OTP         string `json:"otp" validate:"required"`
}
