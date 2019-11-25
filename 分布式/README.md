#分布式

## 分布式id生成器

snowflake算法


| unused | time in milliseconds | datacenter_id | worker_id | sequence_id |
| :---: | :---: | :---: | :---: | :---: |
| 0 | 0000 0000 0000 0000 0000 0000 0000 0000 0000 00000 0 | 00000 | 00000 | 0000 0000 0000 |
| 1 bit | 41 bit | 5bit | 5bit | 12bit |
| 符号位 | 毫秒时间戳 | 数据中心的id | 机器的实例id | 循环自增id |

说明

| 数据段               | 说明                                 |
| -------------------- | ------------------------------------ |
| time in milliseconds | 运行期生成的                         |
| datacenter_id        | 部署阶段就能够获取得到，通过api获取  |
| worker_id            | 部署阶段就能够获取得到，开始工具生成 |
| sequence_id          | 运行期生成的                         |

```go
"github.com/bwmarrin/snowflake" // 标准的实现
"github.com/sony/sonyflake" // sony实现
```

## 分布式锁

1. 进程内加锁

2. trylock

3. 基于Redis的setnx

4. 基于ZooKeeper
5. 基于etcd

## 延时任务系统

1. 实现一套类似crontab的分布式定时任务管理系统。
2. 实现一个支持定时发送消息的消息队列。

时间堆

任务分发

## 分布式搜索引擎

搜索引擎

查询dsl