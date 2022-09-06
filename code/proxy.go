package main

import "fmt"

/*
代理是一种结构型设计模式， 让你能提供真实服务对象的替代品给客户端使用。 代理接收客户端的请求并进行一些处理 （访问控制和缓存等）， 然后再将请求传递给服务对象。
代理对象拥有和服务对象相同的接口， 这使得当其被传递给客户端时可与真实对象互换。
*/
// 主体
type Server interface {
	handRequest(method string, url string) (status int, resp string)
}

//代理
type Proxy struct {
	app             *Application
	maxAllowRequest int
	ratelimit       map[string]int
}

func NewProxy() *Proxy {
	return &Proxy{
		app:             &Application{},
		maxAllowRequest: 2,
		ratelimit:       make(map[string]int),
	}
}

func (p *Proxy) handRequest(method string, url string) (status int, resp string) {
	if ok := p.check(url); !ok {
		return 403, "not allowed"
	}
	return p.app.handRequest(method, url)
}

func (p *Proxy) check(url string) bool {
	if p.ratelimit[url] == 0 {
		p.ratelimit[url] = 1
	}
	if p.ratelimit[url] > p.maxAllowRequest {
		return false
	}
	p.ratelimit[url] = p.ratelimit[url] + 1
	return true
}

//真实主体
type Application struct {
}

func (a *Application) handRequest(method string, url string) (statuscode int, resp string) {
	if url == "/app/status" && method == "GET" {
		return 200, "ok"
	}
	if url == "/create/user" && method == "POST" {
		return 201, "User Created"
	}
	return 404, "Not Found"
}

func main() {
	n := NewProxy()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"
	statuscode, resp := n.handRequest("GET", appStatusURL)
	fmt.Printf("\nRequest URL: %s\nStatus Code: %d\nResponse Body: %s\n", appStatusURL, statuscode, resp)

	statuscode, resp = n.handRequest("GET", appStatusURL)
	fmt.Printf("\nRequest URL: %s\nStatus Code: %d\nResponse Body: %s\n", appStatusURL, statuscode, resp)

	statuscode, resp = n.handRequest("GET", appStatusURL)
	fmt.Printf("\nRequest URL: %s\nStatus Code: %d\nResponse Body: %s\n", appStatusURL, statuscode, resp)

	statuscode, resp = n.handRequest("POST", createuserURL)
	fmt.Printf("\nRequest URL: %s\nStatus Code: %d\nResponse Body: %s\n", appStatusURL, statuscode, resp)
}
