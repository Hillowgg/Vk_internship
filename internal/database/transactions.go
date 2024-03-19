package database

import (
    "context"
    "errors"
    "time"

    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgtype"
    "github.com/jackc/pgx/v5/pgxpool"
)

func (q *Queries) AddFilmWithActors(ctx context.Context, film *AddFilmParams, ids []int32) (int32, error) {
    conn := q.db.(*pgxpool.Pool)
    tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
    if err != nil {
        return 0, err
    }
    defer tx.Rollback(ctx)
    qtx := q.WithTx(tx)
    filmId, err := qtx.AddFilm(ctx, film)
    if err != nil {
        return 0, err
    }
    for _, id := range ids {
        err = qtx.AddActorToFilm(ctx, &AddActorToFilmParams{id, filmId})
        if err != nil {
            return 0, err
        }
    }
    return filmId, tx.Commit(ctx)
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
        t, err := time.Parse("2006-01-02", d)
        if err != nil {
            return errors.New("release date has wrong format")
        }
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

func (q *Queries) UpdateActor(ctx context.Context, actor map[string]any) error {
    conn := q.db.(*pgxpool.Pool)
    tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
    if err != nil {
        return err
    }
    defer tx.Rollback(ctx)
    qtx := q.WithTx(tx)

    id := actor["Id"].(int32)

    if name, ok := actor["Name"]; ok {
        n, ok := name.(string)
        if !ok {
            return errors.New("name has wrong type")
        }
        err = qtx.updateActorName(ctx, &updateActorNameParams{id, n})
    }
    if birthday, ok := actor["Birthday"]; ok {
        b, ok := birthday.(string)
        if !ok {
            return errors.New("birthday has wrong type")
        }
        t, err := time.Parse("2006-01-02", b)
        if err != nil {
            return errors.New("birthday has wrong format")
        }
        err = qtx.updateActorBirthday(
            ctx,
            &updateActorBirthdayParams{
                ID:       id,
                Birthday: pgtype.Date{Time: t, InfinityModifier: pgtype.Finite, Valid: true}},
        )
    }
    if gender, ok := actor["Gender"]; ok {
        var g Gender
        err := g.Scan(gender)
        if err != nil {
            return errors.New("gender has wrong type")
        }
        err = qtx.updateActorGender(ctx, &updateActorGenderParams{id, g})
    }
    if err != nil {
        return err
    }
    return tx.Commit(ctx)
}
