package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)
const (
	url1 uint8 = iota + 1
	url2 
	url3
	url4
	url5
)
var(
	wg sync.WaitGroup
	timeoutErr = errors.New("timeOut while fetching news")
    newsUrl = map[uint8]string{
	url1 : "https://hn.algolia.com/api/v1/search?query=tech&tags=story",
	url2 : "https://hn.algolia.com/api/v1/search?tags=front_page",
	url3 : "https://news.google.com/rss/search?q=technology&hl=en-US&gl=US&ceid=US:en",
	url4 : "https://news.google.com/rss/search?q=AI+OR+ChatGPT+OR+Google&hl=en-US&gl=US&ceid=US:en&tbs=qdr:d",
}
)


func main(){
	getNews()
}

func getNews(){
	ctx, cancel := context.WithDeadlineCause(context.Background(),time.Now().Add(10*time.Second),timeoutErr)
	defer cancel()

	urlSend := func(ctx context.Context,response chan http.Response,nUrl string, wg *sync.WaitGroup){

         req, err := http.NewRequestWithContext(ctx,http.MethodGet,nUrl,nil)	
          if err != nil{
               log.Println(err)
              return
          }	

          resp, err := http.DefaultClient.Do(req)
          if err != nil{
               log.Println(err)
              return
          }
          defer func(){
            close(response)
            wg.Done()
          }()
          
         select{
          case <- ctx.Done():
          close(response)
          case response <- *resp:
         }
          
	}

	urlRecv := func(response chan http.Response,id int, wg *sync.WaitGroup){
		  fmt.Print(id)	
           buf := make([]byte, uint32(1<<20))
			resp := <-response
			
			defer func(){
				io.Copy(io.Discard,resp.Body)
				resp.Close = true
				wg.Done()
			}()

			if resp.Body != nil{
			   err := json.NewDecoder(resp.Body).Decode(&buf)
				if err != nil{
					return
				}
				data, _ :=io.ReadAll(resp.Body)
				fmt.Print(string(data))
				//os.WriteFile("/home/iambronnix/project101/server/news/java.txt",buf,0644)
				fmt.Printf("id : %v is done writing response", id)
			}
				     
		
	}
	

	resp := make(chan http.Response)

	for j, k := range newsUrl{
		wg.Add(1)
		go urlSend(ctx,resp,k,&wg)
		fmt.Printf("fetching %v:%v\n",j,k)
	}
	
    for i :=0; i<=8; i++{
    fmt.Print(i)
      wg.Add(1)
      go urlRecv(resp,i, &wg)      
    }
	wg.Wait()
	fmt.Print("done...good job buddy")
}                                                                                                                                                                            