package main

import (
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"strconv"
)

func aliyunDDNS(client *alidns.Client, ip string, config DNSRecord) {
	record, err := describeSubDomainRecords(client, config)
	if err != nil {
		fmt.Println("解析记录不存在，添加记录")
		addDomainRecord(client, ip, config)
		return
	}
	if record.Value != ip {
		fmt.Println("解析记录地址与公网地址不一致，更新记录")
		updateDomainRecord(client, record, ip, config)
		return
	}
	fmt.Println("解析记录地址与公网地址一致")
}

func describeSubDomainRecords(client *alidns.Client, config DNSRecord) (alidns.Record, error) {
	request := alidns.CreateDescribeSubDomainRecordsRequest()
	request.Scheme = "https"

	request.SubDomain = config.RR + "." + config.Domain
	response, err := client.DescribeSubDomainRecords(request)
	if err != nil {
		panic(err.Error())
	}
	if len(response.DomainRecords.Record) != 0 {
		return response.DomainRecords.Record[0], nil
	}
	return alidns.Record{}, errors.New("record is nil")
}

func addDomainRecord(client *alidns.Client, ip string, config DNSRecord) (string, string) {
	request := alidns.CreateAddDomainRecordRequest()
	request.Scheme = "https"

	request.DomainName = config.Domain
	request.RR = config.RR
	request.Type = config.Type
	request.Value = ip
	request.Priority = requests.Integer(strconv.FormatInt(config.Priority, 10))
	request.TTL = requests.Integer(strconv.FormatInt(config.TTL, 10))

	response, err := client.AddDomainRecord(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println("添加记录完成")
	return response.RecordId, response.RequestId
}

func updateDomainRecord(client *alidns.Client, record alidns.Record, ip string, config DNSRecord) (string, string) {
	request := alidns.CreateUpdateDomainRecordRequest()
	request.Scheme = "https"

	request.RecordId = record.RecordId
	request.RR = config.RR
	request.Type = config.Type
	request.Value = ip
	request.Priority = requests.Integer(strconv.FormatInt(config.Priority, 10))
	request.TTL = requests.Integer(strconv.FormatInt(config.TTL, 10))

	response, err := client.UpdateDomainRecord(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println("更新记录完成")
	return response.RecordId, response.RequestId
}
