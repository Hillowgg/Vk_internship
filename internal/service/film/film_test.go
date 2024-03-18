package film

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

func (m *mockDb) AddFilm(ctx context.Context, film *database.AddFilmParams) (int32, error) {
    args := m.Called(ctx, film)
    return int32(args.Int(0)), args.Error(1)
}

func (m *mockDb) GetFilmById(ctx context.Context, id int32) (*database.Film, error) {
    args := m.Called(ctx, id)
    return args.Get(0).(*database.Film), args.Error(1)
}

func (m *mockDb) GetFilms() ([]*Film, error) {
    args := m.Called()
    return args.Get(0).([]*Film), args.Error(1)
}

func (m *mockDb) DeleteFilmById(ctx context.Context, id int32) error {
    args := m.Called(id)
    return args.Error(0)
}

func (m *mockDb) UpdateFilm(ctx context.Context, film map[string]any) error {
    args := m.Called(ctx, film)
    return args.Error(0)
}

func (m *mockDb) GetFilmByTitle(title string) (*Film, error) {
    args := m.Called(title)
    return args.Get(0).(*Film), args.Error(1)
}

func (m *mockDb) GetFilmsByUserId(userId int) ([]*Film, error) {
    args := m.Called(userId)
    return args.Get(0).([]*Film), args.Error(1)
}
func (m *mockDb) AddFilmWithActors(ctx context.Context, film *database.AddFilmParams, ids []int32) error {
    args := m.Called(ctx, film, ids)
    return args.Error(0)
}

func (m *mockDb) GetFilmsASCReleaseDate(ctx context.Context) ([]*database.Film, error) {
    args := m.Called(ctx)
    return args.Get(0).([]*database.Film), args.Error(1)
}

func (m *mockDb) GetFilmsDESCReleaseDate(ctx context.Context) ([]*database.Film, error) {
    args := m.Called(ctx)
    return args.Get(0).([]*database.Film), args.Error(1)
}

func (m *mockDb) GetFilmsASCRating(ctx context.Context) ([]*database.Film, error) {
    args := m.Called(ctx)
    return args.Get(0).([]*database.Film), args.Error(1)
}

func (m *mockDb) GetFilmsDESCRating(ctx context.Context) ([]*database.Film, error) {
    args := m.Called(ctx)
    return args.Get(0).([]*database.Film), args.Error(1)
}

func (m *mockDb) GetFilmsASCTitle(ctx context.Context) ([]*database.Film, error) {
    args := m.Called(ctx)
    return args.Get(0).([]*database.Film), args.Error(1)
}

func (m *mockDb) GetFilmsDESCTitle(ctx context.Context) ([]*database.Film, error) {
    args := m.Called(ctx)
    return args.Get(0).([]*database.Film), args.Error(1)
}

func (m *mockDb) SearchFilmByTitleAndActor(ctx context.Context, arg *database.SearchFilmByTitleAndActorParams) (*database.Film, error) {
    args := m.Called(ctx, arg)
    return args.Get(0).(*database.Film), args.Error(1)
}

func (m *mockDb) SearchFilmsByTitle(ctx context.Context, title pgtype.Text) ([]*database.Film, error) {
    args := m.Called(ctx, title)
    return args.Get(0).([]*database.Film), args.Error(1)
}

var _ database.FilmQuerierWithTx = (*mockDb)(nil)

func TestService_CreateFilm(t *testing.T) {
    // write test for CreateFilm
    db := new(mockDb)
    s := NewService(db)
    db.On("AddFilm", mock.Anything, mock.Anything).Return(1, nil)
    _, err := s.CreateFilm(context.Background(), &NewFilm{Title: "title", Description: "description", ReleaseDate: time.Now(), Rating: 1})
    if err != nil {
        t.Error("CreateFilm returned error")
    }
}

func TestService_CreateFilmWithActors(t *testing.T) {
    // write test for CreateFilmWithActors
    db := new(mockDb)
    s := NewService(db)
    db.On("AddFilmWithActors", mock.Anything, mock.Anything, mock.Anything).Return(nil)
    err := s.CreateFilmWithActors(context.Background(), &NewFilm{Title: "title", Description: "description", ReleaseDate: time.Now(), Rating: 1}, []int32{1, 2})
    if err != nil {
        t.Error("CreateFilmWithActors returned error")
    }
}

func TestService_DeleteFilm(t *testing.T) {
    // write test for DeleteFilm
    db := new(mockDb)
    s := NewService(db)
    db.On("DeleteFilmById", mock.Anything, mock.Anything).Return(nil)
    err := s.DeleteFilm(context.Background(), 1)
    if err != nil {
        t.Error("DeleteFilm returned error")
    }
}

func TestService_GetFilm(t *testing.T) {
    // write test for GetFilm
    db := new(mockDb)
    s := NewService(db)
    db.On("GetFilmById", mock.Anything, mock.Anything).Return(&database.Film{}, nil)
    _, err := s.GetFilm(context.Background(), 1)
    if err != nil {
        t.Error("GetFilm returned error")
    }
}

func TestService_GetFilms(t *testing.T) {
    // write test for GetFilms
    db := new(mockDb)
    s := NewService(db)
    db.On("GetFilms").Return([]*Film{}, nil)
    _, err := s.GetFilms(context.Background(), "title", "asc")
    if err != nil {
        t.Error("GetFilms returned error")
    }
}

func TestService_SearchFilmByActor(t *testing.T) {
    // write test for SearchFilmByActor
    db := new(mockDb)
    s := NewService(db)
    db.On("SearchFilmByTitleAndActor", mock.Anything, mock.Anything).Return(&database.Film{}, nil)
    _, err := s.SearchFilmByActor(context.Background(), "title", "actor")
    if err != nil {
        t.Error("SearchFilmByActor returned error")
    }
}

func TestService_SearchFilms(t *testing.T) {
    // write test for SearchFilms
    db := new(mockDb)
    s := NewService(db)
    db.On("SearchFilmsByTitle", mock.Anything, mock.Anything).Return([]*Film{}, nil)
    _, err := s.SearchFilms(context.Background(), "title")
    if err != nil {
        t.Error("SearchFilms returned error")
    }
}

func TestService_UpdateFilm(t *testing.T) {
    // write test for UpdateFilm
    db := new(mockDb)
    s := NewService(db)
    db.On("UpdateFilm", mock.Anything, mock.Anything).Return(nil)
    err := s.UpdateFilm(context.Background(), map[string]any{"id": 1})
    if err != nil {
        t.Error("UpdateFilm returned error")
    }
}
