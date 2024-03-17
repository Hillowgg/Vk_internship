package database

import (
    "context"
    "errors"
    "time"

    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgtype"
    "github.com/jackc/pgx/v5/pgxpool"
)

type QuerierWithTx interface {
    Querier
    AddFilmWithActors(ctx context.Context, film *AddFilmParams, ids []int32) error
    UpdateFilm(ctx context.Context, film map[string]any) error
    UpdateActor(ctx context.Context, actor *OptUpdateActor) error
    DeleteFilmById(ctx context.Context, id int32) error

    DeleteActorById(ctx context.Context, id int32) error
}

func (q *Queries) AddFilmWithActors(ctx context.Context, film *AddFilmParams, ids []int32) error {
    conn := q.db.(*pgxpool.Pool)
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
        err = qtx.AddActorToFilm(ctx, &AddActorToFilmParams{id, filmId})
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

func (q *Queries) UpdateFilm(ctx context.Context, film map[string]any) error {
    conn := q.db.(*pgxpool.Pool)
    tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
    if err != nil {
        return err
    }
    defer tx.Rollback(ctx)
    qtx := q.WithTx(tx)

    id := film["Id"].(int32)

    if title, ok := film["Title"]; ok {
        t, ok := title.(string)
        if !ok {
            return errors.New("title has wrong type")
        }
        err = qtx.updateFilmTitle(ctx, &updateFilmTitleParams{id, t})
    }
    if desc, ok := film["Description"]; ok {
        d, ok := desc.(string)
        if !ok {
            return errors.New("description has wrong type")
        }
        err = qtx.updateFilmDescription(ctx, &updateFilmDescriptionParams{id, d})
    }
    if date, ok := film["ReleaseDate"]; ok {
        d, ok := date.(string)
        if !ok {
            return errors.New("release date has wrong type")
        }
        t, _ := time.Parse("2006-01-02", d)
        err = qtx.updateFilmReleaseDate(ctx,
            &updateFilmReleaseDateParams{
                id,
                pgtype.Date{t, pgtype.Finite, true}},
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

func (q *Queries) UpdateActor(ctx context.Context, actor *OptUpdateActor) error {
    conn := q.db.(*pgxpool.Pool)
    tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
    if err != nil {
        return err
    }
    defer tx.Rollback(ctx)
    qtx := q.WithTx(tx)

    if actor.Name != nil {
        err = qtx.updateActorName(ctx, &updateActorNameParams{actor.Id, *actor.Name})
    }
    if actor.Birthday != nil {
        err = qtx.updateActorBirthday(
            ctx,
            &updateActorBirthdayParams{
                ID:       actor.Id,
                Birthday: pgtype.Date{Time: *actor.Birthday, InfinityModifier: pgtype.Finite, Valid: true}},
        )
    }
    if actor.Gender != nil {
        err = qtx.updateActorGender(ctx, &updateActorGenderParams{actor.Id, *actor.Gender})
    }
    if err != nil {
        return err
    }
    return tx.Commit(ctx)
}

func (q *Queries) DeleteFilmById(ctx context.Context, id int32) error {
    conn := q.db.(*pgxpool.Pool)
    tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
    if err != nil {
        return err
    }
    defer tx.Rollback(ctx)
    qtx := q.WithTx(tx)
    err = qtx.deleteConnectionsByFilmId(ctx, id)
    if err != nil {
        return err
    }
    err = q.deleteFilmById(ctx, id)
    if err != nil {
        return err
    }
    return tx.Commit(ctx)
}

func (q *Queries) DeleteActorById(ctx context.Context, id int32) error {
    conn := q.db.(*pgx.Conn)
    tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
    if err != nil {
        return err
    }
    defer tx.Rollback(ctx)
    qtx := q.WithTx(tx)
    err = qtx.deleteConnectionsByActorId(ctx, id)
    if err != nil {
        return err
    }
    err = q.deleteActorById(ctx, id)
    if err != nil {
        return err
    }
    return tx.Commit(ctx)
}
