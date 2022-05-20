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
	Resource string
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

func GetConf() Conf {
	var conf Conf
	// 加载文件
	yamlFile, err := ioutil.ReadFile("./src/config/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	// 将读取的yaml文件解析为响应的 struct
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	return conf
}