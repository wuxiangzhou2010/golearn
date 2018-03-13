### Compile go

1. go/gccgo
2. supported arch amd64/386/arm/arm64/ppc64/ppc64le/mips/mipsle/mips64/,ips64le/s39x
3. build the bootstrap(go1.4)/crosscompile
4. fetch the source 
 git clone https://go.googlesource.com/go
5.  checkout branch 
6. build
cd src
./all.bash

environment
    $GOOS $GOARCH
    $GOROOT_BOOTSTRAP

###  [Install go ](https://golang.org/doc/install)
Download [link](https://golang.org/dl/)
```
wget https://dl.google.com/go/go1.10.linux-amd64.tar.gz
# creating a Go tree in /usr/local/go
tar -C /usr/local -xzf go1.10.linux-amd64.tar.gz

# Add /usr/local/go/bin to the PATH environment variable.
export PATH=$PATH:/usr/local/go/bin
export GOPATH=~/go
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin

```
### Installing to a custom location
```
# need to set the following
export GOROOT=$HOME/go1.X
export PATH=$PATH:$GOROOT/bin
```

Uninstalling Go
```
rm -rf /usr/local/go
# Remove  PATH environment variable
```

[ A tour of Go](https://tour.golang.org/list)


[go command](https://www.imooc.com/video/7208)
```
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

## [Go by examples](https://gobyexample.com/)

- type interface {}
- channel
- buffered channel
- channel-directions
- [select](https://gobyexample.com/select)
- [timeout](https://gobyexample.com/timeouts)
- [unblocking_channel](https://gobyexample.com/non-blocking-channel-operations)
- closing channel
- range over channel 
- timer
- ticker
  
    time.NewTicker(500 * time.Millisecond)

- sync/atomic
  
    atomic.AddUint64(&ops, 1)
    atomic.LoadUint64(&ops)
- mutex
- stateful goroutines
- sorting
- panic

- Defer 

Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.

- regular expression
- Json
- time/ time parsing
- utl/url parsing
- sha1 hash
- base64 encoding
- reading/writing files
- line filters
- command line arguments
- environment variables
- executing processes
- signals os/signal syscall
- exit