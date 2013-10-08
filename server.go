package main

import(
  "net/http"
  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
)

func main() {
  http.HandleFunc("/deploy", deploy)
  port := os.Getenv("PORT")
  if len(port) == 0 {
    port = "9999"
  }
  fmt.Println("simpaas listening on ", port)
  err := http.ListenAndServe(":" + port, nil)
  if err != nil {
    panic(err)
  }
}

//routes
func deploy(res http.ResponseWriter, req *http.Request) {
  
  if req.Method != "POST" {
    fmt.Fprintf(res, "Up, waiting for a POST request from github.com")
    return
  }

  fmt.Println("Processing deployment")
  //reading body
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    fmt.Println("[ERROR] reading body : ", err)
    return
  }
  fmt.Println("Body: ", string(body))

  //parsing JSON
  var data interface{}
  err = json.Unmarshal(body, &data)
  if err != nil {
    fmt.Println("[ERROR] parsing body : ", err)
    return
  }

  fmt.Println("Data: ", data)




}