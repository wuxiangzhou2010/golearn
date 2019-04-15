## 协程 Coroutine

- 轻量级 “线程”
- 非抢占式多任务处理， 由协程主动交出控制权
- 编译器/解释器/虚拟机层面的多任务
- 多个协程可能在一个或多个线程上执行。

Subroutines are special cases of more genereal program conoonents, called coroutines. In contrast to the unsymmetric.

- 子程序是协程的一个特例

python 协程

- 使用 yield 关键字实现协程
- python3.5 加入了 async def 对协程原生支持

- 使用 go run --race 来检测数据访问冲突

goroutine 可能的切换点

- I/O, select
- channel
- 等待锁
- 函数调用
- runtime.Gosched()
