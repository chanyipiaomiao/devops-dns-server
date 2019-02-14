#
# BUILD 阶段
#
FROM golang:1.11 AS build

# 设置我们应用程序的工作目录
WORKDIR /go/src/devops-dns-server

# 添加所有需要编译的应用代码
ADD . .

# 编译一个静态的go应用（在二进制构建中包含C语言依赖库）
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo .

ADD app.conf .
ADD hosts .

EXPOSE 53


# 设置我们应用程序的启动命令
CMD ["./devops-dns-server"]



#
# 生产阶段
#
FROM scratch AS prod

# 设置我们应用程序的工作目录
WORKDIR /devops-dns-server

# 从buil阶段拷贝二进制文件
COPY --from=build /go/src/devops-dns-server/devops-dns-server .

ADD app.conf .
ADD hosts .

EXPOSE 53

CMD ["./devops-dns-server"]




