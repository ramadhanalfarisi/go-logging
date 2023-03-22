# go-logging
The go-logging package used to log all errors, infos, warnings, and debugs in your code

### Install
`go get github.com/ramadhanalfarisi/go-logging`

### Preparation
Create your log folder in your root code or anywhere, in the case, i created my log folder in my root path

### Simple Example Usage
```go
package main

import (
	"fmt"
	go_logging "github.com/ramadhanalfarisi/go-logging"
)


func main(){
	get_config := go_logging.ConfigLogging(true, "logs/")
	get_config.Debug("This is debug")
	get_config.Info("This is info")
	get_config.Warning("This is warning")
	get_config.Error(fmt.Errorf("This is error"))
}
```
