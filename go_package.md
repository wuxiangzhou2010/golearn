# [golang 的包管理](http://www.infoq.com/cn/articles/golang-package-management)

## 包管理的挑战

1. 网络环境是一个瓶颈， 尤其是遇到大量的依赖包下载的时候， 国家防火墙问题
2. 第三方包没有中央库统一管理， 要检查包的可用性，以及版本管理、 1.5 引进 vendor 目录

## 推荐工具

    glide
    gvt
    gopkg.in

[golang 中国](https://golangtc.com/members/city/%E4%B8%8A%E6%B5%B7?p=2)

[包的注意事项](http://www.cnblogs.com/dajianshi/p/3596492.html)

- 一个 package 可以有几个文件组成
- 不要求 package 名称 和所在目录一致， 但是最好一致
- 每个目录中只能有一个 package

init 同一个 go 文件的 init 从上到下依次执行
同一个 package 中的 go 文件， 根据文件名字符串到下， 依次执行。
先 import 后调用

如果存在依赖， 先调用最早被依赖的 package 中 init

先初始化包级变量和常量， 接着 init()函数

### import

\_ 仅仅执行包中的 init()函数

别名 import(f "fmt")

import 的时候其实是执行了该包里面的 init 函数，初始化了里面的变量，\_操作只是说该包引入了，只初始化里面的 init 函数和一些变量，不能通过包名来调用其它的函数，这有什么用呢？往往这些 init 函数里面是注册自己包里面的引擎，让外部可以方便的使用，就很多实现 database/sql 的引起，在 init 函数里面都是调用了 sql.Register(name string, driver driver.Driver)注册自己，然后外部就可以使用了。

### [interface{} 空接口可以代表任何类型， 有点类似于 Java 中的 Object 类](http://blog.csdn.net/chuangrain/article/details/9358737)

### interface 类型

可以被任意对象实现， 一个对象也可以实现多个接口
方法不能重载

```go
type DNSclient interface {
    Feature
    LookupIP(host string)([]net.IP, err)
}
```

interface 值
interface 组合
嵌套但是不能有重复的方法
方法相同可以相互赋值
大的 interface 可以赋值给小的 interface， 反之不行， 因为有的方法没有实现。 被赋值的接口同名方法被覆盖

### interface 查询

[reflect](http://www.cnblogs.com/coder2012/p/4881854.html)

- 从接口值到反射

```go
var x int = 1
fmt.Println("type: ", reflect.TypeOf(x))
```

- 从反射到接口值

```go
func (v Value) Interface() interface {}
// Interface 以 interface{} 返回 v 的值
y := v.Interface().(float64)
```

- 修改反射对象

```go
p := reflect.ValueOf(&x) // 获取x的地址
fmt.Println("settability of p: ", p.CanSet())
v := p.Elem()
fmt.Println("settability of v: ", v.CanSet())
```

interface{}到函数反射

some links

- [type Kind](https://golang.org/pkg/reflect/)
- [laws-of-reflection](https://blog.golang.org/laws-of-reflection)
- [go interfaces](https://www.airs.com/blog/archives/277)
- [interfaces](https://research.swtch.com/interfaces)
- [2 4 Google Understanding Go Interfaces](https://www.youtube.com/watch?v=F4wUrj6pmSI)
- [context](https://yq.aliyun.com/articles/69662)

  the bigger the interface, the weaker the abstration
  Return concrete types, receive interfaces as parameters

## strconv

    Atoi
    Itoa

## go modules(from go 1.11)

1. Module support is enabled only when the current directory is outside GOPATH/src and itself contains a go.mod file or is below a directory containing a go.mod file.
2. In module-aware mode, GOPATH no longer defines the meaning of imports during a build, but it still stores downloaded dependencies (in GOPATH/pkg/mod) and installed commands (in GOPATH/bin, unless GOBIN is set).
