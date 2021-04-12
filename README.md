# AliyunDDNS_Multi
Aliyun DDNS. Supports multiple network interfaces for independent dynamic resolution of public network address.

## 功能

- 支持绑定指定网卡的域名解析

## 编译

```
go get github.com/aliyun/alibaba-cloud-sdk-go
```

> coder只知道怎么下载依赖，不知道怎么编译 （没错，由于用的idea，没有从命令行做过编译，所以我也不知道怎么通过命令行编译它），所以自行研究吧。（coder真的太懒太懒了）

## 使用方法

初始化config.toml文件（文件名可以自行指定）

```shell
aliyunddns config.toml
```

初始化文件生成以后，根据提示进行编辑。

> *文件会默认生成网卡信息，可以根据生成的网卡信息配置IpFinder中的参数。*

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