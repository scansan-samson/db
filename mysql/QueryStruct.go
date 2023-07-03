package mysql

import (
    "reflect"
    
    l "golang.org/x/exp/slog"
)

func QueryStruct[T any](sql string, parameters ...any) ([]T, error) {
    
    // First of all, get all the database records, ising the old Record/Field method.
    allRecords, err := DB.Query(sql, parameters...)
    if err != nil {
        return make([]T, 0), err
    }
    
    results := make([]T, 0)
    
    for i, record := range allRecords {
        var newStructRecord T
        
        for k, v := range record {
            // Use Reflection to set the value.
            
            structFieldName, structFieldType := getStructDetails[T](k)
            
            // l.INFO("index:%d Key:%s Value:%v structFieldName:%v structFieldType:%v", i, k, "", structFieldName, structFieldType)
            
            switch structFieldType {
            case "int", "int32", "int64":
                // l.INFO("Setting Int64 field: %s to %v type: %T", structFieldName, v.Value, v.Value)
                reflect.ValueOf(&newStructRecord).Elem().FieldByName(structFieldName).SetInt(v.AsInt64())
            
            case "float32", "float64":
                // l.INFO("Setting flaot64 field: %s to %v", structFieldName, v.Value)
                reflect.ValueOf(&newStructRecord).Elem().FieldByName(structFieldName).SetFloat(v.AsFloat())
            
            case "string":
                // l.INFO("Setting String field: %s to %v", structFieldName, v.Value)
                reflect.ValueOf(&newStructRecord).Elem().FieldByName(structFieldName).SetString(v.AsString())
            
            case "Time":
                // l.INFO("Setting Time field: %s to %v", structFieldName, v.Value)
                reflect.ValueOf(&newStructRecord).Elem().FieldByName(structFieldName).Set(reflect.ValueOf(v.AsDate("")))
                
                // Add Blob Support.
            case "[]uint8":
                reflect.ValueOf(&newStructRecord).Elem().FieldByName(structFieldName).Set(reflect.ValueOf(v.AsByte()))
                // l.INFO("Setting Blob field: %s to %v", structFieldName, v.Value)
            
            default:
                l.Error("Database column '%s' (index %d) was not found Type:'%T' Struct Field:'%s' StructValue:'%s'", k, i, structFieldType, structFieldName, structFieldType)
            }
        }
        
        results = append(results, newStructRecord)
    }
    return results, nil
}

func QuerySingleStruct[T any](sql string, parameters ...any) (T, error) {
    
    var SingleResult T
    
    results, err := QueryStruct[T](sql, parameters...)
    if err != nil {
        return SingleResult, err
    }
    if len(results) == 0 {
        return SingleResult, nil
    }
    return results[0], nil
}
