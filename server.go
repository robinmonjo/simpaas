package main

import(
  "net/http"
  "fmt"
  "os"
  "encoding/json"
  "os/exec"
  "io/ioutil"
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
  cloneApp(repo, appName)
  buildApp(appName)
}

func cloneApp(repo string, appName string) {
  fmt.Print("Downloading ", appName, " on: ", repo, " ...")
  //1: removing old version
  dir, _ := os.Getwd()
  err := exec.Command("rm", "-rf", dir + "/clones/" + appName).Run()
  if err != nil {
    fmt.Println("failed to erase last app version")
  }
  git, err := exec.LookPath("git")
  if err != nil {
    fmt.Println("git not installed")
  }
  err = exec.Command(git, "clone", repo, dir + "/clones/" + appName).Run()
  if err != nil {
    fmt.Println("[ERROR] git clone failed: ", err)
  }
  fmt.Println(" done")
}

func buildApp(appName string) {
  dir, _ := os.Getwd()
  config, err := ioutil.ReadFile(dir + "/clones/" + appName + "/simpaas_config.json")
  if err != nil {
    fmt.Println("[ERROR] failed to read simpaas_config: ", err)
  }
  var conf interface{}
  err = json.Unmarshal(config, &conf)
  if err != nil {
    fmt.Println("[ERROR] failed to parse simpaas_config: ", err)
  }
  mapConf := conf.(map[string]interface{})
  buildpackRepo := mapConf["buildpack_repo"].(string)
  fmt.Println("Will use ", buildpackRepo, " to build app")
}
