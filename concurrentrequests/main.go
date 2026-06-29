package main

import "sync"

// URLs is a fixed list of 50 public URLs to fetch concurrently
// with bounded concurrency (no more than N in flight at once).
var URLs = []string{
	"https://httpbin.org/get",
	"https://httpbin.org/delay/1",
	"https://httpbin.org/delay/2",
	"https://httpbin.org/uuid",
	"https://httpbin.org/ip",
	"https://httpbin.org/user-agent",
	"https://httpbin.org/headers",
	"https://httpbin.org/status/200",
	"https://httpbin.org/status/404",
	"https://httpbin.org/json",
	"https://api.github.com",
	"https://api.github.com/zen",
	"https://api.github.com/users/golang",
	"https://api.github.com/users/torvalds",
	"https://api.github.com/repos/golang/go",
	"https://jsonplaceholder.typicode.com/posts/1",
	"https://jsonplaceholder.typicode.com/posts/2",
	"https://jsonplaceholder.typicode.com/posts/3",
	"https://jsonplaceholder.typicode.com/comments/1",
	"https://jsonplaceholder.typicode.com/users/1",
	"https://jsonplaceholder.typicode.com/todos/1",
	"https://jsonplaceholder.typicode.com/albums/1",
	"https://jsonplaceholder.typicode.com/photos/1",
	"https://example.com",
	"https://example.org",
	"https://example.net",
	"https://www.google.com",
	"https://www.wikipedia.org",
	"https://en.wikipedia.org/wiki/Go_(programming_language)",
	"https://en.wikipedia.org/wiki/Concurrency_(computer_science)",
	"https://go.dev",
	"https://go.dev/doc/",
	"https://pkg.go.dev/net/http",
	"https://pkg.go.dev/context",
	"https://pkg.go.dev/sync",
	"https://news.ycombinator.com",
	"https://www.reddit.com/r/golang.json",
	"https://httpstat.us/200",
	"https://httpstat.us/301",
	"https://httpstat.us/500",
	"https://dog.ceo/api/breeds/list/all",
	"https://catfact.ninja/fact",
	"https://api.coindesk.com/v1/bpi/currentprice.json",
	"https://api.publicapis.org/entries",
	"https://reqres.in/api/users/1",
	"https://reqres.in/api/users/2",
	"https://reqres.in/api/users?page=1",
	"https://postman-echo.com/get",
	"https://postman-echo.com/headers",
	"https://postman-echo.com/time/now",
}

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main(){
	for _, url := range(URLs){
		wg.Add(1)
		go process(url)
	}
	wg.Wait()
}