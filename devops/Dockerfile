FROM golang:1.16.6 as build
COPY . . /build/
WORKDIR /build
#使用buildx构建多平台镜像(GOOS=linux GOARCH=amd64)
#scratch为基础镜像指定一个用户来运行程序(echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd && cat /etc_passwd && )
RUN export GO111MODULE=on && export GOPROXY=https://goproxy.cn && CGO_ENABLED=0 go build -ldflags "-X 'main.time=$(date)' -X 'main.user=$(id -u -n)' -X github.com/lyf571321556/manhour-reminder/cmd.version=v1.0.2" -o manhour-robot
# 运行阶段
#scratch 是内置关键词，并不是一个真实存在的镜像。 FROM scratch 会使用一个完全干净的文件系统，不包含任何文件。 因为Go语言编译后不需要运行时，也就不需要安装任何的运行库。 FROM scratch 可以使得最后生成的镜像最小化，其中只包含了 server 程序。
FROM alpine:3.14
# 从编译阶段的中拷贝编译结果到当前镜像中
# 应用程序需要发出HTTPS请求,HTTP库需要建立SSL连接所需的证书,而scratch是一个空白虚拟镜像，因此需要从上一步镜像（golang:1.16.6 as build）复制证书到过来
#COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#scratch为基础镜像指定一个用户来运行程序
#COPY --from=build /etc_passwd /etc/passwd
RUN mkdir /app
COPY --from=build /build/manhour-robot /app/
RUN ls /app && chmod +x /app/manhour-robot  && ls /app
#scratch为基础镜像指定一个用户来运行程序
#USER nobody
ENTRYPOINT ["/app/manhour-robot"]
CMD ["start"]
