package goweb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]any

type Context struct {
	//origin object
	w http.ResponseWriter
	r *http.Request
	//request info
	path string
	method string
	Params map[string]string       //存储解析后的参数
	//response info
	statusCode int
	//middlewares
	handler []HandlerFunc
	index int
}
func newContext(w http.ResponseWriter,r *http.Request)*Context{
	return &Context{
		w:          w,
		r:          r,
		path:       r.URL.Path,
		method:     r.Method,
		index:		-1,
	}
}
func (c *Context)Next(){
	c.index++
	s := len(c.handler)
	for ;c.index<s;c.index++{
		c.handler[c.index](c)
	}
}
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}
func (c *Context)PostForm(key string)string{
	return c.r.FormValue(key)
}
func (c *Context)Query(key string)string{
	return c.r.URL.Query().Get(key)
}
func (c *Context)Status(code int){
	c.statusCode = code
	c.w.WriteHeader(code)
}
func (c *Context)SetHeader(key string,value string){
	c.w.Header().Set(key,value)
}
func (c *Context)String(code int,format string,values...any){
	c.SetHeader("Content-Type","text/plain")
	c.Status(code)
	c.w.Write([]byte(fmt.Sprintf(format,values...)))
}
func (c *Context)JSON(code int,obj interface{}){
	c.SetHeader("Content-Type","application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.w)
	if err:=encoder.Encode(obj);err!=nil{
		http.Error(c.w,err.Error(),http.StatusInternalServerError)
	}
}
func (c *Context)HTML(code int,html string){
	c.SetHeader("Content-Type","text/html")
	c.Status(code)
	c.w.Write([]byte(html))
}
func (c *Context)Data(code int,data []byte){
	c.Status(code)
	c.w.Write(data)
}
//Header.set()->writerHeader()->writer()
