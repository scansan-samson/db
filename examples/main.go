package main

import (
    "fmt"
    "os"
    "time"
    
    "log/slog"
    
    MySQL "github.com/SpaceTent/db/mysql"
)

type UpdatePerson struct {
    Id      int       `db:"column=id primarykey=yes table=Users"`
    Name    string    `db:"column=name"`
    Dtadded time.Time `db:"column=dtadded"`
    Status  int       `db:"column=status"`
    Ignored int       `db:"column=ignored omit=yes"`
}

func main() {
    
    DSN := ""
    textHandler := slog.NewTextHandler(os.Stdout, nil)
    l := slog.New(textHandler)
    
    MySQL.New(DSN, l)
    
    p := UpdatePerson{
        Id:      12,
        Name:    "Test",
        Dtadded: time.Now(),
        Status:  1,
    }
    
    sql, err := MySQL.DB.Update(p)
    if err != nil {
        l.Error(err.Error())
    }
    fmt.Println(sql)
    
}
