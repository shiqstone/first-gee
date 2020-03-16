从零实现Web框架

这是参考[7天用Go从零实现Web框架Gee教程](https://geektutu.com/post/gee.html)，用来学习Golang Web框架开发的实验代码
未完成，更新中...

# http.Handle
## 标准库启动Web服务
net/http
http.HandleFunc
func demoHandler(w http.ResponseWriter, req *http.Request) {
	...
}
http.ListenAndServe(":9999", nil)

## 实现http.Handle接口
type Engine struct{}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    switch req.URL.Path {
        case 
        default
    }
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}

## 实现框架雏形
### 首先定义了类型HandlerFunc
这是提供给框架用户的，用来定义路由映射的处理方法。我们在Engine中，添加了一张路由映射表router，key 由请求方法和静态路由地址构成，例如GET-/、GET-/hello、POST-/hello，这样针对相同的路由，如果请求方法不同,可以映射不同的处理方法(Handler)，value 是用户映射的处理方法。
### 路由映射表
当用户调用(*Engine).GET()方法时，会将路由和处理方法注册到映射表 router 中，(*Engine).Run()方法，是 ListenAndServe 的包装。
### Engine实现ServeHTTP
Engine实现的 ServeHTTP 方法的作用就是，解析请求的路径，查找路由映射表，如果查到，就执行注册的处理方法。如果查不到，就返回 404 NOT FOUND 。

# 上下文Context
将路由(router)独立出来，方便之后增强
设计上下文(Context)，封装 Request 和 Response
## Router
## Context 
## Engine
## main.go

# 使用 Tire 树实现动态路由(dynamic route)

# 分组控制Group

# 中间件
中间件(middlewares)，简单说，就是非业务的技术类组件
- 插入点在哪?
- 中间件的输入是什么？