package STmySQL

import (
    "testing"
    "time"
)

type Person struct {
    Id      int       `db:"column=id primarykey=yes table=Users"`
    Name    string    `db:"column=name"`
    Dtadded time.Time `db:"column=dtadded omit=yes"`
    Status  int       `db:"column=status"`
}

func TestInsert(t *testing.T) {
    
    p := Person{
        Id:      0,
        Name:    "Test",
        Dtadded: time.Now(),
        Status:  1,
    }
    
    New("", nil)
    
    sql, err := DB.Insert(p)
    if err != nil {
        t.Error(err)
    }
    
    if sql != "INSERT INTO Users(name,status) VALUES (X'54657374',1);" {
        t.Error("SQL Statement is not correct")
    }
}
