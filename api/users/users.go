package users

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"net/http"
)
type User struct{
	UserID int 
	UserName,
	UserHash string
}

func Users(ctx context.Context, handlerAddr string)error{
    log.SetFlags(log.Llongfile)
    apiUrl := "http://" + handlerAddr
	req, err := http.NewRequestWithContext(ctx,"LISTusers",apiUrl,nil)
    if err != nil{
    return err
    }
    resp, err := http.DefaultClient.Do(req)
    if err != nil{
        if !errors.Is(err, context.DeadlineExceeded){
        return err 
        }
        newErr := errors.New("deadline exceeded")
    return newErr
    }    
   defer func(){
     resp.Close= true
   }() 
   
   
  if resp.StatusCode == http.StatusOK {
      r := csv.NewReader(resp.Body)
      r.FieldsPerRecord = -1
      data, err := r.ReadAll()
      log.Println(data)
      if err != nil{
       return err
      }
      
      return nil
  }
  log.Println(resp.Header)
    io.Copy(io.Discard,resp.Body)
 
  Nerr := errors.New(resp.Status) 
  return Nerr  
     
}

