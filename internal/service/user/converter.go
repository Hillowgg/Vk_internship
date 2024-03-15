package user

import (
    "github.com/google/uuid"
    "main/internal/database"
)

func dbUserToUser(dbUser *database.User) *User {
    id, _ := uuid.FromBytes(dbUser.ID.Bytes[:])
    user := &User{
        Id:       id,
        Nickname: dbUser.Nickname,
        Email:    dbUser.Email,
        Password: "",
        IsAdmin:  dbUser.IsAdmin.Bool,
    }
    return user
}
