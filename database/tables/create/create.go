package createusers

import (
	"fmt"
	db "project101/database/dbInit"
)

func NewTable()(error, string){//changed oids
	mydb, err := db.Config()
	if err != nil {
		return err, fmt.Sprintln()
	}
	defer mydb.Close()
	createStatement := `
	CREATE TABLE users(	
                        ID int GENERATED ALWAYS AS IDENTITY	,
						UserName text NOT NULL,
					    UserHash text NOT NULL,
									PRIMARY KEY(ID),
									UNIQUE(UserHash));
									`


rs, createErr := mydb.Exec(createStatement)
if createErr!=nil{
	return createErr, fmt.Sprintln()
}
return nil,fmt.Sprintln(rs)
}


// 
//  ID integer NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,