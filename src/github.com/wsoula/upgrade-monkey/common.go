package main
import (
  "io/ioutil"
  "log"
  "net/http"
  "time"
)
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
