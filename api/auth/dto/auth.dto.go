package dto

type SendOTPInput struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type VerifyOTPInput struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	OTP         string `json:"otp" validate:"required"`
}
