package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
)

func getPublicNetIP(IF int, IFName string, uri string) (string, error) {
	var tr *http.Transport
	var lIP net.IP
	if IF <= 0 && IFName == "" {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	} else if IF > 0 {
		var err error
		lIP, err = selectInterfaceByIndex(IF)
		if err != nil {
			return "", err
		}
	} else if IFName != "" {
		var err error
		lIP, err = selectInterfaceByName(IFName)
		if err != nil {
			return "", err
		}
	}
	if lIP != nil {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			DialContext: func(ctx context.Context, network string, addr string)(net.Conn, error) {
				lAddr, err := net.ResolveTCPAddr(network, lIP.String() + ":0")
				if err != nil {
					return nil, err
				}
				rAddr, err := net.ResolveTCPAddr(network, addr)
				if err != nil {
					return nil, err
				}
				conn, err := net.DialTCP(network, lAddr, rAddr)
				return conn, err
			},
		}
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(uri)
	if err != nil {
		return "", err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	reg := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	rIP := reg.Find(body)
	if len(rIP) == 0 {
		return "", errors.New("IP is nil")
	}
	return string(rIP), nil
}

func selectInterfaceByIndex(IF int) (net.IP, error) {
	inter, err := net.InterfaceByIndex(IF)
	if err != nil {
		return nil, err
	}
	if (inter.Flags & net.FlagUp) == 0 {
		return nil, errors.New("interface not working")
	}
	addrArr, err := inter.Addrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrArr {
		if ip, bool := addr.(*net.IPNet); bool {
			if ipv4 := ip.IP.To4(); ipv4 != nil {
				return ipv4, nil
			}
		}
	}
	return nil, errors.New("interface not bind any v4 ip")
}

func selectInterfaceByName(IFName string) (net.IP, error)  {
	inter, err := net.InterfaceByName(IFName)
	if err != nil {
		return nil, err
	}
	if (inter.Flags & net.FlagUp) == 0 {
		return nil, errors.New("interface not working")
	}
	addrArr, err := inter.Addrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrArr {
		if ip, bool := addr.(*net.IPNet); bool {
			if ipv4 := ip.IP.To4(); ipv4 != nil {
				return ipv4, nil
			}
		}
	}
	return nil, errors.New("interface not bind any v4 ip")
}
