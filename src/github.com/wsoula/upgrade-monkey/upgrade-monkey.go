package main
import (
  "encoding/json"
  "io/ioutil"
  "log"
  "reflect"
  "net/http"
  "os"
  "time"
)
type people struct {
        Name string `json:"name"`
}
func main() {
  // Nomad
  githubLatestRelease("hashicorp/nomad","NOMAD_VERSION")
  // Hashi-UI
  githubLatestRelease("jippi/hashi-ui","HASHIUI_VERSION")
  // Consul
  githubLatestRelease("hashicorp/consul","CONSUL_VERSION")
}
func getUrl(url string) []byte {
  spaceClient := http.Client{
    Timeout: time.Second * 2, // Max of 2 seconds
  }
  req, err := http.NewRequest(http.MethodGet, url, nil)
  if err != nil {
    log.Fatal(err)
  }
  res, getErr := spaceClient.Do(req)
  if getErr != nil {
    log.Fatal(getErr)
  }
  body, readErr := ioutil.ReadAll(res.Body)
  if readErr != nil {
    log.Fatal(readErr)
  }
  return body
}
func githubLatestRelease(orgrepo string, version_env_var string) string {
  var nomadUrl string = "https://api.github.com/repos/"+orgrepo+"/tags"
  var objs interface{}
  json.Unmarshal([]byte(getUrl(nomadUrl)), &objs)
  objArr, ok := objs.([]interface{})
  if !ok {
    log.Fatal("expected an array of objects")
  }
  var tags []string
  for i, obj := range objArr {
    obj, ok := obj.(map[string]interface{})
    if !ok {
      log.Fatalf("expected type map[string]interface{}. got %s", reflect.TypeOf(objArr[i]))
    }
    for key, value := range obj {
      if key == "name" {
        // Below will show list of strings as it came back
        // Design decision: trust above order since sort.Strings(tags) will mark 0.8.4-rc1 newer than 0.8.4
        //println(value.(string))
        tags = append(tags,value.(string))
      }
    }
  }
  var latest string = tags[0]
  var current string = os.Getenv(version_env_var)
  if latest != current {
    println("Upgrade "+version_env_var+" to "+latest)
  } else {
    println(version_env_var+" up-to-date")
  }
  return "true"
}
func nomad() {
  //githubLatestRelease("hashicorp/nomad")
}
