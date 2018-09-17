# go

## variables

2.3.2 a short variable declearation must declear at least one new variable, otherwise it will not compile.

## 2.6 Packages and Files

- 2.6.1 Imports
- 2.6.2 Package Initialization
- Scope

## types

## Basic data types

- numbers
  - int
  - float:
    Go provides two sizes of floating-point numbers, float32 and float64
    math.MaxFloat32/math.MaxFloat64

- booleans: trure or false
- strings: immutable sequence of bytes, can use index s[1], len, s[i:j]
    four standard packages are particularly important for manipulating strings: bytes, strings, strconv and unicode.

- constants constant generator iota(from 0)

## aggregate types 聚合类型

fixed size

- arrays: fixed length, a[0], len(a), range a,
- structs

## Reference types 引用类型（composite types）

Pointers, slices, maps, functions, channels

they all refer to program variables or state indirectly,

for reference type the empty value is nil

- slice []string, append, s[1],range s, len(s)

- array has its fixed length while slice is not

- map: use make, range, m['a']

m = make(map[string]int)
m["k1"] = 7
m["k2"] = 8
len(m)
delete["k1"]
_, pres = m["k2"]
    pres indicate whether the value presents with the given key

- struct
- struct literals

- JSON: basic types are numbers, bolleans, string
    json array and objects

    Converting a Go dat a structure like movies to JSON is called marshaling. `json.Marshal(movies)`, produce a byte slice. `json.Unmarshl`
- channel
    ch = make(chan int)
    send  `ch <- val`
    receive `val <- ch`
    can be closed `close(ch)`

Channels support a third operat ion, close, which sets a flag indic ating that no more values will
ever be sent on this channel; subsequent attempts to send will panic.Receive operat ions on a closed channel yield the values that have been sent until no more values are left; any receive
operations thereafter complete immediately and yield the zero value of the channel’s element type.

buffered an unbuffered channel
    ch = make(chan, int,  6)
unbuffered channel : synchronous channels
When a value is sent on an unbuffered channel, the receipt of the value happens before the reawakening of the sending goroutine.

pipeline: counter -> squarer --> printer
after a channel has been closed, any further send operations on it will panic. After the closed channel has been drained, that is, after the last sent element has been received, all subsquent receive operation will processd without blocking but will yield a zero value.

iterate over a channel

unidirectional channel types `chan<- int`, `<-chan int`

cap(ch) // channel capacity
len(ch) // current buffered size

send and receive operations on a nil channel block forever, a case in a select statement whose channel is nil is never selected.

slice channel map 需要用make 生成
range on arrays and slices provides both the index and value for each entry.
first index and then entry.
range on map provide the key and value, if only one is provided(key), it works too.
range over string, indexn and char

``` go
type Name struct {

}
type Name interface {

}
```

## function

- Variadic Functions

```go
func sum(vals ...int)int {
    total ：= 0
    for _, val : range vals {
        total +=val
    }
    return total
}
```

- Defer

Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.

- panic

- Recover

## Methods

traditional function
function with receiver
function with a pointer receiver

method value and method expression

## Interfaces

Concrete Typ | Interface Type
-|-
`map` `struct` `int` `string` `englishBot`| `bot`

- interface are not generic types
- interfaces are 'implicit'
- interfaces are a contract to help us manage types

## 7.5. Interface Values

The zero value for an interface has both its type and value components set to nil
Calling any met hod of a nil interface value causes a panic

The interface value’s dynamic type is set to the type descriptor for the pointer type *os.File, and its dynamic value holds a copy of os.Stdout, which is a pointer to the os.File variable representing the standard output of the process

call through an interface must use dynamic dispatch

interface|
-|
type descriptor|
value|
|

slice, maps, function type is not comparable
When handling errors, or during debugging, it is often helpful to report the dynamic type of an interface value. For that, we use the fmt package’s `%T `verb:

## Goroutine and Channels

## Concurrency with shared variables

- race condition
- mutual exclusion: sync.Mutex
- Read write mutexes: sync.RWMutex

multiple readers, single writer lock

A semaphore that counts only to 1 is called a `binary semaphore`.
a deferred Unlock will run even if the critical section panics.

## Testing

tests, benchmarks, and examples

Test Functions `*_test.go`, `*testing.T`
Coverarge
Profiling `*testing.B`

```go
go test bench=.
```

## reflection

## package and the go tool

import path
the package declearation
import declearation
blank imoport
Packages and Namings
The Go Tool

- Documenting Packages

```go
go doc time
go doc time.Since
go doc time.Duration.Seconds
godoc -http :8080
```

- Querying Packages

``` go
go list ...
go list -json hash
```

## Low-Level Programming

- unsafe.Sizeof,  Alignof, and Offsetof
- unsafe.Pointer

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
    channel need to be closed, otherwise there is deadlock

``` go
  for s:= range c{

    }
```

- timer

```go
    timer1 := time.NewTimer(2 * time.Second)
    <-timer1.C
    time2 := time.NewTimer(time.Second)
    go func() {
        <- timer2.C
    }()
    stop2 := timer2.Stop()
```

     One reason a timer may be useful is that you can cancel the timer before it expires.
- ticker

``` go
    ticker := time.NewTicker(500 * time.Millisecond)
    go func(){
        for t:= range ticker.C{
            fmt.Println("Tick at", t)
        }
    }()
    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    fmt.Println("Ticker stopped")
```

- worker pools

- sync/atomic

    atomic.AddUint64(&ops, 1)
    atomic.LoadUint64(&ops)

- mutex
- stateful goroutines
- sorting

- regular expression
- Json
- time/ time parsing
- url/url parsing
- sha1 hash
- base64 encoding
- reading/writing files
- line filters
- command line arguments
- environment variables
- executing processes
- signals os/signal syscall

  gracefully shutdown the server when receive a SIGTERM

```go
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    go func(){
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }
    fmt.Println("awaiting signal")
    <- done
    fmt.Println("exiting")
```

- exit

    os.Exit(3) //defer will not be run when using os.Exit

channel 只能用make 创建 c := make(chan int)
unbuffered channel c := make(chan init) c := make(chan int, 0)
buffered channel  c := make(chan int, 10)

goroutine 不同于线程， 一个线程可以包含数个goroutine, 如果线程被阻塞， 则使用其他线程， 如果没有线程， 则开辟新线程，
， 开辟后不删除， 即动态增加的线程

close channel twice `panic`
read from closed channel `ok`
write to close channel `panic`

value receiver and pointer receiver

    值類型的receiver 需要copy一份，不會修改原來的值
    指針類型的可以修改原來的值

The select statement is like a switch, but the decision is based on ability to communicate rather than equal values.

Some elements of Go step further from C, even C++ and Java:

    concurrency
    garbage collection
    interface types
    reflection
    type switch

resources-for-new-go-programmers: `https://dave.cheney.net/resources-for-new-go-programmers`

- concurrency is about dealing a lot of things at once
- parallelism is about doing a lot of things at once