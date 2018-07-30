package main

import (
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
)

func main() {
	//configor.Load(&Config, "config.yml")
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("哈喽世界！")
	})
	s.SetPort(8000)
	s.Run()

}

