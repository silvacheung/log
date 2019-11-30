# log
一个简单快速的日志库

# 使用方式
go get -u github.com/marcosxz/log

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

## 接管标准库或者其他日志库(以标准库为例)
```
log：代表标准库
mylog：代表我们的库

log.SetOutput(mylog.Writer())
```

## 一些简单的测试
[点击查看](./logger_test.go)