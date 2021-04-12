package main

type TomlConfig struct {
	Access AliyunAccess
	IpFinder IpFinder
	Record DNSRecord
}

type AliyunAccess struct {
	AccessKeyId string
	AccessKeySecret string
}

type IpFinder struct {
	Interface int
	InterfaceName string
	Uri string
}

type DNSRecord struct {
	Domain string
	RR     string
	Type string
	Priority int64
	TTL int64
}

var initialConfig =
`
[Access]
# Aliyun AccessKey
AccessKeyId=""
# Aliyun AccessKeySecret
AccessKeySecret=""

[IpFinder]
# Interface Index. (If You don't know what is it, just set it to 0.)
# Set value less than or equal to 0 to disable. (suggest set 0)
interface=0
# Interface Name. (If You don't know what is it, just set it to "" (empty string))
# If you specify the interface index, interfaceName wouldn't be work.
# Set value "" (empty string ) to disable.
interfaceName=""
# Website which use to find public ip 
uri="http://www.net.cn/static/customercare/yourip.asp"

# Record
[Record]
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
`
