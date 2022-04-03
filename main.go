package main

import (
    "fmt"
    "io/ioutil"
    "gopkg.in/yaml.v3"
    "os/exec"
    "os"
    "log"
)

type serverConf struct {
    Host string `yaml:"host"`
    User string `yaml:"user"`
    Command string `yaml:"command"`
}

func readConf(filename string) (*serverConf, error) {
    buf, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    c := &serverConf{}
    err = yaml.Unmarshal(buf, c)
    if err != nil {
        return nil, fmt.Errorf("in file %q: %v", filename, err)
    }

    return c, nil
}

func main() {
    c, err := readConf("./servers.yml")
    if err != nil {
      log.Fatal(err)
    }

    cmd := exec.Command("ssh", fmt.Sprintf("%v@%v", c.User, c.Host))
    if (len(c.Command) > 0) {
      cmd = exec.Command("ssh", fmt.Sprintf("%v@%v", c.User, c.Host), "-t", c.Command)
    }
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    _ = cmd.Run() // TODO: add error checking
}
