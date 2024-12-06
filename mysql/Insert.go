package mysql

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

func (db *Database) Insert(dbStructure any) (string, error) {

	t := reflect.TypeOf(dbStructure)
	InsertTable := ""
	endsql := ""
	buildsql := ""
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("db")
		dbStructureMap := decodeTag(tag)

		if reflect.ValueOf(dbStructure).Field(i).CanInterface() {
			value := reflect.ValueOf(dbStructure).Field(i).Interface()
			// l.INFO("%d. Value='%v'  %v (%v), tag: '%v'\n", i+1, value, field.Name, field.Type.Name(), tag)

			if dbStructureMap["column"] == "" {
				return "", errors.New("no column name specified for field" + field.Type.Name())
			}

			if dbStructureMap["primarykey"] == "yes" {
				// l.INFO("Primary Key Found: %s", dbStructureMap["table"])
				InsertTable = dbStructureMap["table"]
			}

			if dbStructureMap["omit"] != "yes" && dbStructureMap["primarykey"] != "yes" {
				buildsql = buildsql + dbStructureMap["column"] + ","

				switch field.Type.Name() {
				case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int", "int32", "int64":
					endsql = endsql + fmt.Sprintf("%v", value) + ","
				case "string":
					endsql = endsql + hexRepresentation(value.(string)) + ","
				case "float32", "float64":
					endsql = endsql + fmt.Sprintf("%v", value) + ","
				case "Time":
					endsql = endsql + fmt.Sprintf("'%s'", value.(time.Time).Format("2006-01-02 15:04:05")) + ","
				default:
					db.Logger.With("type", field.Type.Name()).With("value", value).Error("type error")
					endsql = endsql + "'" + value.(string) + "',"
				}
			}
		}
	}
	// Get Rid of Trailling Comma
	buildsql = strings.TrimSuffix(buildsql, ",")
	endsql = strings.TrimSuffix(endsql, ",")

	SQL := "INSERT INTO " + InsertTable + "(" + buildsql + ") VALUES (" + endsql + ");"

	return SQL, nil
}
