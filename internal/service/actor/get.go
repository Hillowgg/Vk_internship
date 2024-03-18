package actor

import (
    "context"

    "github.com/jackc/pgx/v5/pgtype"
    "main/internal/logs"
    "main/internal/service/film"
)

func (s *Service) GetActor(ctx context.Context, id int32) (*Actor, error) {
    actor, err := s.db.GetActorById(ctx, id)
    if err != nil {
        logs.Log.Errorw("Failed to get an actor", "actorId", id, "err", err)
        return nil, err
    }
    logs.Log.Infow("Got an actor", "actorId", id)
    return dbActorToActor(actor), nil
}

func (s *Service) SearchActors(ctx context.Context, name string) ([]*Actor, error) {
    actors, err := s.db.SearchActorsByName(ctx, pgtype.Text{name, true})
    if err != nil {
        logs.Log.Errorw("Failed to search actors", "info", name, "err", err)
        return nil, err
    }
    logs.Log.Infow("Searched actors", "info", name)
    ret := make([]*Actor, 0, len(actors))
    for _, actor := range actors {
        ret = append(ret, dbActorToActor(actor))
    }
    return ret, nil
}

func (s *Service) GetActorsWithFilms(ctx context.Context) (map[Actor][]film.Film, error) {
    dbRows, err := s.db.GetActorsWithFilms(ctx)
    if err != nil {
        logs.Log.Errorw("Failed to get actors with films", "err", err)
        return nil, err
    }
    logs.Log.Infow("Got actors with films", "rows", len(dbRows))
    ret := make(map[Actor][]film.Film, len(dbRows)/2)
    for _, row := range dbRows {
        a := Actor{row.ID, row.Name, row.Birthday.Time, row.Gender}
        if row.FilmID.Valid {
            f := film.Film{Id: row.FilmID.Int32, Title: row.Title.String, Description: row.Description.String,
                ReleaseDate: row.ReleaseDate.Time, Rating: int8(row.Rating.Int16)}
            ret[a] = append(ret[a], f)
        } else {
            ret[a] = []film.Film{}
        }
    }
    return ret, nil
}
