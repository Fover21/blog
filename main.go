package main

import (
	"fmt"
	"blog/app/blog"
	"blog/app/shop"
	"blog/app/login"
	"blog/app/appall"
	"blog/routers"
)

func main() {
	// 加载多个APP的路由配置
	routers.Include(shop.Routers, blog.Routers, login.Routers, appall.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(":8080"); err != nil {
	fmt.Println("startup service failed, err:%v\n", err)
	}
}