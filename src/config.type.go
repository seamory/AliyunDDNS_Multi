package main

type TomlConfig struct {
	Access AliyunAccess
	IpFinder IpFinder
	Record map[string]DNSRecord
}

type AliyunAccess struct {
	AccessKeyId string
	AccessKeySecret string
}

type IpFinder struct {
	Uri string
}

type DNSRecord struct {
	Interface int `toml:"if"`
	InterfaceName string `toml:"ifName"`
	Domain string
	RR     string
	Type string
	Priority int64
	TTL int64
}

var initialConfig =
`
[access]
# Aliyun AccessKey
AccessKeyId=""
# Aliyun AccessKeySecret
AccessKeySecret=""

[ipFinder]
# Website which use to find public ip 
uri="http://www.net.cn/static/customercare/yourip.asp"

# Record
[record.ex1]
# Interface Index. (If You don't know what is it, just set it to 0.)
# Set value less than or equal to 0 to disable. (suggest set 0)
if=0
# Interface Name. (If You don't know what is it, just set it to "" (empty string))
# If you specify the interface index, interfaceName wouldn't be work.
# Set value "" (empty string ) to disable.
ifName=""
# Domain
domain=""
# RR
rr=""
# Type
type="A"
# Priority 
priority=1
# Time to life
ttl=600

# [record.ex2]
# Interface Index. (If You don't know what is it, just set it to 0.)
# Set value less than or equal to 0 to disable. (suggest set 0)
# if=0
# Interface Name. (If You don't know what is it, just set it to "" (empty string))
# If you specify the interface index, interfaceName wouldn't be work.
# Set value "" (empty string ) to disable.
# ifName=""
# Domain
# domain=""
# RR
# rr=""
# Type
# type="A"
# Priority 
# priority=1
# Time to life
# ttl=600
`
