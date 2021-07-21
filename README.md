# manhour-reminder

> 工程构建相关
*下载全部依赖包   
go mod tidy   
*交叉编译时注入版本号   
go build -ldflags "-X main.version=v1.0.1"  -o manhour-reminder main.go