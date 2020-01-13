## go中的interface [转载](http://kangkona.github.io/oo-in-golang/)

记得有位大学老师把Java中的接口比作资格证书，你要去考资格证书，并达到资格证书中的所有要求，才算具有某些资质。golang的接口则不太一样，只要你能做到资格证书中规定的那些事情，不管你去不去考这个资格证，都认为你具有了资质。其实这种想法在一些动态语言中实现过，且有诗为证：

如果一个人看起来像鸭子,走起来像鸭子,叫起来像鸭子,那么他就是个基佬。

这种非侵入式的设计方式，很大程度上也是为了解耦。接口和类本就是不同的东西：类是为了把数据和代码包装在一起，是为了对内实现；接口则更像是一种契约，是为了对外展示。基于这种抽象层面的接口进行编程，很容易达到设计模式中的依赖倒置，接口隔离以及迪米特法则几个原则。

编程语言的进化固然可以带来一些工程上的好处， 但千万不要忘了那句古训：

`There is no silver bullet.`

## 继承与组合 [转载](https://segmentfault.com/a/1190000021504638?utm_source=tag-newest)

__Go 并没有支持类继承体系和多态，Go 是面向对象却不是一般所理解的那种面向对象.__

```
package main

import (
  "fmt"
)

type CrossCompiler struct {
}

func (v CrossCompiler) crossCompile() {
  v.collectSource()
  v.compileToTarget()
}

func (v CrossCompiler) collectSource() {
  fmt.Println("CrossCompiler.collectSource")
}

func (v CrossCompiler) compileToTarget() {
  fmt.Println("CrossCompiler.compileToTarget")
}

type IPhoneCompiler struct {
  CrossCompiler
}

func (v IPhoneCompiler) collectSource() {
  fmt.Println("IPhoneCompiler.collectSource")
}

func (v IPhoneCompiler) compileToTarget() {
  fmt.Println("IPhoneCompiler.compileToTarget")
}

func main() {
  iPhone := IPhoneCompiler{}
  iPhone.crossCompile()
  iPhone.collectSource()
}
```

```
# Expect
IPhoneCompiler.collectSource
IPhoneCompiler.compileToTarget
IPhoneCompiler.collectSource

# Output
CrossCompiler.collectSource
CrossCompiler.compileToTarget
IPhoneCompiler.collectSource
```

OOAD中，除了继承，还有一种解决方法是组合。类继承的访问范围比组合要大，组合能够获取最小的信息。
优先使用组合而非继承，以上案例是Go的内嵌语法，内嵌属于组合。

## 面向接口编程

```
type SourceCollector interface {
  collectSource()
}

type TargetCompiler interface {
  compileToTarget()
}

type CrossCompiler struct {
  collector SourceCollector
  compiler  TargetCompiler
}

func (v CrossCompiler) crossCompile() {
  v.collector.collectSource()
  v.compiler.compileToTarget()
}

type IPhoneCompiler struct {
}

func (v IPhoneCompiler) collectSource() {
  fmt.Println("IPhoneCompiler.collectSource")
}

func (v IPhoneCompiler) compileToTarget() {
  fmt.Println("IPhoneCompiler.compileToTarget")
}

iPhone := IPhoneCompiler{}
compiler := CrossCompiler{iPhone, iPhone}
compiler.crossCompile()
```

接口应该定义最小的功能范围，以便重用。

__在 Go 中接口都比较小，非常小，只有一两个函数；但是对象却会比较大，会使用很多的接口。这种方式能够以最灵活的方式重用代码，而且保持接口的有效性和最小化，也就是接口隔离。__

隐式实现接口有个很好的作用，就是两个类似的模块实现同样的服务时，可以无缝的提供服务，甚至可以同时提供服务。

## 内嵌

__Embeding 只是将内嵌的数据和函数自动全部代理了一遍而已，本质上还是使用这个内嵌对象的服务。__

Outer 内嵌了Inner，和 Outer 继承 Inner 的区别在于：内嵌 Inner 是不知道自己被内嵌，调用 Inner 的函数，并不会对 Outer 有任何影响，Outer 内嵌 Inner 只是自动将 Inner 的数据和方法代理了一遍，但是本质上 Inner 的东西还不是 Outer 的东西；对于继承，调用 Inner 的函数有可能会改变 Outer 的数据，因为 Outer 继承 Inner，那么 Outer 就是 Inner，二者的依赖是更紧密的。

