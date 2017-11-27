package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Cluster struct {
	Name       string   `yaml:"name"`
	DataCentre string   `yaml:"datacentre"`
	Nodes      []string `yaml:"nodes"`
}

type Configuration struct {
	Clusters    []Cluster `yaml:"clusters"`
	MinReplicas int       `yaml:"min_replicas"`
	MaxReplicas int       `yaml:"max_replicas"`
}

func saveConfig(c Configuration, filename string) error {
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0644)
}

func loadConfig(filename string) (Configuration, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return Configuration{}, err
	}

	var c Configuration
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return Configuration{}, err
	}

	return c, nil
}
