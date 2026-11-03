package login

import (
	"bufio"
	"fmt"
	"html"
	"log"
	"os"
)

var (
	loginName    = make(chan string,1)
    loginPass    = make(chan string,1)
)

type User struct{
	userName,
	secretKey string
}

func (u User)sanitizeUserInput()*User{
	return &User{
		userName: html.EscapeString(u.userName),
		secretKey: html.EscapeString(u.secretKey),
	}
	
}
func getUserInput()*User{
	usr := &User{}
	fmt.Println("\nEnter username:")
	scanner := bufio.NewScanner(os.Stdin)
	//check for any scanner err 
	if err := scanner.Err(); err != nil{
		log.Fatal(err)
	}
	for scanner.Scan() {
		usr.userName = scanner.Text()
		break
	}
	fmt.Println("Enter password:")
	//get password
	for scanner.Scan() {
		usr.secretKey = scanner.Text()
		break
	}
	
	return &User{
		userName: usr.userName,
		secretKey: usr.secretKey,
	}

}

func PostUser()(UserName, Password <- chan string){
	
   sanitizedInput := getUserInput().sanitizeUserInput()
   //transmit sanitized inputs
   loginName  <- sanitizedInput.userName
   loginPass   <- sanitizedInput.secretKey	

   return loginName, loginPass
	
}