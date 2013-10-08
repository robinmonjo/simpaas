package main

import(
  "net/http"
  "fmt"
)

func main() {
  http.HandleFunc("/deploy", root)
  port := os.Getenv("PORT")
  if len(port) == 0 {
    port = "9999"
  }
  fmt.Println("simpaas listening on ", port)
}

//routes
func deploy(res http.ResponseWriter, req *http.Request) {
  
}