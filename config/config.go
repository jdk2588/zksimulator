package config

import (
  "io/ioutil"
  "gopkg.in/yaml.v2"
)

type Settings struct {
  Servers []string `yaml:"servers"`
  NoOfClients int `yaml:"numberofclients"`
  MessageCount int `yaml:"messages"`
  Debug bool `yaml:"debug"`
}

func DefaultConfig() Settings {
  var settings Settings
  settings.Servers = []string{}
  settings.NoOfClients = 3
  settings.MessageCount = 1
  settings.Debug = true
  return settings

}

func LoadConfig(confPath string) (Settings, error) {
  var settings Settings

  configFile, err := ioutil.ReadFile(confPath)

  if err != nil {
    return settings, err
  }

  err = yaml.Unmarshal([]byte(configFile), &settings)

  if err != nil {
    return settings, err
  }

  return settings, nil
}
