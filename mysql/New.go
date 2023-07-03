package mysql

import (
    "database/sql"
    "errors"
    "sync"
    "time"
    
    "golang.org/x/exp/slog"
)

type Database struct {
    dbConnection               *sql.DB
    DSN                        string
    Logger                     *slog.Logger
    Timed                      bool
    Lock                       sync.Mutex
    connected                  bool
    MaxDatabaseOpenConnections int
    MaxDatabaseIdleConnections int
    DatabaseIdleTimeout        time.Duration
}

var DB *Database

func New(DSN string, L slog.Logger) (*Database, error) {
    
    newDatabaseConnection := &Database{
        connected: false,
    }
    
    if DSN == "" {
        return newDatabaseConnection, errors.New("empty database dsn")
    }
    
    var err error
    
    // attempt 3 times to connect, then give up
    for i := 0; i < 3; i++ {
        newDatabaseConnection.dbConnection, err = sql.Open("mysql", DSN)
        
        if err == nil {
            // Open may just validate its arguments without creating a connection to the database.
            // To verify that the data source name is valid, call Ping.
            err = newDatabaseConnection.dbConnection.Ping()
            if err == nil {
                break // connection was fine
            }
            L.With("attempt", i).With("error", err.Error()).Error("Unable to Ping Database")
            time.Sleep(500 * time.Millisecond) // wait a short while before trying again
            continue
        }
        time.Sleep(500 * time.Millisecond)
        L.With("attempt", i).With("error", err.Error()).Error("Unable to Ping Database")
    }
    
    if err != nil {
        return nil, err
    }
    
    newDatabaseConnection.Logger = &L
    newDatabaseConnection.dbConnection.SetMaxOpenConns(25)
    newDatabaseConnection.dbConnection.SetMaxIdleConns(25)
    newDatabaseConnection.dbConnection.SetConnMaxIdleTime(5 * time.Minute)
    newDatabaseConnection.connected = true
    newDatabaseConnection.Lock.Unlock()
    
    return newDatabaseConnection, nil
}

func getConnection() (*sql.DB, error) {
    
    DB.Lock.Lock()
    // check once more - in case a prev goroutine has established a connection
    if DB.connected && DB.dbConnection != nil {
        DB.Lock.Unlock()
        return DB.dbConnection, nil
    }
    return nil, errors.New("no database connection")
}
