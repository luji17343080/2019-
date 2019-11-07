package main

import (
	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	flag "github.com/spf13/pflag"
)

const (
	//设置默认端口8080
	PORT string = "8080"
)

func main() {
	//如果没有监听到端口，则使用默认端口8080
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	//端口号解析，用户使用-p设置端口号
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	app := iris.New() //创建app结构体对象
	app.Use(recover.New())
	app.Use(logger.New())
	// 输出html
	// 请求方式: GET
	// 访问地址: http://localhost:8080/welcome
	app.Handle("GET", "/welcome", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})
	// 输出字符串
	// 类似于 app.Handle("GET", "/ping", [...])
	// 请求方式: GET
	// 请求地址: http://localhost:port/cloudgo
	app.Get("/cloudgo", func(ctx iris.Context) {
		ctx.WriteString("简单web服务器cloudgo!")
	})
	// 输出json格式信息
	// 请求方式: GET
	// 请求地址: http://localhost:port/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})
	app.Run(iris.Addr(":" + port)) //启动服务器监听端口
}
