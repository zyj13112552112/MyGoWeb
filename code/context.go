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
	//response info
	statusCode int
}
func newContext(w http.ResponseWriter,r *http.Request)*Context{
	return &Context{
		w:          w,
		r:          r,
		path:       r.URL.Path,
		method:     r.Method,
	}
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
func (c *Context)string(code int,format string,values...any){
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

