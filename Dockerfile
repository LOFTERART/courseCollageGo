#源镜像
FROM golang:latest
#作者
MAINTAINER Razil "niyuelin1990@163.com"
#设置工作目录
WORKDIR /Users/metis/Desktop/PC
#将服务器的go工程代码加入到docker容器中
ADD . /Users/metis/Desktop/PC
#go构建可执行文件
RUN go mod tidy
RUN go build .
#暴露端口
EXPOSE 6064
#最终运行docker的命令
ENTRYPOINT  ["./pc"]
