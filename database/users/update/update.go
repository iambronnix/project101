package update

import(
	db "project101/database/dbInit"
)


func UpdatetUser(userName, userHash string)error{
    mydb,error := db.Config()
    if error != nil{
    return error
    }
    defer mydb.Close()
	query, prepErr := mydb.Prepare("UPDATE users SET UserHash=$1 WHERE UserName=$2")
	 if prepErr != nil{
			return prepErr
		}

		_, queryErr := query.Exec(userHash,userName)
		if queryErr!= nil{
			return queryErr
		}

		return nil
	
}