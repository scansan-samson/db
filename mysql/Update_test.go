package STmySQL

import (
    "testing"
    "time"
)

type UpdatePerson struct {
    Id      int       `db:"column=id primarykey=yes table=Users"`
    Name    string    `db:"column=name"`
    Dtadded time.Time `db:"column=dtadded omit=yes"`
    Status  int       `db:"column=status"`
}

func TestUpdate(t *testing.T) {
    
    p := UpdatePerson{
        Id:      12,
        Name:    "Test",
        Dtadded: time.Now(),
        Status:  1,
    }
    
    New("", nil)
    
    sql, err := DB.Update(p)
    if err != nil {
        t.Error(err)
    }
    
    // UPDATE Users SET name = X'54657374',status=1 WHERE id=12;
    // UPDATE Users SET name=X'54657374',status=1 WHERE id=12;
    
    if sql != "UPDATE Users SET name=X'54657374',status=1 WHERE id=12;" {
        t.Errorf("SQL Statement is not correct SQL: %s", sql)
    }
}
