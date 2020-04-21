# go 常用设计模式

23 种设计模式

- [单例模式](https://github.com/senghoo/golang-design-pattern/blob/master/03_singleton/singleton.go)

确保一个类只有一个实例，并提供该实例的全局访问点。

- [抽象工厂](https://github.com/senghoo/golang-design-pattern/blob/master/00_simple_factory/simple.go)
- [简单工厂](https://learnku.com/articles/33703)

提供一个接口用于创建相关对象的家族；

- [代理模式](https://github.com/senghoo/golang-design-pattern/blob/master/09_proxy/proxy.go)

  为其他对象提供一种代理以控制对这个对象的访问。

- [适配器模式](https://github.com/senghoo/golang-design-pattern/blob/master/02_adapter/adapter.go)

适配另一个不兼容的接口来一起工作；

- [装饰器模式](https://learnku.com/articles/33699)

装饰器模式动态地将责任附加到对象上。若要扩展功能，装饰者提供了比继承更有弹性的替代方案。

- [策略模式](https://learnku.com/articles/33698)

定义一系列的算法，把它们一个个封装起来，并且使它们可相互替换。

- [观察者模式](https://github.com/senghoo/golang-design-pattern/blob/master/10_observer/obserser.go)

对象间的一种一对多的依赖关系，以便一个对象的状态发生变化时，所有依赖于它的对象都得到通知并自动刷新；

- [责任链模式](https://learnku.com/articles/33708)

为某个请求创建一个对象链，每个对象依次检查此请求，并对其进行处理，或者将它传给链中的下一个对象

reference:

- [CS-note 设计模式 全](http://cyc2018.gitee.io/cs-notes/#/notes/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F)

- [go 设计模式 github](https://github.com/senghoo/golang-design-pattern)

- [golang 设计模式-以 kubernetes 源码为例
  ](https://juejin.im/post/5a113e686fb9a0452936596c)

- [超赞的 GO 语言设计模式和成例集锦](https://studygolang.com/articles/8230)
