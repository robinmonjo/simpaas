package main

import(
  "net/http"
  "fmt"
  "os"
  "encoding/json"
)

func main() {
  http.HandleFunc("/deploy", startDeployment)
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
func startDeployment(res http.ResponseWriter, req *http.Request) {
  
  if req.Method != "POST" {
    fmt.Fprintf(res, "Up, waiting for a POST request from github.com")
    return
  }

  fmt.Println("Processing deployment")

  payloadData := []byte(req.FormValue("payload"))
  var payload interface{}
  err := json.Unmarshal(payloadData, &payload)
  if err != nil {
    fmt.Println("[ERROR] parsing body : ", err)
    return
  }

  deploy(payload)
}

func deploy(payload interface{}) {
  mapPayload := payload.(map[string]interface{})
  repository := mapPayload["repository"].(map[string]interface{})
  repo_url := repository["url"]
  fmt.Println("Getting ready to deploy app on: ", repo_url)
}