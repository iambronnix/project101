package ui1

import (
	handler "project101/handler/httpHandler"
	chat "project101/server/chat"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gookit/color"
)

var (
	handlerAddr = make(chan string,1)
	chatAddr = make(chan string,1)
	wg =&sync.WaitGroup{}
)


func CoreInit(){
	wg.Add(2)
	go handler.Start(wg)
	s := spinner.New(spinner.CharSets[35],250*time.Millisecond)
    s.Color("green")
    s.Prefix = color.Style{color.BgBlue,color.FgBlue,color.OpConcealed}.Sprint("Initialising handler :")   	
    s.Start()
          time.Sleep(5 *time.Second)
    s.Stop()  
	
    go chat.Start(wg)	
   	n := spinner.New(spinner.CharSets[35],250*time.Millisecond)
       n.Color("green")
       n.Prefix = color.Style{color.BgBlue,color.FgBlue,color.OpConcealed}.Sprint("Initialising chat server :")   	
       n.Start()
            time.Sleep(5 *time.Second)
       n.Stop()  
       wg.Wait()
    	
}
