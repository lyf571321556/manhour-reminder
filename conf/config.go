package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

const (
	defaultConfigPath = "conf/config.yaml"
)

type Config struct {
	OnesProjectUrl string      `yaml:"ones_project_url"`
	TeamUUID       string      `yaml:"team_uuid"`
	MsgContent     string      `yaml:"msg_content"`
	TaskCrontab    string      `yaml:"task_crontab"`
	LogPath        string      `yaml:"log_path"`
	Debug          bool        `yaml:"debug"`
	RobotList      []RobotInfo `yaml:"robot_list"`
}

type RobotInfo struct {
	RobotName      string        `yaml:"robot_name"`
	RobotKey       string        `yaml:"robot_key"`
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

func Init(configPath string) (err error) {
	if configPath == "" {
		configPath = defaultConfigPath
	}
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("failed to read yaml file : %+v\n", err)
		return err
	}

	err = yaml.Unmarshal(yamlFile, &AppConfig)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}
