package mysql

import (
	"testing"
	"time"
)

type InsertPerson[StatusType uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64 | string] struct {
	Id      int        `db:"column=id primarykey=yes table=Users"`
	Name    string     `db:"column=name"`
	Dtadded time.Time  `db:"column=dtadded omit=yes"`
	Status  StatusType `db:"column=status"`
}

func generateInsertPerson[StatusType uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64 | string](value StatusType) InsertPerson[StatusType] {
	return InsertPerson[StatusType]{
		0, "Test", time.Now(), value,
	}
}

func testInsertNumericalErrorValueHelper(t *testing.T, sql string, err error) {
	if err != nil {
		t.Error(err)
	}

	if sql != "INSERT INTO Users(name,status) VALUES (X'54657374',1);" {
		t.Errorf("SQL Statement is not correct, found: %s", sql)
	}
}

func testInsertStringErrorValueHelper(t *testing.T, sql string, err error) {
	if err != nil {
		t.Error(err)
	}

	if sql != `INSERT INTO Users(name,status) VALUES (X'54657374',X'31');` {
		t.Errorf("SQL Statement is not correct, found: %s", sql)
	}
}

func TestInsert(t *testing.T) {
	New("", nil)

	sql, err := DB.Insert(generateInsertPerson(uint8(1)))
	testInsertNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generateInsertPerson(uint16(1)))
	testInsertNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generateInsertPerson(uint32(1)))
	testInsertNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generateInsertPerson(uint64(1)))
	testInsertNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generateInsertPerson(int(1)))
	testInsertNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generateInsertPerson(int8(1)))
	testInsertNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generateInsertPerson(int16(1)))
	testInsertNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generateInsertPerson(int32(1)))
	testInsertNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generateInsertPerson(int64(1)))
	testInsertNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generateInsertPerson(float32(1)))
	testInsertNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generateInsertPerson(float64(1)))
	testInsertNumericalErrorValueHelper(t, sql, err)

	sql, err = DB.Insert(generateInsertPerson("1"))
	testInsertStringErrorValueHelper(t, sql, err)
}
