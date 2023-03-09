package services

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// CheckPassword checks if a plaintext password matches a hashed password.
func CheckPassword(password, hashedPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}
