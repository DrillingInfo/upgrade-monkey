Upgrade Monkey
===

Like the ci-monkey that goes around finding repos to give ci builds to automatically,
the upgrade-monkey goes around finding infrastructure to upgrade.

Design
---
* different "modules" for each upgradeable piece
  * jenkins "module" goes to find if there are updates for jenkins and does them
  * module could be file or function or separate repo
* current version is stored in env vars but if this does the upgrade how does it update that

Roadmap
---
- [ ] Nomad
  - [x] Get list of releases
  - [x] compare to current version
  - [ ] alert
- [ ] Hashi-UI
- [ ] Consul

Repo Notes
---
* I don't know go, using this to learn
* Folder structure is because of this: https://golang.org/doc/code.html, specifically the "Your first program" section

```
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:~/git/upgrade-monkey/bin
export GOPATH=~/git/upgrade-monkey
vi src/github.com/wsoula/upgrade-monkey/upgrade-monkey.go
go install github.com/wsoula/upgrade-monkey
bin/upgrade-monkey
```

Links I Used
---
* https://blog.alexellis.io/golang-json-api-client/
* https://stackoverflow.com/questions/42377989/unmarshal-json-array-of-arrays-in-go
* https://stackoverflow.com/questions/29463791/golang-reading-an-array-of-json-objects
