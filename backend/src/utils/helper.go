package utils

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Compare hashed password
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), plainPwd)
	return err == nil
}

// Session expiration time
func SessionExpires() time.Time {
	return time.Now().Add(24 * time.Hour)
}

// Format Date
func FormatDate(DateOfBirth string) (time.Time, error) {
	format := "2006-01-02"
	dateOfBirth, err := time.Parse(format, DateOfBirth)
	if err != nil {
		return time.Time{}, err
	}
	return dateOfBirth, nil
}


func FormatDateString(DateOfBirth string) (string, error) {
	format := "2006-01-02"
	dateOfBirth, err := time.Parse(format, DateOfBirth)
	if err != nil {
		return "", err
	}
	return dateOfBirth.Format("02 January 2006"), nil
}


// Generate Random Password
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomPassword(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(password)
}

// Generate accunt number
func GenerateAccountNumber() string {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(1_000_000_000) // 9 digit random
	return fmt.Sprintf("92%09d", randomNumber)
}

