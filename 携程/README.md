# Swoole协程与Go协程的区别

来自：https://zhuanlan.zhihu.com/p/93905248

## 进程、线程、协程的概念

**进程是什么？**

进程就是应用程序的启动实例。
例如：打开一个软件，就是开启了一个进程。
进程拥有代码和打开的文件资源，数据资源，独立的内存空间。

**线程是什么？**

线程属于进程，是程序的执行者。
一个进程至少包含一个主线程，也可以有更多的子线程。
线程有两种调度策略，一是：分时调度，二是：抢占式调度。

**协程是什么？**

协程是轻量级线程, 协程的创建、切换、挂起、销毁全部为内存操作，消耗是非常低的。
协程是属于线程，协程是在线程里执行的。
协程的调度是用户手动切换的，所以又叫用户空间线程。
协程的调度策略是：协作式调度。

## **Swoole 协程**

Swoole 的协程客户端必须在协程的上下文环境中使用。

```php
// 第一种情况：Request 回调本身是协程环境
$server->on('Request', function($request, $response) {
    // 创建 Mysql 协程客户端
    $mysql = new Swoole\Coroutine\MySQL();
    $mysql->connect([]);
    $mysql->query();
});

// 第二种情况：WorkerStart 回调不是协程环境
$server->on('WorkerStart', function() {
    // 需要先声明一个协程环境，才能使用协程客户端
    go(function(){
        // 创建 Mysql 协程客户端
        $mysql = new Swoole\Coroutine\MySQL();
        $mysql->connect([]);
        $mysql->query();
    });
});
```

Swoole 的协程是基于单线程的, 无法利用多核CPU，同一时间只有一个在调度。

```php
// 启动 4 个协程
$n = 4;
for ($i = 0; $i < $n; $i++) {
    go(function () use ($i) {
        // 模拟 IO 等待
        Co::sleep(1);
        echo microtime(true) . ": hello $i " . PHP_EOL;
    });
};

echo "hello main \n";

// 每次输出的结果都是一样
$ php test.php 
hello main 
1558749158.0913: hello 0 
1558749158.0915: hello 3 
1558749158.0915: hello 2 
1558749158.0915: hello 1
```

Swoole 协程使用示例及详解

```php
// 创建一个 Http 服务
$server = new Swoole\Http\Server('127.0.0.1', 9501, SWOOLE_BASE);

// 调用 onRequest 事件回调函数时，底层会调用 C 函数 coro_create 创建一个协程，
// 同时保存这个时间点的 CPU 寄存器状态和 ZendVM stack 信息。
$server->on('Request', function($request, $response) {
    // 创建一个 Mysql 的协程客户端
    $mysql = new Swoole\Coroutine\MySQL();

    // 调用 mysql->connect 时发生 IO 操作，底层会调用 C 函数 coro_save 保存当前协程的状态，
    // 包括 Zend VM 上下文以及协程描述的信息，并调用 coro_yield 让出程序控制权，当前的请求会挂起。
    // 当协程让出控制权之后，会继续进入 EventLoop 处理其他事件，这时 Swoole 会继续去处理其他客户端发来的 Request。
    $res = $mysql->connect([
        'host'     => '127.0.0.1',
        'user'     => 'root',
        'password' => 'root',
        'database' => 'test'
    ]);

    // IO 事件完成后，MySQL 连接成功或失败，底层调用 C 函数 coro_resume 恢复对应的协程，恢复 ZendVM 上下文，继续向下执行 PHP 代码。
    if ($res == false) {
        $response->end("MySQL connect fail");
        return;
    }

    // mysql->query 的执行过程和 mysql->connect 一致，也会进行一次协程切换调度
    $ret = $mysql->query('show tables', 2);

    // 所有操作完成后，调用 end 方法返回结果，并销毁此协程。
    $response->end('swoole response is ok, result='.var_export($ret, true));
});

// 启动服务
$server->start();
```

## **Go 的协程 goroutine**

goroutine 是轻量级的线程，Go 语言从语言层面就支持原生协程。
Go 协程与线程相比，开销非常小。
Go 协程的堆栈开销只用2KB，它可以根据程序的需要增大和缩小，
而线程必须指定堆栈的大小，并且堆栈的大小都是固定的。

goroutine 是通过 GPM 调度模型实现的。
M: 表示内核级线程，一个 M 就是一个线程，goroutine 跑在 M 之上的。
G: 表示一个 goroutine，它有自己的栈。
P: 全称是 Processor，处理器。它主要用来执行 goroutine 的，同时它也维护了一个 goroutine 队列。

Go 在 runtime、系统调用等多个方面对 goroutine 调度进行了封装和处理，当遇到长时间执行或进行系统调用时，
会主动把当前协程的 CPU 转让出去，让其他协程调度执行。



Go 语言原生层面就支持协层，不需要声明协程环境。

```go
package main

import "fmt"

func main() {
    // 直接通过 Go 关键字，就可以启动一个协程。
    go func() {
        fmt.Println("Hello Go!")
    }()
}
```

Go 协程是基于多线程的，可以利用多核 CPU，同一时间可能会有多个协程在执行。

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 设置这个参数，可以模拟单线程与 Swoole 的协程做比较
    // 如果这个参数设置成 1，则每次输出的结果都一样。
    // runtime.GOMAXPROCS(1)

    // 启动 4 个协程
    var i int64
    for i = 0; i < 4; i++ {
        go func(i int64) {
            // 模拟 IO 等待
            time.Sleep(1 * time.Second)
            fmt.Printf("hello %d \n", i)
        }(i)
    }

    fmt.Println("hello main")

    // 等待其他的协程执行完，如果不等待的话，
    // main 执行完退出后，其他的协程也会相继退出。
    time.Sleep(10 * time.Second)
}

// 第一次输出的结果
$ go run test.go
hello main
hello 2 
hello 1 
hello 0 
hello 3 

// 第二次输出的结果
$ go run test.go
hello main
hello 2 
hello 0 
hello 3 
hello 1 

// 依次类推，每次输出的结果都不一样
```

go 协程使用示例及详解

```go
package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    "net/http"
    "time"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
    dsn := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
        "root",
        "root",
        "127.0.0.1",
        "3306",
        "fastadmin",
    )
    db, err := gorm.Open("mysql", dsn)
    if err != nil {
        fmt.Printf("mysql connection failure, error: (%v)", err.Error())
        return
    }
    db.DB().SetMaxIdleConns(10)  // 设置连接池
    db.DB().SetMaxOpenConns(100) // 设置与数据库建立连接的最大数目
    db.DB().SetConnMaxLifetime(time.Second * 7)

    http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
        // http Request 是在协程中处理的
        // 在 Go 源码 src/net/http/server.go:2851 行处 `go c.serve(ctx)` 给每个请求启动了一个协程
        var name string
        row := db.Table("fa_auth_rule").Where("id = ?", 1).Select("name").Row()
        err = row.Scan(&name)
        if err != nil {
            fmt.Printf("error: %v", err)
            return
        }
        fmt.Printf("name: %v \n", name)
    })
    http.ListenAndServe("0.0.0.0:8001", nil)
}

```

## **总结**

- 协程是轻量级的线程，开销很小。
- Swoole 的协程客户端需要在协程的上下文环境中使用。
- 在 Swoole v4.3.2 版本之后，已经支持协程 CPU 密集场景调度。
- Go 语言层面就已经完全支持协程了。