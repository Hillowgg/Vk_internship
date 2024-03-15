package database

import (
    "context"

    "github.com/jackc/pgx/v5"
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
