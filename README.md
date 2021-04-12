# AliyunDDNS_Multi
Aliyun DDNS. Supports multiple network interfaces for independent dynamic resolution of public network address.

## 功能

- 支持绑定指定网卡的域名解析


## 使用方法

初始化config.toml文件（文件名可以自行指定）

```shell
aliyunddns config.toml
```

初始化文件生成以后，根据提示进行编辑

```toml
# IF MTU NAME MAC FLAGS
# 1 -1 Loopback Pseudo-Interface 1  up|loopback|multicast 

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
```

编辑完成后 运行

```shell
aliyunddns config.toml
```