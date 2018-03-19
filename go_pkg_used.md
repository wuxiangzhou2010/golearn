## go package met

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

- sync

    sync.Mutex
    r := sync.RWMutex
    r.RLock
    r.RuLock
    r.Lock

- context

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

- Bytes

    bytes.Equal()

- time

    time.Duration
    time.NewTimer()