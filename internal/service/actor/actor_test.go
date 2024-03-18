package actor

import (
    "context"
    "testing"
    "time"

    "github.com/jackc/pgx/v5/pgtype"
    "github.com/stretchr/testify/mock"
    "main/internal/database"
)

type mockDb struct {
    mock.Mock
}

func (m *mockDb) AddActor(ctx context.Context, actor *database.AddActorParams) (int32, error) {
    args := m.Called(ctx, actor)
    return int32(args.Int(0)), args.Error(1)
}

func (m *mockDb) AddActorToFilm(ctx context.Context, arg *database.AddActorToFilmParams) error {
    args := m.Called(ctx, arg)
    return args.Error(0)
}

func (m *mockDb) GetActorById(ctx context.Context, id int32) (*database.Actor, error) {
    args := m.Called(ctx, id)
    return args.Get(0).(*database.Actor), args.Error(1)
}

func (m *mockDb) GetActorsWithFilms(ctx context.Context) ([]*database.GetActorsWithFilmsRow, error) {
    args := m.Called(ctx)
    return args.Get(0).([]*database.GetActorsWithFilmsRow), args.Error(1)
}

func (m *mockDb) SearchActorsByName(ctx context.Context, dollar_1 pgtype.Text) ([]*database.Actor, error) {
    args := m.Called(ctx, dollar_1)
    return args.Get(0).([]*database.Actor), args.Error(1)
}

func (m *mockDb) UpdateActor(ctx context.Context, actor map[string]any) error {
    args := m.Called(ctx, actor)
    return args.Error(0)
}

func (m *mockDb) DeleteActorById(ctx context.Context, id int32) error {
    args := m.Called(ctx, id)
    return args.Error(0)

}
func TestService_CreateActor(t *testing.T) {
    c := new(mockDb)
    s := NewService(c)
    c.On("AddActor", mock.Anything, mock.Anything).Return(1, nil)
    _, err := s.CreateActor(context.Background(), &NewActor{"name", time.Now(), "male"})
    if err != nil {
        t.Error("CreateActor returned error")
    }
}

func TestService_GetActor(t *testing.T) {
    c := new(mockDb)
    s := NewService(c)
    c.On("GetActorById", mock.Anything, mock.Anything).Return(&database.Actor{}, nil)
    _, err := s.GetActor(context.Background(), 1)
    if err != nil {
        t.Error("GetActor returned error")
    }
}

func TestService_GetActorsWithFilms(t *testing.T) {
    c := new(mockDb)
    s := NewService(c)
    c.On("GetActorsWithFilms", mock.Anything).Return([]*database.GetActorsWithFilmsRow{}, nil)
    _, err := s.GetActorsWithFilms(context.Background())
    if err != nil {
        t.Error("GetActorsWithFilms returned error")
    }
}

func TestService_SearchActors(t *testing.T) {
    c := new(mockDb)
    s := NewService(c)
    c.On("SearchActorsByName", mock.Anything, mock.Anything).Return([]*database.Actor{}, nil)
    _, err := s.SearchActors(context.Background(), "name")
    if err != nil {
        t.Error("SearchActors returned error")
    }
}

func TestService_UpdateActor(t *testing.T) {
    c := new(mockDb)
    s := NewService(c)
    c.On("UpdateActor", mock.Anything, mock.Anything).Return(nil)
    err := s.UpdateActor(context.Background(), map[string]any{"id": 1, "name": "name"})
    if err != nil {
        t.Error("UpdateActor returned error")
    }
}

func TestService_DeleteActor(t *testing.T) {

    c := new(mockDb)
    s := NewService(c)
    c.On("DeleteActorById", mock.Anything, mock.Anything).Return(nil)
    err := s.DeleteActor(context.Background(), 1)
    if err != nil {
        t.Error("DeleteActor returned error")
    }
}

func TestService_AddActorToFilm(t *testing.T) {
    c := new(mockDb)
    s := NewService(c)
    c.On("AddActorToFilm", mock.Anything, mock.Anything).Return(nil)
    err := s.AddActorToFilm(context.Background(), 1, 1)
    if err != nil {
        t.Error("AddActorToFilm returned error")
    }
}
