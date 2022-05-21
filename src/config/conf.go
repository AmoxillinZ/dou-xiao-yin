/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/18 15:33
 */

package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	Database Database
	Resource Resource
	Oss      Oss
}
type Database struct {
	Dbtype   string
	Host     string
	Port     string
	Config   string
	Dbname   string
	Username string
	Password string
}
type Resource struct {
	StaticPath string `yaml:"staticPath"`
	Ip         string
	Port       string
}

type Oss struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	BucketName      string `yaml:"bucketName"`
}

// conf: 配置结构体，InitConf中初始化
var conf Conf

// InitConf : 初始化配置，从conf.yaml加载配置到conf中
func InitConf(configPath string) {
	// 加载文件
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 将读取的yaml文件解析为相应的struct
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetConf() Conf {
	return conf
}
