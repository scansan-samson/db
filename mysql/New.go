package mysql

import (
	"database/sql"
	"errors"
	"sync"
	"time"

	"log/slog"
)

type Database struct {
	dbConnection               *sql.DB
	DSN                        string
	Logger                     *slog.Logger
	ShowSQL                    bool
	Timed                      bool
	Lock                       sync.Mutex
	connected                  bool
	MaxDatabaseOpenConnections int
	MaxDatabaseIdleConnections int
	DatabaseIdleTimeout        time.Duration
}

var DB *Database

var dbs sync.Map
func New(newDSN string, L *slog.Logger) {
	newDB(newDSN, newDSN, L)
}

func NewWithName(name, newDSN string, L *slog.Logger) {
	newDB(name, newDSN, L)
}

func newDB(name, newDSN string, L *slog.Logger) {
	DB = &Database{
		connected: false,
		DSN:       newDSN,
		Logger:    L,
		ShowSQL:   false,
	}

	dbs.Store(newDSN, DB)
}

func getConnection(db *Database) (*sql.DB, error) {
	copiedDB := DB
	if db != nil {
		copiedDB = db
	}

	copiedDB.Lock.Lock()
	// check once more - in case a prev goroutine has established a connection
	if copiedDB.connected && copiedDB.dbConnection != nil {
		copiedDB.Lock.Unlock()
		return copiedDB.dbConnection, nil
	}

	if copiedDB.DSN == "" {
		return nil, errors.New("empty database dsn")
	}

	var err error

	// attempt 3 times to connect, then give up
	for i := 0; i < 3; i++ {
		copiedDB.dbConnection, err = sql.Open("mysql", copiedDB.DSN)

		if err == nil {
			// Open may just validate its arguments without creating a connection to the database.
			// To verify that the data source name is valid, call Ping.
			err = copiedDB.dbConnection.Ping()
			if err == nil {
				break // connection was fine
			}
			copiedDB.Logger.With("attempt", i).With("error", err.Error()).Error("Unable to Ping Database")
			time.Sleep(500 * time.Millisecond) // wait a short while before trying again
			continue
		}
		time.Sleep(500 * time.Millisecond)
		copiedDB.Logger.With("attempt", i).With("error", err.Error()).Error("Unable to Ping Database")
	}

	if err != nil {
		return nil, err
	}

	copiedDB.dbConnection.SetMaxOpenConns(25)
	copiedDB.dbConnection.SetMaxIdleConns(25)
	copiedDB.dbConnection.SetConnMaxIdleTime(5 * time.Minute)
	copiedDB.connected = true
	copiedDB.Lock.Unlock()

	return copiedDB.dbConnection, nil
}
