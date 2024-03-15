package actor

import (
    "github.com/jackc/pgx/v5/pgtype"
    "main/internal/database"
)

func newActorToDb(actor *NewActor) *database.AddActorParams {
    ret := &database.AddActorParams{
        Name:     actor.Name,
        Birthday: pgtype.Date{Time: actor.Birthday, InfinityModifier: pgtype.Finite, Valid: true},
        Gender:   actor.Gender,
    }
    return ret
}

func dbActorToActor(dbActor *database.Actor) *Actor {
    ret := &Actor{
        Id:       dbActor.ID,
        Name:     dbActor.Name,
        Birthday: dbActor.Birthday.Time,
        Gender:   dbActor.Gender,
    }
    return ret
}
