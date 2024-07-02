package utils

import (
	"os"
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GenerateUUID() string {
	uuid := uuid.New()
	return uuid.String()
}

func ValidateUser(
	firstName,
	lastName,
	userName,
	password,
	confirmPassword string,
) (bool, string) {
	if len(strings.TrimSpace(password)) < 8 {
		return false, "Password must be at least 8 characters long"
	}

	if len(strings.TrimSpace(userName)) < 3 || len(strings.TrimSpace(userName)) > 20 {
		return false, "Username must be between 3 and 20 characters"
	}

	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return false, "Password must contain at least one number"
	}

	if !regexp.MustCompile(`[!@#~$%^&*(),.?":{}|<>]`).MatchString(password) {
		return false, "Password must contain at least one special character"
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(userName) {
		return false, "Username must only contain alphanumeric characters"
	}

	if strings.Contains(password, firstName) || strings.Contains(password, lastName) {
		return false, "Password must not contain your first or last name"
	}

	if len(strings.TrimSpace(firstName)) < 1 {
		return false, "First name must be provided"
	}

	if len(strings.TrimSpace(lastName)) < 1 {
		return false, "Last name must be provided"
	}

	if len(strings.TrimSpace(confirmPassword)) < 1 {
		return false, "Confirm password must be provided"
	}

	if password != confirmPassword {
		return false, "Passwords do not match"
	}

	if userName == password {
		return false, "Username and password must not be the same"
	}

	return true, ""
}

func ComparePassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return err
	}
	return nil
}

func CreateToken(username string) (string, error) {
	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().AddDate(0, 0, 7).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
