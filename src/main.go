package main

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"os"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			os.Exit('0')
		}
	}()
	filePath := os.Args[1]
	if filePath == "" {
		panic(errors.New("config file must be specify"))
	}
	var config TomlConfig
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		fmt.Println(err)
		createConfigFile(filePath)
		fmt.Println("config.toml had created from you, and you can configure it now!")
		return
	}
	client, err := alidns.NewClientWithAccessKey("cn-hangzhou", config.Access.AccessKeyId, config.Access.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	waiter := len(config.Record)
	for _, record := range config.Record {
		record := record
		go func() {
			defer func() {
				waiter--
			}()
			ip, err := getPublicNetIP(record.Interface, record.InterfaceName, config.IpFinder.Uri)
			if err != nil {
				fmt.Println(err)
				return
			}
			aliyunDDNS(client, ip, record)
		}()
	}
	for 0 != waiter {
	}
}
