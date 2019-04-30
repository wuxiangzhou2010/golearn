# Go Concurrency patterns

- [Google I/O 2012 - Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)

- Generator: functions that returns a channel

  - channels as a handle on a service
  - multiplexing --> Fan-in
  - restore sequencing.
  - select
  - Timeout using select `time.After(1 * time.Second)`
  - Timeout for whole conversion using select
  - Quit channel
  - Daisy-chain
  - system software

- [Rob Pike - 'Concurrency Is Not Parallelism'](https://www.youtube.com/-watch?v=cN_DpYBzKso&t=3s)
