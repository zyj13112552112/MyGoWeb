package goweb

import (
	"net/http"
)

//参照gin框架的接口

type HandlerFunc func(c *Context)

type Engine struct {
	router *router
}

func New()*Engine{
	return &Engine{router: newRouter()}
}

func (engine *Engine)addRouter(method string,pattern string,handler HandlerFunc){
	engine.router.addRouter(method,pattern,handler)
}

func (engine *Engine)Get(pattern string,handler HandlerFunc){
	engine.addRouter("Get",pattern,handler)
}

func (engine *Engine)Post(pattern string,handler HandlerFunc){
	engine.addRouter("Post",pattern,handler)
}

func (engine *Engine)Run(addr string)(err error){
	return http.ListenAndServe(addr,engine)
}

func (engine *Engine)ServeHTTP(w http.ResponseWriter,r *http.Request){
	c := newContext(w, r)
	engine.router.handle(c)
}


