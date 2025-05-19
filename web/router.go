package web

import (
	"fmt"
	"strings"
)

// 路由树实现
type router struct {
	trees map[string]*node
}
type node struct {
	children map[string]*node
	path     string
	handler  HandleFunc
}

func newRouter() router {
	return router{
		trees: map[string]*node{},
	}
}
func (r *router) addRoute(method, path string, handleFunc HandleFunc) {
	if path == "" {
		panic("path can't be empty")
	}
	root, ok := r.trees[method]
	if !ok {
		root = &node{
			path: "/",
		}
		r.trees[method] = root
	}
	if path[0] != '/' {
		panic("path must begin with '/'")
	}
	if path != "/" && path[len(path)-1] == '/' {
		panic("path end cant with '/'")
	}
	if path == "/" {
		if root.children != nil {
			panic("path 重复注册")
		}

		root.handler = handleFunc
		return
	}
	segs := strings.Split(path[1:], "/")
	for _, seg := range segs {
		if seg == "" {
			panic("path  can't be continuous")
		}
		child := root.childOrCreate(seg)
		root = child
	}
	if root.handler != nil {
		panic(fmt.Sprintf("web:路由冲突，重复注册[%s]", path))
	}
	root.handler = handleFunc
}
func (r *router) findRoute(method, path string) (*node, bool) {
	root, ok := r.trees[method]
	if !ok {
		return nil, false
	}
	if path == "/" {
		return root, true
	}
	path = strings.Trim(path, "/")
	segs := strings.Split(path, "/")
	for _, seg := range segs {
		child, found := root.childOf(seg)
		if !found {
			return nil, false
		}
		root = child
	}
	return root, true
}

func (n *node) childOrCreate(seg string) *node {
	if n.children == nil {
		n.children = map[string]*node{}
	}
	res, ok := n.children[seg]
	if !ok {
		res = &node{
			path: seg,
		}
		n.children[seg] = res
	}
	return res
}

func (n *node) childOf(path string) (*node, bool) {
	if n.children == nil {
		return nil, false
	}
	child, ok := n.children[path]
	return child, ok
}
