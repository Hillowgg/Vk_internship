// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

func (e *Gender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Gender(s)
	case string:
		*e = Gender(s)
	default:
		return fmt.Errorf("unsupported scan type for Gender: %T", src)
	}
	return nil
}

type NullGender struct {
	Gender Gender
	Valid  bool // Valid is true if Gender is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullGender) Scan(value interface{}) error {
	if value == nil {
		ns.Gender, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Gender.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullGender) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Gender), nil
}

type Actor struct {
	ID       int32
	Name     string
	Birthday pgtype.Date
	Gender   Gender
}

type ActorsFilm struct {
	ActorID int32
	FilmID  int32
}

type Film struct {
	ID          int32
	Title       string
	Description string
	ReleaseDate pgtype.Date
	Rating      int16
}

type User struct {
	ID           pgtype.UUID
	Nickname     string
	Email        string
	PasswordHash pgtype.Text
	Salt         pgtype.Text
	IsAdmin      pgtype.Bool
}
