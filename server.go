package main

import(
  "net/http"
  "fmt"
  "os"
)

func main() {
  http.HandleFunc("/deploy", deploy)
  port := os.Getenv("PORT")
  if len(port) == 0 {
    port = "9999"
  }
  fmt.Println("simpaas listening on ", port)
}

//routes
func deploy(res http.ResponseWriter, req *http.Request) {
  fmt.Println("Getting a deploy action")
}