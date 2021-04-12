package main

import (
	"fmt"
	"net"
	"os"
)

func printInterfaces() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	str := fmt.Sprintf("# IF MTU NAME MAC FLAGS\n")
	for _, v := range interfaces {
		str += fmt.Sprintf("# %d %d %s %s %s \n", v.Index, v.MTU, v.Name, v.HardwareAddr, v.Flags)
	}
	return str
}

func createConfigFile(filePath string) {
	//filePath := "config.toml"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	bytes := []byte(printInterfaces() + initialConfig)
	_, err = file.Write(bytes)
	if err != nil {
		fmt.Println(err)
		return
	}
}
