package main
import (
  "flag"
  "os"
  "strconv"
)
type people struct {
  Name string `json:"name"`
}
var config map[string]string
var conf_file string
var debug bool
var upgrades bool
func main() {
  flag.StringVar(&conf_file, "c", "config.txt", "config file of technologies to check for upgrades")
  flag.Parse()
  config = ReadConfig(conf_file)
  debug, err := strconv.ParseBool(config["DEBUG"])
  if err != nil {
    debug = false
  }
  upgrades=false
  for tech := range config {
    switch tech {
    case "NOMAD":
      if !githubLatestRelease("hashicorp/nomad",config["NOMAD"],debug) {
        upgrades=true
      }
    case "CONSUL":
      if !githubLatestRelease("hashicorp/consul",config["CONSUL"],debug) {
        upgrades=true
      }
    case "TERRAFORM":
      if !githubLatestRelease("hashicorp/terraform",config["TERRAFORM"],debug) {
        upgrades=true
      }
    case "HASHI_UI":
      if !githubLatestRelease("jippi/hashi-ui",config["HASHI_UI"],debug) {
        upgrades=true
      }
    case "JENKINS_LTS":
      if !jenkinsLTSRelease(config["JENKINS_LTS"],debug) {
        upgrades=true
      }
    case "JENKINS":
      if !jenkinsLatestRelease(config["JENKINS"],debug) {
        upgrades=true
      }
    case "RUNDECK_DOCKER":
      if !githubLatestRelease("jjethwa/rundeck",config["RUNDECK_DOCKER"],debug) {
        upgrades=true
      }
    case "JAEGER":
      if !githubLatestRelease("jaegertracing/jaeger",config["JAEGER"],debug) {
        upgrades=true
      }
    case "ATHENS":
      if !githubLatestRelease("gomods/athens",config["ATHENS"],debug) {
        upgrades=true
      }
    case "DEBUG":
    default:
      println("Unknown type "+tech)
      os.Exit(1)
    }
  }
  if !upgrades {
    println("No upgrades to do")
  }
}
func nomad() {
  //githubLatestRelease("hashicorp/nomad")
}
