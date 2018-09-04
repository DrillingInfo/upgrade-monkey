package main
import (
  "flag"
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
    case "JENKINS_LTS":
      jenkinsLTSRelease(config["JENKINS_LTS"])
    case "JENKINS":
      jenkinsLatestRelease(config["JENKINS"])
    case "RUNDECK_DOCKER":
      githubLatestRelease("jjethwa/rundeck",config["RUNDECK_DOCKER"])
    default:
      println("Unknown type")
    }
  }
}
func nomad() {
  //githubLatestRelease("hashicorp/nomad")
}
