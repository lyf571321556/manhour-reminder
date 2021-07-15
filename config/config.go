package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

const (
	defaultConfigPath = "config/config.yaml"
)

type Config struct {
	OnesProjectUrl string    `yaml:"ones_project_url"`
	BotList        []BotInfo `yaml:"bot_list"`
}

type BotInfo struct {
	BotName        string        `yaml:"bot_name"`
	BotKey         string        `yaml:"bot_key"`
	DepartmentUUID string        `yaml:"department_uuid"`
	UserMappings   []UserMapping `yaml:"user_mappings"`
}

type UserMapping struct {
	OnesUserid   string `yaml:"ones_userid"`
	WechatUserid string `yaml:"wechat_userid"`
}

func (config *Config) toString() string {
	return "test"
}

var AppConfig Config

func init() {
	yamlFile, err := ioutil.ReadFile(defaultConfigPath)
	if err != nil {
		fmt.Printf("failed to read yaml file : %v\n", err)
		return
	}

	err = yaml.Unmarshal(yamlFile, &AppConfig)
	if err != nil {
		log.Println(err)
		return
	}
}
