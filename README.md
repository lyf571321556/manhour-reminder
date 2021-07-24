# manhour-robot

## 工程构建相关

## 下载全部依赖包

***go mod tidy***

## 交叉编译时注入变量值(在main包中定义的变量通过-X main.var_name=value,在其他包下的变量通过-X package_path.variable_name=value)

go build -ldflags "-X main.Time=value -X main.User=value -X package_path.variable_name=value" -0 outputfile main.go

## 查询变量所在的package_path,分别执行一下命令

1. go build -o manhour-robot,编译出二进制可执行文件

2. go tool nm ./manhour-robot | grep manhour-robot 查找定义变量所在package_path,输出如下：

> 17d9aa0 D github.com/lyf571321556/manhour-reminder/bot..inittask   
> 142f9e0 T github.com/lyf571321556/manhour-reminder/bot.InitBot   
> 142fd20 T github.com/lyf571321556/manhour-reminder/bot.SendMsgToUser   
> 181b380 B github.com/lyf571321556/manhour-reminder/bot.wechatbot   
> 17dcee0 D github.com/lyf571321556/manhour-reminder/cmd..inittask   
> 143acc0 T github.com/lyf571321556/manhour-reminder/cmd.Execute   
> ***181bc50 B github.com/lyf571321556/manhour-reminder/cmd.Version***

3. 找到定义的Version变量的package_path，然后再次编译的时候注入变量值   
   go build -ldflags "-X 'main.time=${date}' -X 'main.user=${id -u -n}' -X
   github.com/lyf571321556/manhour-reminder/cmd.version=v1.0.2"  -o manhour-robot

4. 容器内运行   
docker build -f ./devops/Dockerfile -t lyf571321556/manhour-robot:1.0.0 .   
docker run --privileged=true --name manhour-robot -v /Users/liuyanfeng/GoWorkspace/manhour-reminder:/app/manhour-robot/ lyf571321556/manhour-robot:1.0.0
--config=./conf/dev_config.yaml start

5. docker buildx构建多平台镜像   
   docker buildx build -f ./devops/Dockerfile -t lyf571321556/manhour-robot:1.0.0 --platform=linux/arm64,linux/amd64 . --push
   //启动一个后台运行，且停止后自动重启的container
   docker run -d --restart always --name manhour-robot -v /apps:/apps  lyf571321556/manhour-robot:1.0.0 --config=/app/conf/dev_config.yaml start
   //测试container中的/app/manhour-robot程序命令
   docker exec -it 60e6fa6835ba /app/manhour-robot test --config=/apps/conf/dev_config.yaml
