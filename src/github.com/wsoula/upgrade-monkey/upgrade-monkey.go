package main
import (
  "flag"
  "strconv"
)
type people struct {
  Name string `json:"name"`
}
var config map[string]string
var conf_file string
var debug bool
func main() {
  flag.StringVar(&conf_file, "c", "config.txt", "config file of technologies to check for upgrades")
  flag.Parse()
  config = ReadConfig(conf_file)
  //debug = config["DEBUG"].(bool)
  //debug = strconv.ParseBool(config["DEBUG"])
  debug, err := strconv.ParseBool(config["DEBUG"])
  if err != nil {
    debug = false
  }
  //debug = config["DEBUG"].(bool)
  for tech := range config {
    switch tech {
    case "NOMAD":
      githubLatestRelease("hashicorp/nomad",config["NOMAD"],debug)
    case "HASHI_UI":
      githubLatestRelease("jippi/hashi-ui",config["HASHI_UI"],debug)
    case "CONSUL":
      githubLatestRelease("hashicorp/consul",config["CONSUL"],debug)
    case "JENKINS_LTS":
      jenkinsLTSRelease(config["JENKINS_LTS"],debug)
    case "JENKINS":
      jenkinsLatestRelease(config["JENKINS"],debug)
    case "RUNDECK_DOCKER":
      githubLatestRelease("jjethwa/rundeck",config["RUNDECK_DOCKER"],debug)
    case "DEBUG":
    default:
      println("Unknown type "+tech)
    }
  }
}
func nomad() {
  //githubLatestRelease("hashicorp/nomad")
}
