package goweb

import (
	"net/http"
	"strings"
)

//参照gin框架的接口

type HandlerFunc func(c *Context)

type Engine struct {
	router *router
	*RouterGroup
	groups []*RouterGroup
}

func New()*Engine{
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine:engine,}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (engine *Engine)addRouter(method string,pattern string,handler HandlerFunc){
	engine.router.addRoute(method,pattern,handler)
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
	var middler []HandlerFunc
	for _,group:=range engine.groups{
		if strings.HasPrefix(r.URL.Path,group.prefix){
			middler = append(middler, group.middlerwares...)
		}
	}
	c := newContext(w, r)
	c.handler = middler
	engine.router.handle(c)

}


