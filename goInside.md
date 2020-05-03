# GO internal

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

## [channel](https://tiancaiamao.gitbooks.io/go-internals/content/zh/07.1.html)

- [数据结构](https://www.youtube.com/watch?v=KBZlN0izeiY)
- buf, 环形队列
- 游标
  - sendx
  - Recvx
- waiting senders sudog![image-20200502182916711](/Users/takesachishuu/go/src/github.com/wuxiangzhou2010/golearn/image-20200502182916711.png)
- waiting receivers
- Mutex

hchan 是 chan 的结构体，在 hchan 结构中 qcount 和 elemsize 指定队列的容量和使用量，dataqsiz 队列的大小，整个 hchan 结构体只记录了队列大小相关的值，带有缓冲区的 chan 需要 make 的时候指定

- runtime.chanrecv
- runtime.chansend

reference:

- [channel 数据结构](https://zhuanlan.zhihu.com/p/27295229)
- [channel 最佳实践](https://zhuanlan.zhihu.com/p/32521576)

- notice
  - channel 的结构在堆区， make chan 返回指针
  - 没有共享的内存，除了 hchan 结构
  - 阻塞的时候调用 gopark, 换设置为 waiting, 置换新的 G
  - 恢复的时候使用 goready

## select

//select 的 case 和 default 编译器最终会编译成 if else

- selectnbsend
- selectnbrecv
- selectnbrecv2

## [interface](https://tiancaiamao.gitbooks.io/go-internals/content/zh/07.2.html)

根据 interface 是否包含 method,底层实现上用两种 struct 来表示，iface 和 eface.

```go
struct Eface
{
    Type*    type;
    void*    data;
};
struct Iface
{
    Itab*    tab;
    void*    data;
};
```

- eface 表示不含 method 地 interface 结构， 或者叫 empty interface。 对于 golang 中的大部分数据类型都可以抽象出来`_type` 结构， 同时针对不同地类型还会有一些其他信息。
- iface 表示 non-empty interface 的底层实现。 相比于 empty interface, non-empty 要包含一些 method. method 的具体实现存放在 itab.func 变量中。

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
- [《深入解析 Go》](https://tiancaiamao.gitbooks.io/go-internals/content/zh/01.0.html)
- [Golang Internals Resources](https://github.com/emluque/golang-internals-resources)
- [How to learn internals of the Go Programming Lanauge? For noob](https://stackoverflow.com/questions/49771806/how-to-learn-internals-of-of-the-go-programming-lanauge-for-noob)

## [调度器](https://www.youtube.com/watch?v=YHRO5WQGh0k)

- 调度器设计理念
  - resuse thread， 创建的 kernel thread 数量和系统 CPU 数量相同， 最大利用多核
  - 每个 CPU 有自己的 runqueue, 以减少对一个全局 queue 访问造成的性能瓶颈。 Lock.
  - 负载均衡： 使用到 work steal(when thread is busy)/ handoff (when thread is block)/ start a thread
  - 抢占长时间不交出控制权的 goroutine preemption ---> sysmon ---> global runqueue=== lower prority than local runqueue
- 调度器的缺点

  - 非强抢占 - strong preemmtion
  - runqueue 的调度室 FIFO 的方式， 并不是 Linux 类似的优先队列
  - 不是 LIFO， 缓存可能不能很好地利用。

- [Go 调度器: M, P 和 G](https://colobu.com/2017/05/04/go-scheduler/)

- G

  Goroutines are lightweight version of threads, with very low cost of starting up. Each goroutine is d`escribed by a struct called`G`, which contains fields necessary to keep track of its stack and current status. So,`G = goroutine`.

  Runtime keeps track of each `G` and maps them onto `Logical Processors`, named `P`. `P` can be seen as a abstract resource or a context, which needs to be acquired, so that `OS thread (called M, or Machine)` can execute `G` .

- [go 内部实现](https://docs.kilvn.com/go-internals/02.2.html)

- [GPM 模型](https://juejin.im/post/5e3a16e06fb9a07ccb7e8472)

![img](https://user-gold-cdn.xitu.io/2020/2/5/17012e959e4664fb?imageslim)

## Slice

- 三部分组成
- slice 的扩容
- make 和 new
  - make 返回普通类型
  - new 返回指针

## map

- 底层用 hash 表实现的

```go
struct Hmap
{
    uint8   B;    // 可以容纳2^B个项
    uint16  bucketsize;   // 每个桶的大小

    byte    *buckets;     // 2^B个Buckets的数组
    byte    *oldbuckets;  // 前一个buckets，只有当正在扩容时才不为空
};
```

## nil

- interface
  var i interface{}

只有上述的 interface 才为空。

- slice string

  - string 不能喝 nil 比较， 即使是空的 string，它的大小也是两个机器字长的
  - slice slice 也类似，它的空值并不是一个空指针，而是结构体中的指针域为空，空的 slice 的大小也是三个机器字长的。

- channel map

  - channel 跟 string 或 slice 有些不同，它在栈上只是一个指针，实际的数据都是由指针所指向的堆上面。

    - 读或者写一个 nil 的 channel 的操作会永远阻塞。
    - 读一个关闭的 channel 会立刻返回一个 channel 元素类型的零值。
    - 写一个关闭的 channel 会导致 panic。

- map map 也是指针，实际数据在堆中，未初始化的值是 nil。
- nil is useful
  - [use nil channel to disable select cases](https://youtu.be/ynoY2xz-F8s?t=1547)

## 函数调用协议

- 多值返回
  - 在传入的参数之上留了两个空位，被调者直接将返回值放在这两空位
- go 关键字
  - 如何实现 runtime.newproc 函数接受的参数分别是：参数大小，新的 goroutine 是要运行的函数，函数的 n 个参数。
  - 在 runtime.newproc 中，会新建一个栈空间，将栈参数的 12 个字节拷贝到新栈空间中并让栈指针指向参数。这时的线程状态有点像当被调度器剥夺 CPU 后一样，寄存器 PC、SP 会被保存到类似于进程控制块的一个结构体 struct G 内。f 被存放在了 struct G 的 entry 域，后面进行调度器恢复 goroutine 的运行，新线程将从 f 开始执行。
- defer 关键字
  - defer 是在 return 之前执行的。这个在 官方文档中是明确说明了的
  - 函数返回的过程是这样的：先给返回值赋值，然后调用 defer 表达式，最后才是返回到调用函数中。
  - 如何实现： runtime.deferproc
- 连续栈
  - goroutine 可以初始时只给栈分配很小的空间，然后随着使用过程中的需要自动地增长。
  - 如何捕获到函数的栈空间不足 runtime.morestack
  - 旧栈数据复制到新栈 runtime.newstack， runtime.lessstack
- 闭包的实现
  - Go 语言支持闭包
  - Go 语言能通过 escape analyze 识别出变量的作用域，自动将变量在堆上分配。将闭包环境变量在堆上分配是 Go 实现闭包的基础。
  - 返回闭包时并不是单纯返回一个函数，而是返回了一个结构体，记录下函数返回地址和引用的环境中的变量地址。

## [垃圾回收](https://www.youtube.com/watch?v=aiv1JOfMjm0)

- Go 语言使用标记清扫的垃圾回收算法
- Three phases
  - Mark setup  STW
    - turn on the write barrier, so have to stop the world
  - Marking  concurrent
    - inspect the stacks and find root pointers to the heap
    - Traverse the heap graph from those root pointers
    - make values on the heap that are still in use
  - Marking termination STW -- aka stop the world

![image-20200503101907028](/Users/takesachishuu/go/src/github.com/wuxiangzhou2010/golearn/image-20200503101907028.png)

## 内存逃逸到 heap 上

查看方法

```sh
go build -gcflags "-m -l "
```

- 逃逸可能出现的情况
  - when a value could possibly be referenced after the function that constructed the value returns
  - when the compiler determines a value is too large to fit on the stack
  - when the compiler doesn't know the size of a value at conpile time



![image-20200503093934221](/Users/takesachishuu/go/src/github.com/wuxiangzhou2010/golearn/image-20200503093934221.png)

![image-20200503093834988](/Users/takesachishuu/go/src/github.com/wuxiangzhou2010/golearn/image-20200503093834988.png)

![image-20200503094020979](/Users/takesachishuu/go/src/github.com/wuxiangzhou2010/golearn/image-20200503094020979.png)



reference：

- [深入解析 Go](https://docs.kilvn.com/go-internals/07.2.html)
