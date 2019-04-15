## 8.1 testing

传统测试

- 测试数据和测试逻辑混在一起
- 出错信息不明确
- 一旦一个数据出错测试全部结束

表格驱动测试

- 分离测试数据和测试逻辑
- 明确的出错信息
- 可以部分失败

### test.T

## 8.2 代码覆盖率和性能测试

- 使用 IDE 查看代码覆盖
- 使用 go test 获取代码覆盖报告
- 使用 go tool cover 查看代码覆盖报告

### test.B

## 8.3 使用 pprof 进行性能调优

```sh
go test -bench . --cpuprofile cpu.out
go tool pprof cpu.out
web
```

b.ResetTimer() 重新计算 time
使用 map 可能会有性能瓶颈 mapaccess/ mapassign

## 8.6 生成文档和示例代码

```go
go doc
go doc Queue

go doc IsEmpty

godoc -http :6060
```

### 示例代码

```go
ExampleQueue_Pop()
//output:
```

reference:

- [8 6 生成文档和示例代码](https://www.youtube.com/watch?v=jUFbKCsVkAk&list=PL9avoKyUyEuy6neui9YrtBQtbbKx0y_y5&index=38)
