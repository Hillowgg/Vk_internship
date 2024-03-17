// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddActor(ctx context.Context, arg *AddActorParams) (int32, error)
	AddActorToFilm(ctx context.Context, arg *AddActorToFilmParams) error
	AddFilm(ctx context.Context, arg *AddFilmParams) (int32, error)
	AddUser(ctx context.Context, arg *AddUserParams) error
	DeleteActorById(ctx context.Context, id int32) error
	DeleteFilmById(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id pgtype.UUID) error
	//ACTORS----------------------------------------------------------------------------------------------------------------
	GetActorById(ctx context.Context, id int32) (*Actor, error)
	//FILMS-----------------------------------------------------------------------------------------------------------------
	GetFilmById(ctx context.Context, id int32) (*Film, error)
	GetFilmsASCRating(ctx context.Context) ([]*Film, error)
	GetFilmsASCReleaseDate(ctx context.Context) ([]*Film, error)
	GetFilmsASCTitle(ctx context.Context) ([]*Film, error)
	GetFilmsDESCRating(ctx context.Context) ([]*Film, error)
	GetFilmsDESCReleaseDate(ctx context.Context) ([]*Film, error)
	GetFilmsDESCTitle(ctx context.Context) ([]*Film, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	//USERS-----------------------------------------------------------------------------------------------------------------
	GetUserById(ctx context.Context, id pgtype.UUID) (*User, error)
	GetUserByLogin(ctx context.Context, nickname string) (*User, error)
	SearchActorsByName(ctx context.Context, dollar_1 pgtype.Text) ([]*Actor, error)
	SearchFilmByTitleAndActor(ctx context.Context, arg *SearchFilmByTitleAndActorParams) (*Film, error)
	SearchFilmsByTitle(ctx context.Context, dollar_1 pgtype.Text) ([]*Film, error)
	updateActorBirthday(ctx context.Context, arg *updateActorBirthdayParams) error
	updateActorGender(ctx context.Context, arg *updateActorGenderParams) error
	updateActorName(ctx context.Context, arg *updateActorNameParams) error
	updateFilmDescription(ctx context.Context, arg *updateFilmDescriptionParams) error
	updateFilmRating(ctx context.Context, arg *updateFilmRatingParams) error
	updateFilmReleaseDate(ctx context.Context, arg *updateFilmReleaseDateParams) error
	updateFilmTitle(ctx context.Context, arg *updateFilmTitleParams) error
}

var _ Querier = (*Queries)(nil)
