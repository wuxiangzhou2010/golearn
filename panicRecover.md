panic

1. 停止当前函数执行
2. 一直向上， 执行每一层的 defer
3. 如果没有遇见 recover， 程序退出

defer:

- 仅在 defer 调用中使用
- 获取 panic 的值
- 如果无法处理， 可重新 panic
