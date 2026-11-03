package insert

import(
	db "project101/database/dbInit"
)


func InsertUser(userName, userHash string)error{
    mydb,error := db.Config()
    if error != nil{
    return error
    }
    defer mydb.Close()
	query, prepErr := mydb.Prepare("INSERT INTO USERS(username, userhash) VALUES($1,$2) ")
	 if prepErr != nil{
			return prepErr
		}
    
		_, queryErr := query.Exec(userName, userHash)
		if queryErr!= nil{
			return queryErr
		}

		return nil
	
}