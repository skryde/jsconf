# JS{c}ON{f} = JSON configuration

Save an `struct` to a JSON file; and then read the JSON file and save its content into an `struct`.

---

# How to install

`go get -v -u github.com/skryde/jsconf`

---

# Usage

```golang

package main

import (
	"fmt"
	"github.com/skryde/jsconf"
)

const configFile = "conf.json"

type configuration struct {
	Port     int    `json:"port"`
	LogFile  string `json:"logfile"`
	Database struct {
		Host string `json:"host"`
		Port string `json:"port"`
		User string `json:"user"`
		Pass string `json:"pass"`
	}               `json:"database"`
}

func main() {
	var c configuration

	if res, err := jsconf.Exist(configFile); err == nil && res == jsconf.IsFile {
		erro := jsconf.LoadFromFile(configFile, c)
		if erro != nil {
			panic(erro)
		}

		fmt.Println(c)

	} else if res == jsconf.Error {
		c.Port = 8080
		c.LogFile = "app.log"
		c.Database.Host = "127.0.0.1"
		c.Database.Port = "3546"
		c.Database.User = "usuario"
		c.Database.Pass = "secret"

		erro := jsconf.SaveToFile(configFile, c)
		if erro != nil {
			panic(erro)
		}

		fmt.Println(c)
	}
}

```
