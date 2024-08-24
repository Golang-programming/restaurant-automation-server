package service

import (
	"errors"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/net/context"
)

var (
	jwtSecret   = []byte(os.Getenv("JWT_SECRET"))
	redisClient *redis.Client
	ctx         = context.Background()
)

type CustomClaims struct {
	StaffID string `json:"staff_id"`
	Email   string `json:"email"`
	jwt.RegisteredClaims
}

func InitRedisClient() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})
}

func AssignToken(staffID, email string) (string, string, error) {
	accessTokenStr, refreshTokenStr, err := createTokens(staffID, email)

	return accessTokenStr, refreshTokenStr, err
}

func ValidateToken(tokenStr string) (*CustomClaims, error) {
	claims, err := decodeToken(tokenStr)
	if err != nil {
		return nil, err
	}
	if time.Now().After(claims.ExpiresAt.Time) {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}

func RefreshToken(refreshTokenStr string) (string, string, error) {
	claims, err := decodeToken(refreshTokenStr)
	if err != nil {
		return "", "", err
	}

	storedRefreshToken, err := redisClient.Get(ctx, "refresh_token:"+claims.StaffID).Result()
	if err != nil || storedRefreshToken != refreshTokenStr {
		return "", "", errors.New("invalid refresh token")
	}

	newAccessTokenStr, newRefreshTokenStr, err := createTokens(claims.StaffID, claims.Email)

	return newAccessTokenStr, newRefreshTokenStr, err
}

func RevokeToken(staffID string) error {
	err := redisClient.Del(ctx, "refresh_token:"+staffID).Err()
	if err != nil {
		return err
	}

	return nil
}

func decodeToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func createTokens(staffID, email string) (string, string, error) {
	expirationTime := time.Now().Add(10 * time.Minute)
	refreshExpirationTime := time.Now().Add(24 * time.Hour)

	claims := CustomClaims{
		StaffID: staffID,
		Email:   email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshClaims := CustomClaims{
		StaffID: staffID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessTokenStr, err := accessToken.SignedString("jwtSecret")
	if err != nil {
		return "", "", err
	}
	refreshTokenStr, err := refreshToken.SignedString("jwtSecret")
	if err != nil {
		return "", "", err
	}

	err = redisClient.Set(ctx, "refresh_token:"+staffID, refreshTokenStr, 24*time.Hour).Err()
	if err != nil {
		return "", "", err
	}

	return accessTokenStr, refreshTokenStr, nil
}
