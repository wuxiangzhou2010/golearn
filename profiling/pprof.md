# pprof

- How it works

  CPU profiling：
  The go program will stop 100 times per second. and look at the stack

- How to use

```go
import _ "net/http/pprof"
```

The package is typically only imported for the side effect of `registering its HTTP handlers`. The handled paths all begin with /debug/pprof/.

If your application is not already running an http server, you need to start one.

```go
go func() {
http.ListenAndServe("localhost:3999", nil)
}()
```

- Then use the pprof tool to look at the `heap profile`:

```sh
go tool pprof http://localhost:6060/debug/pprof/heap
```

- to look at`a 30-second CPU profile`:

```sh
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

- to look at the goroutine `blocking profile`, after calling runtime.SetBlockProfileRate in your program:

```sh
go tool pprof http://localhost:6060/debug/pprof/block
```

- to collect `a 5-second execution trace`:

```sh
wget http://localhost:6060/debug/pprof/trace?seconds=5
```

- to look at `the holders of contended mutexes`, after calling runtime.SetMutexProfileFraction in your program:

```sh
go tool pprof http://localhost:6060/debug/pprof/mutex
```

To view all available profiles, open http://localhost:6060/debug/pprof/ in your browser.

- 查看对象分配

```sh
go tool pprof --alloc_objects http://localhost:6060/debug/pprof/heap
go tool pprof http://localhost:6060/debug/pprof/heap
```

- 查看内存占用

```sh
go tool pprof -inuse_space http://127.0.0.1:6060/debug/pprof/heap
go tool pprof -alloc_space -cum -svg http://127.0.0.1:6060/debug/pprof/heap > heap.svg
```

- 查看 goroutine 栈：

```sh
go tool pprof http://127.0.0.1:6060/debug/pprof/goroutine?debug=1
```

- /debug/pprof/profile：访问这个链接会自动进行 CPU profiling，持续 30s，并生成一个文件供下载
- /debug/pprof/block：Goroutine 阻塞事件的记录。默认每发生一次阻塞事件时取样一次。
- /debug/pprof/goroutines：活跃 Goroutine 的信息的记录。仅在获取时取样一次。
- /debug/pprof/heap： 堆内存分配情况的记录。默认每分配 512K 字节时取样一次。
- /debug/pprof/mutex: 查看争用互斥锁的持有者。
- /debug/pprof/threadcreate: 系统线程创建情况的记录。 仅在获取时取样一次。

- 常用交互分析命令

```sh
topN
web
list
```

参考：

- [golang pprof 性能分析工具](https://blog.csdn.net/moxiaomomo/article/details/77096814)
- [profiling-go-programs](https://blog.golang.org/profiling-go-programs)

---

## FlameGraph

- 安装

```sh
go get -v github.com/brendangregg/FlameGraph
go get -v github.com/uber/go-torch
```

- CPU profiling

```sh
go-torch -u http://localhost:6060/debug/pprof/profile
```

- Heap profiling

all time

```sh
go-torch -alloc_space http://localhost:6060/debug/pprof/heap --colors=mem
```

onetime

```sh
go-torch -inuse_space http://localhost:6060/debug/pprof/heap --colors=mem
```

reference:

- [brendangregg](http://www.brendangregg.com/flamegraphs.html)
- [Go:strings.Replace 与 bytes.Replace 调优](https://zhuanlan.zhihu.com/p/56217644)
- [Profiling your Golang app in 3 steps](https://coder.today/tech/2018-11-10_profiling-your-golang-app-in-3-steps/)

- [GopherCon 2018: George Tankersley - Micro optimizing Go Code](https://www.youtube.com/watch?time_continue=102&v=keydVd-Zn80)

write benchmark

1. inline
1. eliminate bound check --> use fixed array
1. allocations and copies
   1. runtime.malloc
   2. runtime.makeslice
   3. runtime.memmove
