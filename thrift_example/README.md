# thrift example

最近看的部分开源代码有使用到thrift，很久之前在刚工作的时候有个课题是用grpc的，
依稀记得grpc需要定义一个通信协议，包括通信结构体、通信方法等，自动生成pb等文件。
今天尝试了一下简单的thrift，发现其整体流程跟grpc差不多，依稀记得好像是thrift需要额外定义实现方法，
而grpc可以更加方便地直接使用。

```shell script
go get github.com/apache/thrift/lib/go/thrift
brew install thrift
thrift -r --gen go:package_prefix=github.com/salmon7/go-learning/ -o thrift_example thrift_example/echo.thrift
```
代码参考：  
[apache thrift golang](https://github.com/apache/thrift/tree/master/lib/go)  
[golang 网络框架之 thrift](https://www.jianshu.com/p/2fabaf897e38)

