## RPC和Protobuf

RPC是远程过程调用的缩写（Remote Procedure Call）通俗地说就是调用远处的一个函数

### 4.1
RPC是远程过程调用的简称，是分布式系统中不同节点间流行的通信方式。在互联网时代，
RPC已经和IPC一样成为一个不可或缺的基础构件。因此Go语言的标准库也提供了一个简单的RPC实现，
我们将以此为入口学习RPC的各种用法。

### 4.1.2 更安全的RPC接口
在涉及RPC的应用中，作为开发人员一般至少有三种角色：
1. 首选是服务端实现RPC方法的开发人员，
2. 其次是客户端调用RPC方法的人员，
3. 最后也是最重要的是制定服务端和客户端RPC接口规范的设计人员。  
在前面的例子中我们为了简化将以上几种角色的工作全部放到了一起，虽然看似实现简单，但是不利于后期的维护和工作的切割。

我们将RPC服务的接口规范分为三个部分
1. 首先是服务的名字
2. 然后是服务要实现的详细方法列表
3. 最后是注册该类型服务的函数

### Protobuf
Protobuf是Protocol Buffers的简称，它是Google公司开发的一种数据描述语言，
并于2008年对外开源。Protobuf刚开源时的定位类似于XML、JSON等数据描述语言，
通过附带工具生成代码并实现将结构化数据序列化的功能。
但是我们更关注的是Protobuf作为接口规范的描述语言，可以作为设计安全的跨语言PRC接口的基础工具。

Protobuf核心的工具集是C++语言开发的，在官方的protoc编译器中并不支持Go语言。
要想基于上面的hello.proto文件生成相应的Go代码，需要安装相应的插件。
首先是安装官方的protoc工具，可以从 `https://github.com/google/protobuf/releases` 下载。
然后是安装针对Go语言的代码生成插件，可以通过go get github.com/golang/protobuf/protoc-gen-go命令安装。

然后通过以下命令生成相应的Go代码：
```sh
$ protoc --go_out=. hello.proto
```

protoc-gen-go内部已经集成了一个名字为grpc的插件，可以针对gRPC生成代码：
```sh
$ protoc --go_out=plugins=grpc:. hello.proto

```