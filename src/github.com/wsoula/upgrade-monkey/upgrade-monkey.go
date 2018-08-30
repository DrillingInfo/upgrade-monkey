package main
import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "reflect"
  "net/http"
  "time"
)
type people struct {
        Name string `json:"name"`
}
func main() {
  url := "https://api.github.com/repos/hashicorp/nomad/tags"
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
  var objs interface{}
  var tags []string
  json.Unmarshal([]byte(body), &objs)
  objArr, ok := objs.([]interface{})
  if !ok {
    log.Fatal("expected an array of objects")
  }
  for i, obj := range objArr {
    obj, ok := obj.(map[string]interface{})
    if !ok {
      log.Fatalf("expected type map[string]interface{}. got %s", reflect.TypeOf(objArr[i]))
    }
    for key, value := range obj {
      if key == "name" {
        println(value.(string))
        tags = append(tags,value.(string))
      }
    }
  }
  fmt.Printf("%v",tags)
}
