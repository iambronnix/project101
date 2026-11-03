package deleteusers

import (
	"fmt"
	db "project101/database/dbInit"
)

func DropTable()(error, string){
	mydb, error := db.Config()
      if error != nil{
      return error, fmt.Sprintln()
      }	
      defer mydb.Close()

      rs, dropErr := mydb.Exec("DROP TABLE users")
      if dropErr != nil{
       return dropErr, fmt.Sprintln()
      }
      return nil, fmt.Sprintln(rs)
}

