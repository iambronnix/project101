package main

import (
	"context"
	"log"
	l "project101/api/login"
	"sync"
	
	// clr "github.com/fatih/color"
)
var (
	wg sync.WaitGroup
	serverAddr, chatAddr string
)


func main(){
	log.SetFlags(log.Llongfile)
	ctx, cancel := context.WithCancel(context.Background())
	 defer func(){
		 cancel()
		}()
	wg.Add(1)
				   	
   	go l.GetUser().LoginApi(ctx,&wg)
           
    wg.Wait()	
	  
		       
	
}