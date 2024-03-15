package user

import (
    "crypto/rand"
    "crypto/sha256"
)

var saltSize = 64

func generateRandomSalt() ([]byte, error) {
    var salt = make([]byte, saltSize)

    _, err := rand.Read(salt[:])

    if err != nil {
        return nil, err
    }

    return salt, nil
}

func hashPassword(pass string, salt []byte) string {
    passBytes := []byte(pass)

    sum := sha256.Sum256(append(passBytes, salt...))
    return string(sum[:])
}

func generateHashedPassword(pass string) (string, string, error) {
    salt, err := generateRandomSalt()
    if err != nil {
        return "", "", err
    }
    hash := hashPassword(pass, salt)
    return hash, string(salt), err
}

func checkPassword(pass, salt, hash string) bool {
    right := hashPassword(pass, []byte(salt))
    return hash == right
}
