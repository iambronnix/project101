package queryusers

import (
	"log"
	db "project101/database/dbInit"
)
type User struct{
	UserId int
	UserName string
}

var Users []User

func QueryUser(userHash string)(error,[]User){ 
	log.SetFlags(log.Llongfile)
	mydb, error := db.Config()
	if error != nil{
		log.Fatalf("%v",error)
		return error,nil
	}
	defer mydb.Close()
	user := &User{}
	scanner, prepErr := mydb.Prepare("SELECT UserName,ID FROM users WHERE UserHash=$1")
	if prepErr != nil{
		return prepErr,nil
	}	
 	if scanErr := scanner.QueryRow(userHash).Scan(&user.UserName,&user.UserId); scanErr!=nil{
     return scanErr,nil
  }
  err := scanner.Close()
	if err != nil{
		return err,nil
	}
Users = append(Users, *user)
	
    return nil, Users
	
}