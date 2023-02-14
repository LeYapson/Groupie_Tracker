package main

import (
	"encoding/json"
   "io/ioutil"
   "log"
   "net/http"
)
//? this function go get the API with is URL and drop it into the terminal
func getApi() {
   resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
   if err != nil {
      log.Fatalln(err)
   }
//*We Read the response body on the line below.
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
//*Convert the body to type string
   sb := string(body)
   log.Printf(sb)
}