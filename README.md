# manhour-reminder

## 工程构建相关

## 下载全部依赖包

***go mod tidy***

## 交叉编译时注入变量值(在main中定义的变量通过-X main.var_name=value,在其他文件下)

go build -ldflags "-X main.Time=value -X main.User=value -X package_path.variable_name=value" -0 outputfile main.go

## 查询变量所在的package_path,分别执行一下命令

1. go build -o manhour-reminder,编译出二进制可执行文件

2. go tool nm ./manhour-reminder | grep manhour-reminder 查找定义变量所在package_path,输出如下：

> 17d9aa0 D github.com/lyf571321556/manhour-reminder/bot..inittask   
> 142f9e0 T github.com/lyf571321556/manhour-reminder/bot.InitBot   
> 142fd20 T github.com/lyf571321556/manhour-reminder/bot.SendMsgToUser   
> 181b380 B github.com/lyf571321556/manhour-reminder/bot.wechatbot   
> 17dcee0 D github.com/lyf571321556/manhour-reminder/cmd..inittask   
> 143acc0 T github.com/lyf571321556/manhour-reminder/cmd.Execute   
> ***181bc50 B github.com/lyf571321556/manhour-reminder/cmd.Version***

3. 找到定义的Version变量的package_path，然后再次编译的时候注入变量值   
   go build -ldflags "-X main.Time=value -X main.User=value -X
   github.com/lyf571321556/manhour-reminder/cmd.Version=v1.0.2"  -o manhour-reminder   