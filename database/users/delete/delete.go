package delete

import (
	"database/sql"
	db "project101/database/dbInit"
)


func deleteUser(userHash string)(error, sql.Result){
	mydb, error := db.Config()
	if error != nil{
		return error, nil
	}
	defer mydb.Close()
	statement, err := mydb.Prepare("DELETE  FROM users WHERE userHash = $1")
	if err != nil{
		return err, nil
	}
     result , err := statement.Exec(userHash)
	if err != nil{
		return err, nil
	}
	return nil, result
 
}

