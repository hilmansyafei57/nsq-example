package platform

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//
// Containt NSQ configuration example: "host" : "127.0.0.1:4150"
//
type NsqConfig struct {
	Host       string
	ChanelName string `json:"chanel_name"`
}

//
// Conaint container environtment of apps
//
type Environment struct {
	DirWork string
	Nsq     NsqConfig `json:"nsq"`
	Debug   int
}

//
// Create Environment config and load data from Environtment.load() function
//
func NewEnvironment() (env *Environment) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("NewEnvironment: " + err.Error())
	}

	env = &Environment{
		DirWork: wd,
	}

	env.load()

	if env.Debug >= 1 {
		fmt.Printf("Environment: %+v\n", env)
	}

	return env
}

//
// load configuration from file
//
func (env *Environment) load() {
	wd := env.DirWork
	configPath := filepath.Join(wd, pathEnv)

	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("platform: Environment.load: cannot find " + configPath)
	}

	fmt.Printf("platform: loading configuration from %q ...\n", configPath)
	err = json.Unmarshal(b, env)
	if err != nil {
		log.Fatalf("platform: Environment.load: %s: %s", configPath, err.Error())
	}

	if env.Debug == 1 {
		log.Println("Debug is active.")
		log.Println(string(b))
	}
}
