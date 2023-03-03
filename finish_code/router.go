package goweb

import (
	"net/http"
	"strings"
)

type router struct {
	roots map[string]*node          //使用 roots 来存储每种请求方式的Trie树根节点
	handlers map[string]HandlerFunc //handlers 存储每种请求方式的 HandlerFunc
}

func newRouter()*router{
	return &router{
		roots: map[string]*node{},
		handlers: map[string]HandlerFunc{},
	}
}



func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.method, c.path)

	if n != nil {
		key := c.method + "-" + n.pattern
		c.Params = params
		c.handler = append(c.handler, r.handlers[key])
	} else {
		c.handler = append(c.handler, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.path)
		})
	}
	c.Next()
}
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

//获取路由参数
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *router)addRoute(method string,pattern string,handler HandlerFunc){
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}