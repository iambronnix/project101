package loginapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	crt "project101/api/createAcc"
	l "project101/services/login"
	"github.com/briandowns/spinner"
	"strings"
	"sync"
	"time"

	"github.com/gookit/color"
)
type User struct{
	UserName string   
	PassWord string
}

type Response struct{
	UserId int 
	UserName string 
}

var (
	 deadLineErr = errors.New("deadline exceeded")
)

func GetUser()*User{
	userName, passWord := l.PostUser()
	return &User{
		UserName: <-userName,
		PassWord: <-passWord,
	}
}


func (u *User)LoginApi(ctx context.Context,wg *sync.WaitGroup){//logs in an existing user into the system
     apiUrl := "http://localhost:8443" 
     log.SetFlags(log.Llongfile)
     defer wg.Done()	
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(u)
	if err != nil{
		Nerr := color.Style{color.BgRed,color.FgBlack.Darken(),color.OpBold}.Sprintf("%v",err)
		log.Println(Nerr)
		return
	}
	
	
   req, err := http.NewRequestWithContext(ctx,"CHECKuser",apiUrl,buf)
   if err != nil{
        Nerr := color.Style{color.BgRed,color.FgBlack.Darken(),color.OpBold}.Sprintf("%v",err)
		log.Println(Nerr)
	    return 
	}

     resp,err := http.DefaultClient.Do(req)     
     if err != nil{
            if !errors.Is(err,context.DeadlineExceeded){
            Nerr := color.Style{color.BgRed,color.FgBlack.Darken(),color.OpBold}.Sprintf("%v",err)
		    log.Println(Nerr)
                return 
            }
        Nerr := color.Style{color.BgRed,color.FgBlack.Darken(),color.OpBold}.Sprintf("%v",deadLineErr)
		log.Println(Nerr)
        return
     }

     
    if resp.StatusCode == http.StatusOK{
    data,_ := io.ReadAll(resp.Body)
    fmt.Print(string(data))
    return
     }else{
     
      data, _ := io.ReadAll(resp.Body)
       b := strings.Contains(string(data), "sql: no rows in result set")
       
        if b{
           s := spinner.New(spinner.CharSets[35],250*time.Millisecond)
           s.Color("green")
           s.Prefix = color.Style{color.BgBlue,color.FgBlue,color.OpConcealed}.Sprint("Let's Create an Account for you\n")   	
           s.Start()
                 time.Sleep(5 *time.Second)
           s.Stop() 
           //create an account for a new user
              crt.GetUser().CreateApi(ctx,wg)
              fmt.Println(color.Style{color.BgBlue,color.FgCyan,color.OpBold}.Sprintf("Account Created Successfully : Login again"))            
              return
              
        }
        Nerr := color.Style{color.BgRed,color.FgBlack.Darken(),color.OpBold}.Sprintf("%v:%s",resp.StatusCode,string(data))
        log.Println(Nerr)      
     }

   
}
     
