package main
import (
  "encoding/xml"
  "fmt"
  "os"
)
func jenkinsLatestRelease(version_var string, debug bool) bool {
  return jenkinsRelease(version_var,"LATEST",debug)
}
func jenkinsLTSRelease(version_var string, debug bool) bool {
  return jenkinsRelease(version_var,"LTS",debug)
}
func jenkinsRelease(version_var string, release_type string, debug bool) bool {
  type Item struct {
    Title string `xml:"title"`
  }
  type Result struct {
    XMLName xml.Name `xml:"rss"`
    Items []Item  `xml:"channel>item"`
  }
  var url string = ""
  switch release_type {
  case "LTS":
    url = "https://jenkins.io/changelog-stable/rss.xml"
  case "LATEST":
    url = "https://jenkins.io/changelog/rss.xml"
  default:
    println("Unknown jenkins release type")
    os.Exit(1)
  }
  objs := Result{}
  err := xml.Unmarshal([]byte(getUrl(url)), &objs)
  if err != nil {
    fmt.Printf("error: %objs", err)
    return false
  }
  var latest string = objs.Items[0].Title[8:]
  var current string = version_var
  if latest != current {
    println("Upgrade Jenkins from "+version_var+" to "+latest)
    return false
  } else {
    if debug {
      println("Jenkins version "+version_var+" up-to-date")
    }
    return true
  }
}
