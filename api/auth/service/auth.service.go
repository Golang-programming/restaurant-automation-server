package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.co/golang-programming/restaurant/api/auth/dto"
	"github.co/golang-programming/restaurant/api/entity"
	"github.co/golang-programming/restaurant/api/redis"
	userService "github.co/golang-programming/restaurant/api/user/service"
)

const (
	otpTTL       = 2 * time.Minute
	requestLimit = 3
	blockTime    = 2 * time.Minute
)

func SendOTP(input *dto.SendOTPInput) error {
	phoneKey := fmt.Sprintf("otp:%s", input.PhoneNumber)
	reqKey := fmt.Sprintf("otp_req:%s", input.PhoneNumber)

	if isBlocked(reqKey) {
		return errors.New("you are blocked for 2 minutes")
	}

	if err := incrementRequestCount(reqKey); err != nil {
		return err
	}

	otp := generateOTP()
	if err := storeOTP(phoneKey, otp); err != nil {
		return err
	}

	fmt.Printf("Sending OTP %s to %s\n", otp, input.PhoneNumber)

	return nil
}

func ValidateOTP(input *dto.ValidateOTPInput) error {
	phoneKey := fmt.Sprintf("otp:%s", input.PhoneNumber)

	storedOTP, err := redis.Get(phoneKey)
	if err != nil || storedOTP == "" {
		return errors.New("OTP expired or invalid")
	}

	if storedOTP != input.OTP {
		return errors.New("incorrect OTP")
	}

	if err := redis.Del(phoneKey); err != nil {
		return err
	}

	return nil
}

func LoginUser(phoneNumber string) (*entity.User, string, string, error) {
	user, err := userService.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, "", "", err
	}

	accessToken, refreshToken, err := AssignToken(string(user.ID), user.PhoneNumber)
	return user, accessToken, refreshToken, err
}

func isBlocked(reqKey string) bool {
	ttl, err := redis.TTL(reqKey)
	return err == nil && ttl > 0
}

func incrementRequestCount(reqKey string) error {
	count, err := redis.Increment(reqKey)
	if err != nil {
		return err
	}

	if count == 4 {
		return redis.SetExpiration(reqKey, blockTime)
	}

	return nil
}

func generateOTP() string {
	// Simple 6-digit OTP generation logic
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func storeOTP(phoneKey, otp string) error {
	return redis.Set(phoneKey, otp, otpTTL)
}
