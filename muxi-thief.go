package main

import (
	"fmt"
	"log"
	"net/http"
	"io"
)
func main(){
	req,err:=http.NewRequest("GET","http://muxithief.muxixyz.com/api/v1/login",nil)
	 if err!=nil {
	 	log.Fatal(err)
	 }
	 req.Header.Add("code","111")
	 resp, err := client.Do(req)  //???
	 if err!=nil {
	 	log.Fatal(err)
	 }
	 defer resp.Body.Close()
	 body, err := io.ReadAll(resp.Body)
	 if err != nil {
	 	log.Fatal(err)
	 }
      fmt.Println(body)

}