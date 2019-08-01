# log
一个简单快速的日志库
# 使用方式
go get -u github.com/marcosxzhang/log
# 依赖
只依赖一个外部的快速json库: github.com/json-iterator/go
# 例子
## 设置参数
```
log.SetOptions(
    log.WithLevel(TraceLevel), // 设置日志级别
    log.WithStdLevel(TraceLevel), // 如果接管其他库log,设置输出级别
    log.WithOutput(file), // 输出目标
    log.WithFileLine(true), // 是否启用行列信息
    log.WithErrorHandler(nil), // 定义一个错误处理器
    log.WithNoLock(true), // 设置无锁模式,追加写模式下可以设置无锁
    log.WithFormatter(&TextFormatter{IgnoreBasicFields: false}), // 设置日志输出的格式，默认提供了text和json两种,可以自己实现Formatter接口来自定义格式,IgnoreBasicFields字段可以忽略输出基础字段
)
```
## 基本使用
```
log.Info(args ...)
```

## 支持自定义键值对
```
log.InfoKvln("key1,key2,key3", value1, value2, value3)
输出(text format)：
... key1=value1 key2=value2 key3=value3 ...
输出(json format)：
{..., "key1":value1, "key2":value2, "key3":value3, ...}
```

## 接管标准库或者其他日志库(以标准库为例)
```
log：代表标准库
mylog：代表我们的库
log.SetOutput(mylog.Writer())
```
## 一些简单的测试(在Windows下进行的,需要验证的话可以自己测试一下)
[点击查看](./logger_test.go)

## 如果有一些没有注意到的问题，欢迎指正