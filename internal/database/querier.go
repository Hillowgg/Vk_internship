// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddActor(ctx context.Context, arg AddActorParams) (int32, error)
	AddActorToFilm(ctx context.Context, arg AddActorToFilmParams) error
	AddFilm(ctx context.Context, arg AddFilmParams) (int32, error)
	// todo: fix deleting from actors_films
	DeleteActorById(ctx context.Context, id int32) error
	// todo: fix deleting from actors_films
	DeleteFilmById(ctx context.Context, id int32) error
	//ACTORS----------------------------------------------------------------------------------------------------------------
	GetActorById(ctx context.Context, id int32) (Actor, error)
	//FILMS-----------------------------------------------------------------------------------------------------------------
	GetFilmById(ctx context.Context, id int32) (Film, error)
	SearchActorsByName(ctx context.Context, dollar_1 pgtype.Text) ([]Actor, error)
	SearchFilmByTitleAndActor(ctx context.Context, arg SearchFilmByTitleAndActorParams) (Film, error)
	SearchFilmsByTitle(ctx context.Context, dollar_1 pgtype.Text) ([]Film, error)
	updateActorBirthday(ctx context.Context, arg updateActorBirthdayParams) error
	updateActorGender(ctx context.Context, arg updateActorGenderParams) error
	updateActorName(ctx context.Context, arg updateActorNameParams) error
	updateFilmDescription(ctx context.Context, arg updateFilmDescriptionParams) error
	updateFilmRating(ctx context.Context, arg updateFilmRatingParams) error
	updateFilmReleaseDate(ctx context.Context, arg updateFilmReleaseDateParams) error
	updateFilmTitle(ctx context.Context, arg updateFilmTitleParams) error
}

var _ Querier = (*Queries)(nil)
