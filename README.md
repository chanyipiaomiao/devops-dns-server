## devops-dns-server

devops-dns-server 是一款简单易用的dns服务器，支持从外部数据源获取数据(比如从CMDB中获取数据,只要API返回固定格式的数据)、支持读取某个文件(必须是/etc/hosts文件一样的格式), 从而可以使解析主机名为IP.

也可以转发DNS请求到指定的DNS服务器, 可以作为内网DNS服务器

## 配置

配置文件app.conf说明

```sh
[server]
# 本服务器的监听地址
# 192.168.1.1:53  只监听该地址 192.168.1.1 的53
# :53             所有接口地址都监听53端口
# 首先从环境变量中获取 DEVOPS_DNS_SERVER_LISTEN 的值 如果没有该值则使用默认的 :53
listen = "${DEVOPS_DNS_SERVER_LISTEN||:53}"

# 当本服务器无法解析的时候会向这个服务器转发DNS查询
nameserver = "${DEVOPS_DNS_SERVER_NAMESERVER||8.8.8.8:53}"

[source]
# 解析的顺序,可以这样设置
# fromFile: 只从文件获取IP
# fromAPI: 只从API获取IP
# fromFile,fromAPI: 首先从文件获取 如果文件里面没有再从API获取
order = "${DEVOPS_DNS_SERVER_ORDER||fromFile}"

[fromAPI]
# api的url地址, 返回的json数据格式如下:
# 必须包含data字段
#{
#    "data": "ip地址"
#    ...
#}
url = "${DEVOPS_DNS_SERVER_FROMAPI_URL||http://192.168.2.116:30080/host/findIPByName?TOKEN=vZQKuspMoUdxDVe}"

[fromFile]
# 文件格式必须要和/etc/hosts格式一样
filepath = "${DEVOPS_DNS_SERVER_FROMFILE||1.txt}"

# 是否要监控该文件,一有修改就重新读取
watch = "${DEVOPS_DNS_SERVER_FROMFILE_WATCH||yes}"
```

## 运行

### Docker

```sh
docker run -p 53:53/udp  chanyipiaomiao/devops-dns-server:latest
```

### 下载[二进制](https://github.com/chanyipiaomiao/devops-dns-server/releases)包直接运行

```bash
./devops-dns-server
```
