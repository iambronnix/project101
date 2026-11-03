package queryusers

import (
	"log"
	db "project101/database/dbInit"
)

type User struct{
	UserID int
	Username,
	Userhash string
}
var Users []User

func QueryUsers()(error,[]User){ 
	mydb, error := db.Config()
	if error != nil{
		log.Fatalf("%v",error)
		return error,nil
	}
	defer mydb.Close()
	
	rows, scanErr := mydb.Query("SELECT * FROM users")
	if scanErr != nil{
		return scanErr,nil
	}
	user := &User{}
	
	if err :=rows.Err(); err != nil{
		return err,nil
	}
   
	for rows.Next(){	  
		rows.Scan(&user.UserID,&user.Username,&user.Userhash)
	    Users = append(Users, *user)		        		
	}
 		
    return nil, Users
	
}