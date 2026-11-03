package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"project101/handler/httpHandler"
	"testing"
	"time"
)
type User struct{
	userName,
	secretKey string
}

func TestHttpHandler(t *testing.T){
	srv := &http.Server{
		Addr: "localhost:8443",
		Handler: http.TimeoutHandler(//requests first pass through the middleware 
			httphandler.DefaultMethodsHandler(), 2 *time.Minute, " "),
		IdleTimeout: 5 *time.Minute,
		ReadHeaderTimeout: time.Minute,
	}

	l, err := net.Listen("tcp", srv.Addr)
	if err != nil{
		t.Fatal(err)
	}

	go func() {
		err := srv.Serve(l)
		if err != http.ErrServerClosed{
			t.Error(err)
		}
	}()

	
	user := &User{}
	user.userName = "erick"
	user.secretKey = "iamerick"
	buf := new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(&user)	
	if err != nil{
		t.Errorf("%v", err)
	}
	fmt.Print(user)
//test parameters are awfull
	testCases := []struct{
		method string
		body io.Reader
		code int
		response string	
	}{
	
		{"CREATEacc",buf,http.StatusCreated,""},
		{"LISTusers",nil,http.StatusOK,""},
		{"CHECKuser",buf, http.StatusOK, ""},
		
		
	}

	client := new(http.Client)
	path := fmt.Sprintf("http://%s/",srv.Addr)

	for i, c := range testCases{
		r, err := http.NewRequest(c.method, path,c.body)
		if err != nil{
			t.Errorf("%d: %v", i, err)
			continue
		}
		resp, err := client.Do(r)
		if err != nil{
			t.Errorf("%d: %v", i, err)
			continue
		}
		
		if resp.StatusCode != c.code{
			t.Errorf("%d: unexpected status code: %q",i,resp.Status)
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil{
			t.Errorf("%d: %v",i,err)
			continue
		}

		_ = resp.Body.Close()

		if c.response != string(b){
			t.Errorf("%d: expected %q; actual %q", i, c.response, b)
		}
		if err := srv.Close();err != nil{
			t.Fatal(err)
		}
	}
}