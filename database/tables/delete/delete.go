package delete 

import (
	db "project101/database/dbInit"
)

func Delete()error{
	mydb, err := db.Config()
	if err != nil{
		return err
	}
	defer mydb.Close()
	statement := `
	DELETE FROM users`

	_, err = mydb.Exec(statement)
	if err !=nil{
		return err
	}
	return nil
	
}
