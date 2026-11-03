package loginapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	lgn "project101/services/login"
	"strconv"
	"sync"
)
const apiUrl = "http://localhost:8443"

type User struct{
	UserName,
	SecretKey string
}
func GetUser()*User{
	username, secretkey := lgn.PostUser()
	return &User{
		UserName: <-username,
		SecretKey: <-secretkey,
		
	}
}

func (usr *User)CreateApi(ctx context.Context,wg *sync.WaitGroup){//creates an account for a new user
//encode user details into a buffer 
   
    defer wg.Done()
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(usr)
	if err != nil{
		fmt.Println(err)
		return 
	}
//prepare a request
   req, err := http.NewRequestWithContext(ctx,"CREATEacc",apiUrl,buf)
   if err != nil{
    fmt.Println(err)
	return 
   	}
//send and receive a response
     resp,err := http.DefaultClient.Do(req)
     //test type of error
     if err != nil{
            if !errors.Is(err,context.DeadlineExceeded){
             fmt.Println(err)
                return 
            }
            fmt.Println(context.DeadlineExceeded)
       return 
     }
     defer func(){
     //drain the contents of the underlying tcp connection
     _, err = io.Copy(io.Discard,resp.Body)
     if err != nil{
     fmt.Println(err)
     return 
     }
     
    resp.Close = true//explicitly close the underlying tcp connection
     }()
     
    if resp.StatusCode == http.StatusCreated{
       return 
    }
    Nerr := errors.New(strconv.Itoa(resp.StatusCode))
    fmt.Println(Nerr)	
	
}
