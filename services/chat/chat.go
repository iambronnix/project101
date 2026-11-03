package chat

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	chatServer "project101/server/chat"
	"time"
)

func Dial(){
	addr := <- chatServer.Addr
	fmt.Println(addr)
	conn, err := net.DialTimeout("tcp",addr,5*time.Second)
	if err != nil{
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func(){
		io.Copy(os.Stdout,conn)
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	

}
func mustCopy(dst io.Writer, src io.Reader)error{
	if _, err := io.Copy(dst, src);err !=nil{
		return err 		
	}
	return nil
}