package main
import (
  "encoding/json"
  "log"
  "reflect"
)
func githubLatestRelease(orgrepo string, version_var string) bool {
  var url string = "https://api.github.com/repos/"+orgrepo+"/tags"
  var objs interface{}
  json.Unmarshal([]byte(getUrl(url)), &objs)
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
  var current string = version_var
  if latest != current {
    println("Upgrade "+version_var+" to "+latest)
    return false
  } else {
    println(version_var+" up-to-date")
    return true
  }
}
