package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	config       *Config
	db           *sql.DB
	AdRepository *AdRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {

	db, err := sql.Open("postgres", s.config.ConnectString)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

// Ad ...
func (s *Store) Ad() *AdRepository {
	if s.AdRepository != nil {
		return s.AdRepository
	}

	s.AdRepository = &AdRepository{
		store: s,
	}

	return s.AdRepository
}
