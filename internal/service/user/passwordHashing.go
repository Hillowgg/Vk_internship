package user

import (
    "crypto/rand"
    "crypto/sha256"
    "fmt"
)

var saltSize = 32

func generateRandomSalt() (string, error) {
    var salt = make([]byte, saltSize)

    _, err := rand.Read(salt[:])

    if err != nil {
        return "", err
    }

    return fmt.Sprintf("%x", salt), nil
}

func hashPassword(pass string, salt string) string {
    bytes := []byte(pass + salt)

    sum := sha256.Sum256(bytes)
    return fmt.Sprintf("%x", sum)
}

func generateHashedPassword(pass string) (string, string, error) {
    salt, err := generateRandomSalt()
    if err != nil {
        return "", "", err
    }
    hash := hashPassword(pass, salt)
    return hash, salt, err
}

func checkPassword(pass, salt, hash string) bool {
    right := hashPassword(pass, salt)
    return hash == right
}
