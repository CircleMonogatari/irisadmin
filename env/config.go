package env

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)


var Config = struct {
	Appname string `json:"AppName"`
	Version string `json:"version"`
	Mysql   struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"mysql"`
	Redis struct {
		IP       string `json:"ip"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"redis"`
}{}



func ConfigInit(fileName string)  {
	fmt.Println("初始化 env config")
	if fileName == "" {
		panic("配置文件 config.json路径不能为``")
	}

	file, err := os.Open(fileName)
	if err != nil {
		panic("文件打开失败 Err:"+err.Error())
	}

	data, err:=io.ReadAll(file)
	if err != nil {
		panic("文件读取失败 Err:"+ err.Error())
	}

	err = json.Unmarshal(data, &Config)
	if err != nil {
		panic("配置文件解析失败 Err:"+ err.Error())
	}

	fmt.Printf("%+v\n", Config)
}
