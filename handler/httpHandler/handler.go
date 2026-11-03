package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	det "project101/database/tables/details"
	ins "project101/database/users/insert"
	qry "project101/database/users/query"
	cipher "project101/handler/hash"
	"sort"
	"strconv"
	"strings"
	"time"
)

type User struct{
	UserName,
	PassWord string	
} 

var Addr = make(chan string, 1)
type Methods map[string]http.Handler

func(h Methods)ServeHTTP(w http.ResponseWriter, r *http.Request){
	defer func(r io.ReadCloser){
		_,_ = io.Copy(io.Discard, r)
	    _= r.Close()
	}(r.Body)

	if handler, ok := h[r.Method];ok{
		if handler == nil{
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}else{
			handler.ServeHTTP(w,r)
		}
			
		return
	}

   	w.Header().Add("allow", h.allowedMethods())
	//skipped testing if methods are allowed just to have custom method without explicitly handling requests
     //it's better to handle requests explicitly rather than test methods but i dont like it either

     //	if r.Method != http.MethodOptions{
	//	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	//	}
}
//just incase the server starts handling external requests 
// and still introduces an attacking vector bcoz it exposes internal methods incase user sends 
// an invalid method 
func (h Methods) allowedMethods()string{
	a := make([]string, 0, len(h))

	for k := range h {
		a =append(a, k)
	}
	sort.Strings(a)

	return strings.Join(a,", ")

}
func DefaultMethodsHandler() http.Handler{
	return Methods{
		//handles post requests for new users into the system
		"CREATEacc": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
				
				user := &User{}
                err := json.NewDecoder(r.Body).Decode(&user)
                if err != nil{
                http.Error(w, fmt.Sprintln(err), http.StatusBadRequest)
                return 
                }
                cipherText := cipher.HashFunction(user.UserName,user.PassWord)
                err = ins.InsertUser(user.UserName,cipherText)
                if err != nil{
                
                http.Error(w, fmt.Sprintln(err),http.StatusInternalServerError)
                 return 
                }

                w.WriteHeader(http.StatusCreated)
                				
			},
		),

		//handles users inthe database requests
		"LISTusers": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			err, details := det.QueryUsers()   
           if err != nil{
              http.Error(w, fmt.Sprintln(err), http.StatusInternalServerError)
              return
           } 		
           
         wrt := csv.NewWriter(w)
         defer wrt.Flush()
         for _, usr := range details{
             if err := wrt.Write([]string{strconv.Itoa(usr.UserID),usr.Username,usr.Userhash});err != nil{
             http.Error(w, fmt.Sprintln(err), http.StatusInternalServerError)
             return 
             }         
         }
          			
		    },
		),
		//handles a new login into the system by returning user details if exists
		"CHECKuser": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			user := &User{}
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil{
				http.Error(w, fmt.Sprintln(err), http.StatusInternalServerError)
				return
			}
		    cipherText := cipher.HashFunction(user.UserName,user.PassWord) 
						err,details := qry.QueryUser(cipherText)
						if err != nil{
							http.Error(w, fmt.Sprintln(err),http.StatusInternalServerError)
							return
								}
			buf := new(bytes.Buffer)
            if err := json.NewEncoder(buf).Encode(details); err != nil{
               http.Error(w, fmt.Sprintln(err),http.StatusInternalServerError)
               return 
            }	
            
           
           if _,err := io.Copy(w,buf); err != nil{
              http.Error(w, fmt.Sprintln(err),http.StatusInternalServerError)
              return            
           }	
						
		},
	),
		
	}
}

func main(){
	
	srv := &http.Server{
		Addr: "localhost:8443",
		Handler: DefaultMethodsHandler(),
		ReadHeaderTimeout: time.Minute,
	}
	fmt.Println(srv.Addr)

	l, err := net.Listen("tcp", srv.Addr)
	if err != nil{
		log.Println(err)
		return
	}
	Addr <- srv.Addr
	close(Addr)
	for {
		err := srv.Serve(l)
		if err != nil{
			log.Fatal(err)
		}
 }


}