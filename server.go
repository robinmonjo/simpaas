package main

import(
  "net/http"
  "fmt"
  "os"
  "encoding/json"
  "os/exec"
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

  //extracting repo URL
  payloadData := []byte(req.FormValue("payload"))
  var payload interface{}
  err := json.Unmarshal(payloadData, &payload)
  if err != nil {
    fmt.Println("[ERROR] parsing body : ", err)
    return
  }

  mapPayload := payload.(map[string]interface{})
  repository := mapPayload["repository"].(map[string]interface{})
  repoURL := repository["url"].(string)
  appName := repository["name"].(string)

  deploy(repoURL, appName)
}

func deploy(repo string, appName string) {
  fmt.Println("Starting deployment for app ", appName, " on: ", repo)
  git, err := exec.LookPath("git")
  if err != nil {
    fmt.Println("git not installed")
  }
  dir, _ := os.Getwd() 
  _, err = exec.Command(git, "clone", repo, dir + "/clones/" + appName).Output()
  if err != nil {
    fmt.Println("[ERROR] git clone failed: ", err)
  }

}