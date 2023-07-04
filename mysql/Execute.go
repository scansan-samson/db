package mysql

func (db *Database) Execute(sql string, parameters ...any) (int64, int64, error) {
    
    DatabaseConnection, err := getConnection()
    if err != nil {
        return 0, 0, err
    }
    
    Result, err := DatabaseConnection.Exec(sql, parameters...)
    if err != nil {
        return 0, 0, err
    }
    
    LastInsertedID, _ := Result.LastInsertId()
    RowsAffected, _ := Result.RowsAffected()
    
    if db.ShowSQL {
        db.Logger.With("lastid", LastInsertedID).With("rows effected", RowsAffected).Info(sql)
    }
    
    return LastInsertedID, RowsAffected, nil
}
