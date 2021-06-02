package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Rule struct {
	Name string `yaml:"name"`
	Src  string `yaml:"src"`
	Dst  string `yaml:"dst"`
}

type Config struct {
	FirewallNetworkRules []Rule `yaml:"firewall_network_rules"`
}

func main() {
	filePtr := flag.String("file", "creds.yml", "file name containing credentials")

	flag.Parse()
	fmt.Println("file: ", *filePtr)
	filename, _ := filepath.Abs(*filePtr)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %v\n", config.FirewallNetworkRules)
}
