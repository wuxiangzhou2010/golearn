- [Golang 实现 epoll 代码](https://zhuanlan.zhihu.com/p/27072761)

网络轮询器

golang 中所有文件描述符都被设置成`非阻塞`的，某个 goroutine 进行网络 io 操作，读或者写文件描述符，如果此刻网络 io 还没准备好，则这个 goroutine 会被放到系统的`等待队列`中，这个 goroutine 失去了运行权，但并不是真正的整个系统“阻塞”于系统调用，后台还有一个 poller 会不停地进行 poll，所有的文件描述符都被添加到了这个 poller 中的，当某个时刻一个文件描述符准备好了，poller 就会唤醒之前因它而阻塞的 goroutine，于是 goroutine 重新运行起来。

netpoll

1. 创建 epoll

```go
func epollcreate(size int32) int32
func epollcreate1(flags int32) int32
```

1. 设置 epoll 事件

```go
func epollctl(epfd, op, fd int32, ev *epollevent) int32
```

1. 等待 epoll 事件

```go
func epollwait(epfd int32, ev *epollevent, nev, timeout int32) int32
```

reference:

- [linux 下非阻塞 io 库 epoll](https://zhuanlan.zhihu.com/p/27050330)

## channel

环形队列， 游标

hchan 是 chan 的结构体，在 hchan 结构中 qcount 和 elemsize 指定队列的容量和使用量，dataqsiz 队列的大小，整个 hchan 结构体只记录了队列大小相关的值，带有缓冲区的 chan 需要 make 的时候指定

runtime.chanrecv
chansend

referenc:

- [channel 数据结构](https://zhuanlan.zhihu.com/p/27295229)

## select

selectnbsend
selectnbrecv
selectnbrecv2

## interface

根据 interface 是否包含 method,底层实现上用两种 struct 来表示，iface 和 eface.

- iface 表示不含 method 地 interface 结构， 或者叫 empty interface。 对于 golang 中地发部分数据类型都可以抽象出来`_type` 结构， 同时针对不同地类型还会有一些其他信息。
- eface 表示 non-empty interface 的底层实现。 相比于 empty interface, non-empty 要包含一些 method. method 的具体实现存放在 itab.func 变量中。

```go
func convT2E(t *type, elem unsafe.Pointer) (e eface) {
    x:= newobject(t)
    typedmmemove(t, x, elem)
    ...
}

func convTI(tab *itab, elem unsafe.Pointer) (i iface) {
    ...
    x:= newobject(t)
    typedmemmove(t, x, elem)
    ...
}
```

reference:

- [Go Interface 源码剖析](https://zhuanlan.zhihu.com/p/27652856)