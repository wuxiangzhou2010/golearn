# go commands

go command: `https://www.imooc.com/video/7208`

``` sh
go env GOPATH
go env GOROOT

go build
    编译源码文件， 代码包， 代码包不产生任何可执行文件，
go run
  -a 强制编译相关代码
  -n/x 打印编译过程中所运行的命令， 而不实际执行, -x 真正执行
  -p n  b 并行编译，n 为并行的数量
  -v 列出被编译的代码包名称
  -a -v 列出所有被编译的代码包的名称
  -work 显示编译时船舰的临时工作目录的路径， 并且不删除它、

go install 编译并安装源码文件或者代码包

go get  从远程代码仓库下载并安装
    -d  只下载， 不安装
    -fix 下载后进行修正动作
    -u update
```

go build tags: `https://golang.org/pkg/go/build/`
