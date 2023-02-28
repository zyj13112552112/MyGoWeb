package goweb

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter()*router{
	return &router{
		handlers: map[string]HandlerFunc{},
	}
}

func (r *router)addRouter(method string,pattern string,handler HandlerFunc){
	log.Printf("Router %4s - %s",method,pattern)
	key := method + "-" + pattern
	r.handlers[key]=handler
}

func (r *router)handle(c *Context){
	key := c.method + "-" + c.path
	if handler, ok := r.handlers[key]; ok{
		handler(c)
	}else{
		c.string(http.StatusNotFound,"404 not found : %s\n",c.path)
	}

}
