package mysql

import (
	"testing"
	"time"
)

type InsertPerson[StatusType uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64 | string] struct {
	Id      int        `db:"column=id primarykey=yes table=Users"`
	Name    string     `db:"column=name"`
	Dtadded time.Time  `db:"column=dtadded omit=yes"`
	Status  StatusType `db:"column=status"`
}

type InsertPersonTime struct {
	Id      int       `db:"column=id primarykey=yes table=Users"`
	Name    string    `db:"column=name"`
	Dtadded time.Time `db:"column=dtadded"`
}

func generateInsertPerson[StatusType uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64 | string](value StatusType) InsertPerson[StatusType] {
	return InsertPerson[StatusType]{
		0, "Test", time.Now(), value,
	}
}

func generateInsertPersonTime(id int) InsertPersonTime {
	return InsertPersonTime{id, "Test", time.Date(2024, time.December, 7, 15, 29, 25, 10, time.UTC)}
}

func generateArrayInsertPerson[StatusType uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64 | string](value StatusType) []InsertPerson[StatusType] {
	return []InsertPerson[StatusType]{
		{0, "Test", time.Now(), value},
		{1, "Test", time.Now(), value},
		{2, "Test", time.Now(), value},
		{3, "Test", time.Now(), value},
		{4, "Test", time.Now(), value},
	}
}

func generateArrayInsertPersonTime() []InsertPersonTime {
	output := make([]InsertPersonTime, 5)
	for i := 0; i < 5; i++ {
		output[i] = generateInsertPersonTime(i)
	}
	return output
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

func testInsertTimeErrorValueHelper(t *testing.T, sql string, err error) {
	if err != nil {
		t.Error(err)
	}
	if sql != `INSERT INTO Users(name,dtadded) VALUES (X'54657374','2024-12-07 15:29:25');` {
		t.Errorf("SQL Statement is not correct, found: %s", sql)
	}
}

func testInsertManyNumericalErrorValueHelper(t *testing.T, sql string, err error) {
	if err != nil {
		t.Error(err)
	}

	if sql != `INSERT INTO Users(name,status) VALUES (X'54657374',1)
(X'54657374',1)
(X'54657374',1)
(X'54657374',1)
(X'54657374',1);` {
		t.Errorf("SQL Statement is not correct, found: %s", sql)
	}
}

func testInsertManyStringErrorValueHelper(t *testing.T, sql string, err error) {
	if err != nil {
		t.Error(err)
	}

	if sql != `INSERT INTO Users(name,status) VALUES (X'54657374',X'31')
(X'54657374',X'31')
(X'54657374',X'31')
(X'54657374',X'31')
(X'54657374',X'31');` {
		t.Errorf("SQL Statement is not correct, found: %s", sql)
	}
}

func testInsertManyTimeErrorValueHelper(t *testing.T, sql string, err error) {
	if err != nil {
		t.Error(err)
	}

	if sql != `INSERT INTO Users(name,dtadded) VALUES (X'54657374','2024-12-07 15:29:25')
(X'54657374','2024-12-07 15:29:25')
(X'54657374','2024-12-07 15:29:25')
(X'54657374','2024-12-07 15:29:25')
(X'54657374','2024-12-07 15:29:25');` {
		t.Errorf("SQL Statement is not correct, found: %s", sql)
	}
}

func TestInsert(t *testing.T) {
	New("", nil)
	sql, err := DB.Insert(generateInsertPerson(uint(1)))
	testInsertNumericalErrorValueHelper(t, sql, err)
	sql, err = DB.Insert(generateInsertPerson(uint8(1)))
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

	sql, err = DB.Insert(generateInsertPersonTime(0))
	testInsertTimeErrorValueHelper(t, sql, err)
}

func TestInsertMany(t *testing.T) {
	New("", nil)

	sql, err := InsertMany[InsertPerson[uint]](generateArrayInsertPerson(uint(1)))
	testInsertManyNumericalErrorValueHelper(t, sql, err)
	sql, err = InsertMany[InsertPerson[uint8]](generateArrayInsertPerson(uint8(1)))
	testInsertManyNumericalErrorValueHelper(t, sql, err)
	sql, err = InsertMany[InsertPerson[uint16]](generateArrayInsertPerson(uint16(1)))
	testInsertManyNumericalErrorValueHelper(t, sql, err)
	sql, err = InsertMany[InsertPerson[uint32]](generateArrayInsertPerson(uint32(1)))
	testInsertManyNumericalErrorValueHelper(t, sql, err)
	sql, err = InsertMany[InsertPerson[uint64]](generateArrayInsertPerson(uint64(1)))
	testInsertManyNumericalErrorValueHelper(t, sql, err)
	sql, err = InsertMany[InsertPerson[int]](generateArrayInsertPerson(int(1)))
	testInsertManyNumericalErrorValueHelper(t, sql, err)
	sql, err = InsertMany[InsertPerson[int8]](generateArrayInsertPerson(int8(1)))
	testInsertManyNumericalErrorValueHelper(t, sql, err)
	sql, err = InsertMany[InsertPerson[int16]](generateArrayInsertPerson(int16(1)))
	testInsertManyNumericalErrorValueHelper(t, sql, err)
	sql, err = InsertMany[InsertPerson[int32]](generateArrayInsertPerson(int32(1)))
	testInsertManyNumericalErrorValueHelper(t, sql, err)
	sql, err = InsertMany[InsertPerson[int64]](generateArrayInsertPerson(int64(1)))
	testInsertManyNumericalErrorValueHelper(t, sql, err)
	sql, err = InsertMany[InsertPerson[float32]](generateArrayInsertPerson(float32(1)))
	testInsertManyNumericalErrorValueHelper(t, sql, err)
	sql, err = InsertMany[InsertPerson[float64]](generateArrayInsertPerson(float64(1)))
	testInsertManyNumericalErrorValueHelper(t, sql, err)

	sql, err = InsertMany[InsertPerson[string]](generateArrayInsertPerson("1"))
	testInsertManyStringErrorValueHelper(t, sql, err)

	sql, err = InsertMany[InsertPersonTime](generateArrayInsertPersonTime())
	testInsertManyTimeErrorValueHelper(t, sql, err)
}
