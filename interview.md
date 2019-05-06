# Interview

## [gotchas-and-common-mistakes-in-go-golang](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html)

## 分析变量的覆盖

go tool vet -shadow your_file.go 分析变量的覆盖， 必须加上 `-shadow`

## 逃逸分析

go run -gcflags -m main.go

go tool vet -shadow your_file.go

The scheduler will run after GC, "go" statements, blocking channel operations, blocking system calls, and lock operations. It may also run when a non-inlined function is called.

- [the-way-to-go](https://www.kancloud.cn/kancloud/the-way-to-go/81396)
- [Go 语言--空结构体 struct{}解析](https://blog.csdn.net/qq_34777600/article/details/87195673)

  struct 空的大小为 0，所以这样可以在一定程度上减少内存使用，特别是在消息管道开辟数量到达一定量级之后。

- [你遇到过哪些高质量的 Go 语言面试题？](https://www.zhihu.com/question/60952598)
- [golang-interviews](https://wuyin.io/2018/03/16/golang-interviews/)
- [Go 面试题答案与解析](https://yushuangqi.com/blog/2017/golang-mian-shi-ti-da-an-yujie-xi.html)
- [笔试面试知识整理](https://hit-alibaba.github.io/interview/)
- [golang面试题解析](https://blog.51cto.com/qiangmzsx/1949904)