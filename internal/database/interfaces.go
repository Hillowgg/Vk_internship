package database

import (
    "context"

    "github.com/jackc/pgx/v5/pgtype"
)

type QuerierWithTx interface {
    Querier
    UpdateActor(ctx context.Context, actor map[string]any) error
    DeleteActorById(ctx context.Context, id int32) error
    AddFilmWithActors(ctx context.Context, film *AddFilmParams, ids []int32) error
    UpdateFilm(ctx context.Context, film map[string]any) error
    DeleteFilmById(ctx context.Context, id int32) error
}

type ActorQuerierWithTx interface {
    AddActor(ctx context.Context, arg *AddActorParams) (int32, error)
    AddActorToFilm(ctx context.Context, arg *AddActorToFilmParams) error
    GetActorById(ctx context.Context, id int32) (*Actor, error)
    GetActorsWithFilms(ctx context.Context) ([]*GetActorsWithFilmsRow, error)
    SearchActorsByName(ctx context.Context, dollar_1 pgtype.Text) ([]*Actor, error)
    UpdateActor(ctx context.Context, actor map[string]any) error
    DeleteActorById(ctx context.Context, id int32) error
}

type FilmQuerierWithTx interface {
    AddFilm(ctx context.Context, arg *AddFilmParams) (int32, error)
    AddFilmWithActors(ctx context.Context, film *AddFilmParams, ids []int32) (int32, error)
    GetFilmById(ctx context.Context, id int32) (*Film, error)
    GetFilmsASCRating(ctx context.Context) ([]*Film, error)
    GetFilmsASCReleaseDate(ctx context.Context) ([]*Film, error)
    GetFilmsASCTitle(ctx context.Context) ([]*Film, error)
    GetFilmsDESCRating(ctx context.Context) ([]*Film, error)
    GetFilmsDESCReleaseDate(ctx context.Context) ([]*Film, error)
    GetFilmsDESCTitle(ctx context.Context) ([]*Film, error)
    SearchFilmByTitleAndActor(ctx context.Context, arg *SearchFilmByTitleAndActorParams) (*Film, error)
    SearchFilmsByTitle(ctx context.Context, dollar_1 pgtype.Text) ([]*Film, error)
    UpdateFilm(ctx context.Context, film map[string]any) error
    DeleteFilmById(ctx context.Context, id int32) error
}

type UserQuerierWithTx interface {
    AddUser(ctx context.Context, arg *AddUserParams) error
    GetUserByEmail(ctx context.Context, email string) (*User, error)
    GetUserByLogin(ctx context.Context, nickname string) (*User, error)
    GetUserById(ctx context.Context, id pgtype.UUID) (*User, error)
    DeleteUser(ctx context.Context, id pgtype.UUID) error
}
