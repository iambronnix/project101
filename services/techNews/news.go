package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)
func main(){
	dial, err := net.Dial("tcp", "localhost:8000")
	if err != nil{
		dialErr := fmt.Sprintf("%v", err)
		os.WriteFile("/home/iambronnix/project101/servers/logs/news.txt",[]byte(dialErr),0666)
		return	
	}
	terminalWrite(dial, 100*time.Microsecond, time.Now().Add(5*time.Second))
}
func terminalWrite(conn net.Conn, delay time.Duration, readTimeout time.Time){
	buff := make([]byte,uint32(1<<30))
	readErr := conn.SetReadDeadline(readTimeout)
	if readErr != nil{
		
	}
	_, err := conn.Read(buff)
	if err != nil{	
	log.Fatal(err)	
	}
	for _, i := range buff{
		fmt.Print(string(i))
		time.Sleep(delay)
	}
}