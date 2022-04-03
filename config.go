package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Shell    string `yaml:"shell"`
	Location string `yaml:"location"`
	Command  string `yaml:"command"`
}

func readConf(fpath string) (*Config, error) {
	buf, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(buf, cfg)
	if err != nil {
		return nil, err
	}

	if len(cfg.Shell) == 0 {
		cfg.Shell = "bash"
	}

	cmd := ""
	if len(cfg.Location) > 0 {
		cmd += fmt.Sprintf("cd %s && ", cfg.Location)
	}

	if len(cfg.Command) > 0 {
		cmd += cfg.Command
	} else {
		cmd += cfg.Shell
	}

	cfg.Command = cmd

	return cfg, nil
}
