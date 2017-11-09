package server

import (
	"database/sql"
	"log"

	"github.com/bobheadxi/calories/config"
)

// Server : Contains the app's database and offers an
// interface to interact with it.
type Server struct {
	db *sql.DB
}

// New : Instantiasafsldfalksdjf
func New(cfg *config.EnvConfig) *Server {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	return &Server{
		db: db,
	}
}

// InsertDataExample : Example use of database insertion
// DEPRECATE ASAP
func (s *Server) InsertDataExample(id string, content string) (int, error) {
	var userid int
	err := s.db.QueryRow(`INSERT INTO test(id, message) VALUES(` + id + `, '` + content + `') RETURNING id`).Scan(&userid)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return userid, nil
}
