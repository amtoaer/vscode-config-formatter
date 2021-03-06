package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"strings"

	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
)

const (
	defaultConfigPath = "~/AppData/Roaming/Code/User/settings.json"
)

func handlePath(path string) (string, error) {
	if path == "" {
		return homedir.Expand(defaultConfigPath)
	}
	return homedir.Expand(path)
}

func format(bytes []byte, tabs bool, indent int) (result []byte, err error) {
	var m map[string]interface{}
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		return
	}
	if tabs {
		return json.MarshalIndent(m, "", "\t")
	}
	var sb strings.Builder
	for i := 0; i < indent; i++ {
		sb.WriteByte(' ')
	}
	return json.MarshalIndent(m, "", sb.String())
}

func main() {
	var (
		path   = flag.String("p", defaultConfigPath, "")
		tabs   = flag.Bool("t", false, "use tabs or not")
		indent = flag.Int("i", 4, "count of space")
		err    error
		bytes  []byte
	)
	flag.Parse()
	*path, err = handlePath(*path)
	if err != nil {
		color.Red("[error] error handle path: %v\n", err)
		return
	}
	bytes, err = ioutil.ReadFile(*path)
	if err != nil {
		color.Red("[error] error read configuration: %v\n", err)
		return
	}
	bytes, err = format(bytes, *tabs, *indent)
	if err != nil {
		color.Red("[error] error format configuration: %v\n", err)
		return
	}
	err = ioutil.WriteFile(*path, bytes, 0777)
	if err != nil {
		color.Red("[error] error write configuration: %v\n", err)
		return
	}
	color.Yellow("[success] process successfully")
}
