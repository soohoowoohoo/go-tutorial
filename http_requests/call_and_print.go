package main

import (
  "fmt"
  "net/http"
  "time"
  "io/ioutil"
)

func main() {
	netClient := &http.Client{
		Timeout: time.Second * 10,
	}
	  
	response, _ := netClient.Get("https://jsonplaceholder.typicode.com/users")
	buf, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(buf));
}