package mysql

import (
	"testing"
	"time"
)

type UpdatePerson[StatusType uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64 | string] struct {
	Id      int        `db:"column=id primarykey=yes table=Users"`
	Name    string     `db:"column=name"`
	Dtadded time.Time  `db:"column=dtadded omit=yes"`
	Status  StatusType `db:"column=status"`
}

func generateUpdatePerson[StatusType uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64 | string](value StatusType) UpdatePerson[StatusType] {
	return UpdatePerson[StatusType]{
		0, "Test", time.Now(), value,
	}
}

func testUpdateNumericalErrorValueHelper(t *testing.T, sql string, err error) {
	if err != nil {
		t.Error(err)
	}

	if sql != "UPDATE Users SET name=X'54657374',status=1 WHERE id=12;" {
		t.Errorf("SQL Statement is not correct, found: %s", sql)
	}
}

func testUpdateStringErrorValueHelper(t *testing.T, sql string, err error) {
	if err != nil {
		t.Error(err)
	}

	if sql != `UPDATE Users SET name=X'54657374',status=X='31 WHERE id=12;` {
		t.Errorf("SQL Statement is not correct, found: %s", sql)
	}
}

func TestUpdate(t *testing.T) {
	New("", nil)

	sql, err := DB.Update(generateUpdatePerson(uint8(1)))
	testUpdateNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Update(generateUpdatePerson(uint16(1)))
	testUpdateNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Update(generateUpdatePerson(uint32(1)))
	testUpdateNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Update(generateUpdatePerson(uint64(1)))
	testUpdateNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Update(generateUpdatePerson(int(1)))
	testUpdateNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Update(generateUpdatePerson(int8(1)))
	testUpdateNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Update(generateUpdatePerson(int16(1)))
	testUpdateNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Update(generateUpdatePerson(int32(1)))
	testUpdateNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Update(generateUpdatePerson(int64(1)))
	testUpdateNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Update(generateUpdatePerson(float32(1)))
	testUpdateNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Update(generateUpdatePerson(float64(1)))
	testUpdateNumericalErrorValueHelper(t, sql, err)

	sql, err = DB.Update(generateUpdatePerson("1"))
	testUpdateStringErrorValueHelper(t, sql, err)
}
