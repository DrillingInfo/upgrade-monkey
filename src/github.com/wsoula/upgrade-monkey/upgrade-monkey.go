package main
import (
  "bufio"
  "encoding/json"
  "flag"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "reflect"
  "regexp"
  "strings"
  "time"
)
type people struct {
        Name string `json:"name"`
}
var config map[string]string
var conf_file string
func main() {
  flag.StringVar(&conf_file, "c", "config.txt", "config file of technologies to check for upgrades")
  flag.Parse()
  config = ReadConfig(conf_file)
  for tech := range config {
    switch tech {
    case "NOMAD":
      githubLatestRelease("hashicorp/nomad",config["NOMAD"])
    case "HASHI_UI":
      githubLatestRelease("jippi/hashi-ui",config["HASHI_UI"])
    case "CONSUL":
      githubLatestRelease("hashicorp/consul",config["CONSUL"])
    }
  }
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
func githubLatestRelease(orgrepo string, version_var string) bool {
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
  var current string = version_var
  if latest != current {
    println("Upgrade "+version_var+" to "+latest)
    return false
  } else {
    println(version_var+" up-to-date")
    return true
  }
}
func nomad() {
  //githubLatestRelease("hashicorp/nomad")
}
func ReadConfig(filename_fullpath string) map[string]string {
  prg := "ReadConfig()"
  var options map[string]string
  options = make(map[string]string)
  file, err := os.Open(filename_fullpath)
  if err != nil {
    log.Printf("%s: os.Open(): %s\n", prg, err)
    return options
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    if strings.Contains(line, "=") == true {
      re, err := regexp.Compile(`([^=]+)=(.*)`)
      if err != nil {
        log.Printf("%s: regexp.Compile(): error=%s", prg, err)
        return options
      } else {
        config_option := re.FindStringSubmatch(line)[1]
        config_value := re.FindStringSubmatch(line)[2]
        options[config_option] = config_value
        //log.Printf("%s: out[]: %s ... config_option=%s, config_value=%s\n", prg, line, config_option, config_value)
      }
    }
  }
  //log.Printf("%s: options[]: %+v\n", prg, options)
  if err := scanner.Err(); err != nil {
    log.Printf("%s: scanner.Err(): %s\n", prg, err)
    return options
  }
  return options
}
