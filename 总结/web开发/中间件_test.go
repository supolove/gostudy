package web开发

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

/** 基本实现方法*/
func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}

// 中间件
func timeMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)

		fmt.Println(timeElapsed)
	})
}

func TestMiddleware(t *testing.T) {
	http.Handle("/", timeMiddleware(http.HandlerFunc(hello)))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

/** 更优雅的中间件框架*/

// 中间件类型
type middleware func(handler http.Handler) http.Handler

// 定义路由器
type Router struct {
	middlewareChain []middleware
	mux             map[string]http.Handler
}

// 构建路由器
func NewRouter() *Router {
	return &Router{}
}

// 中间件方法链添加中间件
func (r *Router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *Router) Add(route string, h http.Handler) {
	var mergedHandler = h

	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = r.middlewareChain[i](mergedHandler)
	}

	r.mux[route] = mergedHandler
}

func TestMiddleware2(t *testing.T) {
	//r := NewRouter()
	//r.Use(logger)
	//r.Use(timeout)
	//r.Use(ratelimit)
	//r.Add("/", helloHandler)

}

const TravelBusiness = 1
const MarketBusiness = 2

type BusinessInstance interface {
	ValidateLogin()
	ValidateParams()
	AntispamCheck()
	GetPrice()
	CreateOrder()
	UpdateUserStatus()
	NotifyDownstreamSystems()
}

var businessInstanceMap = map[int]BusinessInstance{
	TravelBusiness: travelorder.New(),
	MarketBusiness: marketorder.New(),
}

func entry() {
	bi := businessInstanceMap[businessType]
}

/*
哪些事情适合在中间件中做
以较流行的开源Go语言框架chi为例：
compress.go
  => 对http的响应体进行压缩处理
heartbeat.go
  => 设置一个特殊的路由，例如/ping，/healthcheck，用来给负载均衡一类的前置服务进行探活
logger.go
  => 打印请求处理处理日志，例如请求处理时间，请求路由
profiler.go
  => 挂载pprof需要的路由，如`/pprof`、`/pprof/trace`到系统中
realip.go
  => 从请求头中读取X-Forwarded-For和X-Real-IP，将http.Request中的RemoteAddr修改为得到的RealIP
requestid.go
  => 为本次请求生成单独的requestid，可一路透传，用来生成分布式调用链路，也可用于在日志中串连单次请求的所有逻辑
timeout.go
  => 用context.Timeout设置超时时间，并将其通过http.Request一路透传下去
throttler.go
  => 通过定长大小的channel存储token，并通过这些token对接口进行限流
*/
