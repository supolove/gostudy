# 缓存
## redis设置
mac brew安装环境

| 文件         | 路径                              |      |
| ------------ | --------------------------------- | ---- |
| redis-cli    | /usr/local/Cellar/redis/3.2.6/bin |      |
| redis-server | /usr/local/Cellar/redis/3.2.6/bin |      |
| redis.conf   | /usr/local/etc/redis.conf         |      |

redis修改临时密码

```shell
redis-cli
# 查询是否有设置密码
127.0.0.1:6379> config get requirepass
# 设置密码 123456
127.0.0.1:6379> config set requirepass 123456
```

redis设置永久密码

修改redis.conf

```
requirepass foobared
requirepass 123456   指定密码123456
```

使用密码重新登录redis

```
redis-cli
127.0.0.1:6379> auth 123456
```



## redis说明

参考：https://www.runoob.com/redis/redis-tutorial.html

REmote DIctionary Server(Redis) 是一个由Salvatore Sanfilippo写的key-value存储系统。

Redis是一个开源的使用ANSI C语言编写、遵守BSD协议、支持网络、可基于内存亦可持久化的日志型、Key-Value数据库，并提供多种语言的API。

它通常被称为数据结构服务器，因为值（value）可以是 字符串(String), 哈希(Hash), 列表(list), 集合(sets) 和 有序集合(sorted sets)等类型。

数据结构

1. String：字符串
2. Hash：散列
3. List：列表
4. Set：集合
5. Sorted Set：有序集合

## go-redis

使用仓库

```shell
go get -u github.com/go-redis/redis/v7
```

如果仓库下载出问题，换个代理试一下

```shell
echo "export GOPROXY=https://goproxy.cn" >> ~/.profile && source ~/.profile
```

## redis按例

### 排行榜

使用有序结合可以实现排行榜，通过score排序

### 计数器

通过`INCR`可以对key做累加

### 消息推送

使用list可以做消息推送，服务器监控list是否有数据，有数据建推给订阅者

### 好友关注，粉丝

关注和粉丝可以用set结合

