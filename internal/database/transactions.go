package database

import (
    "context"
    "time"

    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgtype"
)

func (q *Queries) AddFilmWithActors(ctx context.Context, film AddFilmParams, ids []int32) error {
    conn := q.db.(*pgx.Conn)
    tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
    if err != nil {
        return err
    }
    defer tx.Rollback(ctx)
    qtx := q.WithTx(tx)
    filmId, err := qtx.AddFilm(ctx, film)
    if err != nil {
        return err
    }
    for _, id := range ids {
        err = qtx.AddActorToFilm(ctx, AddActorToFilmParams{id, filmId})
        if err != nil {
            return err
        }
    }
    return tx.Commit(ctx)
}

type OptUpdateFilm struct {
    Id          int32
    Title       *string
    Description *string
    ReleaseDate *time.Time
    Rating      *uint8
}

func (q *Queries) UpdateFilm(ctx context.Context, film OptUpdateFilm) error {
    conn := q.db.(*pgx.Conn)
    tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
    if err != nil {
        return err
    }
    defer tx.Rollback(ctx)
    qtx := q.WithTx(tx)

    if film.Title != nil {
        err = qtx.updateFilmTitle(ctx, updateFilmTitleParams{film.Id, *film.Title})
    }
    if film.Description != nil {
        err = qtx.updateFilmDescription(ctx, updateFilmDescriptionParams{film.Id, *film.Description})
    }
    if film.ReleaseDate != nil {
        err = qtx.updateFilmReleaseDate(ctx,
            updateFilmReleaseDateParams{
                film.Id,
                pgtype.Date{*film.ReleaseDate, 0, true}},
        )
    }
    if err != nil {
        return err

    }
    return tx.Commit(ctx)
}

type OptUpdateActor struct {
    Id       int32
    Name     *string
    Birthday *time.Time
    Gender   *Gender
}
