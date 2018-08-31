package main
import (
  "bufio"
  "log"
  "os"
  "regexp"
  "strings"
)
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
