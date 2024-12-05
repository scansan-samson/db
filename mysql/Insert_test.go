package mysql

import (
	"testing"
	"time"
)

type Person[StatusType uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64 | string] struct {
	Id      int        `db:"column=id primarykey=yes table=Users"`
	Name    string     `db:"column=name"`
	Dtadded time.Time  `db:"column=dtadded omit=yes"`
	Status  StatusType `db:"column=status"`
}

func generatePerson[StatusType uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64 | string](value StatusType) Person[StatusType] {
	return Person[StatusType]{
		0, "Test", time.Now(), value,
	}
}

func testNumericalErrorValueHelper(t *testing.T, sql string, err error) {
	if err != nil {
		t.Error(err)
	}

	if sql != "INSERT INTO Users(name,status) VALUES (X'54657374',1);" {
		t.Errorf("SQL Statement is not correct, found: %s", sql)
	}
}

func testStringErrorValueHelper(t *testing.T, sql string, err error) {
	if err != nil {
		t.Error(err)
	}

	if sql != `INSERT INTO Users(name,status) VALUES (X'54657374',X'31');` {
		t.Errorf("SQL Statement is not correct, found: %s", sql)
	}
}

func TestInsert(t *testing.T) {
	New("", nil)

	sql, err := DB.Insert(generatePerson(uint8(1)))
	testNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generatePerson(uint16(1)))
	testNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generatePerson(uint32(1)))
	testNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generatePerson(uint64(1)))
	testNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generatePerson(int(1)))
	testNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generatePerson(int8(1)))
	testNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generatePerson(int16(1)))
	testNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generatePerson(int32(1)))
	testNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generatePerson(int64(1)))
	testNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generatePerson(float32(1)))
	testNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generatePerson(float64(1)))
	testNumericalErrorValueHelper(t, sql, err)

	sql, err = DB.Insert(generatePerson("1"))
	testStringErrorValueHelper(t, sql, err)
}
