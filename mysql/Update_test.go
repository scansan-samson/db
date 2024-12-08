package mysql

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type UpdatePerson[StatusType uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64 | string] struct {
	Id      int        `db:"column=id primarykey=yes table=Users"`
	Name    string     `db:"column=name"`
	Dtadded time.Time  `db:"column=dtadded omit=yes"`
	Status  StatusType `db:"column=status"`
}

type UpdatePersonTime struct {
	Id      int       `db:"column=id primarykey=yes table=Users"`
	Name    string    `db:"column=name"`
	Dtadded time.Time `db:"column=dtadded"`
}

func generateUpdatePerson[StatusType uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64 | string](value StatusType) UpdatePerson[StatusType] {
	return UpdatePerson[StatusType]{
		0, "Test", time.Now(), value,
	}
}

func generateUpdatePersonTime(id int) UpdatePersonTime {
	return UpdatePersonTime{
		id, "Test", time.Date(2024, time.December, 7, 15, 29, 25, 10, time.UTC),
	}
}

func testUpdateNumericalErrorValueHelper(t *testing.T, sql string, err error) {
	assert.NoError(t, err)
	assert.Equal(t, "UPDATE Users SET name=X'54657374',status=1 WHERE id=0;", sql)
}

func testUpdateStringErrorValueHelper(t *testing.T, sql string, err error) {
	assert.NoError(t, err)
	assert.Equal(t, "UPDATE Users SET name=X'54657374',status=X'31' WHERE id=0;", sql)
}

func testUpdateTimeErrorValueHelper(t *testing.T, sql string, err error) {
	assert.NoError(t, err)
	assert.Equal(t, "UPDATE Users SET name=X'54657374',dtadded='2024-12-07 15:29:25' WHERE id=0;", sql)
}

func TestUpdate(t *testing.T) {
	New("", nil)

	sql, err := DB.Update(generateUpdatePerson(uint(1)))
	testUpdateNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Update(generateUpdatePerson(uint8(1)))
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

	sql, err = DB.Update(generateUpdatePersonTime(0))
	testUpdateTimeErrorValueHelper(t, sql, err)
}
