package main

import (
	//"encoding/json"
   "io/ioutil"
   "log"
   "net/http"
)
//? this function go get the API with is URL and drop it into the terminal
func getApi() string {
   resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
   if err != nil {
      log.Fatalln(err)
   }
//*We Read the response body on the line below.
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
//*Convert the body to type string and print it into the terminal
   sb := string(body)
    return sb;

    //? a example of decoding.json

    /*var dat map[string]interface{}

    if err := json.Unmarshal(body, &dat); err != nil {
      panic(err)
  }
  fmt.Println(dat)

  id := dat["id"].(float64)
    fmt.Println(id)*/
    }

//*encoding the brute JSON  to update the format
