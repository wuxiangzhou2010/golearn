# go package met

- path/filepath

  filepath.Join()
  filepath.Dir() //Dir returns all but the last element of path, typically the path's directory.

- os

  v, found := os.LookupEnv(f.AltName); found
  os.Executable() //Executable returns the path name for the executable that started the current process

- strings

  strings.ToUpper(strings.TrimSpace())
  strings.Replace
  strings.ToLower
  strings.LastIndexByte(filename, '.') //turns the index of the last instance of c in s, or -1 if c is not present in s.

- strconv

  strconv.ParseInt
  strconv.ParseUint

- net

  net.DialTCP
  net.DialUDP
  net.DialUnix
  net.Listen
  net.ListTCP
  net.ListenUDP
  net.LoopupIp

  net.FileConnc //FileConn returns a copy of the network connection corresponding to the open file f
  net.ParseIP(s string) //ParseIP parses s as an IP address,
  net.SplitHostPort
  net.Addr
  net.Conn
  net.CIDRMask
  net.TCPAddr
  net.UDPAddr
  net.TCPConn
  net.UDPConn
  net.UnixAddr
  net.UnixConn
  net.IP
  net.IPMask
  net.IPNet

  DNS resolve can use buildin go or system, buildin is using a goroutine while c is using a block call(thread).

  it seems internal go routine is more efficient.

- sync

  sync.Mutex
  r := sync.RWMutex
  r.RLock
  r.RuLock
  r.Lock
  atomic
  // 原子增减 atomic.AddInt32(&value, n)
  // 原子读取 atomic.LoadInt32(&value)
  // 原子写入 atomic.StoreInt32(&value, n)
  // 原子交换 swap
  // 原子比较交换 compare and swap

- sync.pool

  reference:

- [《GO 并发编程实战》—— 临时对象池](http://ifeve.com/go-concurrency-object-pool/)

- context

Package contex defines the Contex type, which carries deadlines, cancelation signals, and other request-scoped values across API boundaries and between processes.

    c := context.Context
    c.Value()
    c.WithValue()

- io

  io.Reader

- math/rand

  rand.Uint32()

- reflect

  reflect.TypeOf()
  reflect.Type

- encoding/hex

  hex.EncodeToString() //returns hexadecimal encoding of src
  hex.Decode()/Encode()

- encoding/json
- regexp

- Bytes

  bytes.Equal()
  bytes.Split()
  bytes.SplitN()
  bytes.Trim()

- time

  time.Duration
  time.NewTimer()

- bufio

- log

- container

  - heap

  - list
    - Package list implements a doubly linked list.
  - ring
    - Package ring implements operations on circular lists
