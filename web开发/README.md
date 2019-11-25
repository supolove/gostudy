瓶颈

磁盘IO瓶颈型

CPU计算瓶颈型

网络带宽瓶颈型

//epoll/kqueue之类的系统提供的IO多路复用接口  
// c10k C10k问题会导致计算机没有办法充分利用CPU来处理更多的用户连接

也就是说如今如果你的程序主要是和网络打交道，那么瓶颈一定在用户程序而不在操作系统内核

web吞吐量，接口的QPS

CDN服务

Proxy服务

流量限制 漏桶、令牌桶

1. 漏桶是指我们有一个一直装满了水的桶，每过固定的一段时间即向外漏一滴水。如果你接到了这滴水，那么你就可以继续服务请求，如果没有接到，那么就需要等待下一滴水。
2. 令牌桶则是指匀速向桶中添加令牌，服务请求时需要从桶中获取令牌，令牌的数目可以按照需要消耗的资源进行相应的调整。如果没有令牌，可以选择等待，或者放弃。

github.com/juju/ratelimit

QoS全称是Quality of Service 顾名思义是服务质量

## 接口和表驱动开发
1. 可以将系统中与业务本身流程无关的部分做拆解和异步化
2. 旁支流程的时延如若非常敏感，那么需要与主流程系统进行RPC通信
3. 如果时延不敏感，我们只要将下游需要的数据打包成一条消息，传入消息队列，之后的事情与主流程一概无关

### 使用函数封装业务流程  
最基本的封装过程，我们把相似的行为放在一起，然后打包成一个一个的函数，让自己杂乱无章的代码变成下面这个样子：
```go
func BusinessProcess(ctx context.Context, params Params) (resp, error){
    ValidateDistrict()    // 判断是否是地区限定商品
    ValidateVIPProduct()  // 检查是否是只提供给 vip 的商品
    GetUserInfo()         // 从用户系统获取更详细的用户信息
    GetProductDesc()      // 从商品系统中获取商品在该时间点的详细信息
    DecrementStorage()    // 扣减库存
    CreateOrderSnapshot() // 创建订单快照
    return CreateSuccess
}
```

### 使用接口来做抽象
如果我们正在做的是平台系统，需要由平台来定义统一的业务流程和业务规范，那么基于接口的抽象就是有意义的
```go
// OrderCreator 创建订单流程
type OrderCreator interface {
    ValidateDistrict()    // 判断是否是地区限定商品
    ValidateVIPProduct()  // 检查是否是只提供给 vip 的商品
    GetUserInfo()         // 从用户系统获取更详细的用户信息
    GetProductDesc()      // 从商品系统中获取商品在该时间点的详细信息
    DecrementStorage()    // 扣减库存
    CreateOrderSnapshot() // 创建订单快照
}
```



在大型系统中容错是重要的，能够让系统按百分比，分批次到达最终用户
灰度发布
1. 通过分批次部署实现灰度发布
2. 通过业务规则进行灰度发布

