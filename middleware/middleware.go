package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func RequestInfos() gin.HandlerFunc {
	return func(context *gin.Context) {

		startTime := time.Now()

		addr := context.Request.RemoteAddr
		uri := context.Request.RequestURI
		method := context.Request.Method
		fullPath := context.FullPath()
		fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
		fmt.Println("请求地址：", addr)
		fmt.Println("请求URI：", uri)
		fmt.Println("请求方法：", method)
		fmt.Println("请求FullPath：", fullPath)

		fmt.Println("------------------------------------------")

		context.Next() // 调用后续函数
		//context.Abort() // 阻止调用后续函数

		fmt.Println("-----------------------------------------")

		//response := context.Request.Response
		//status := response.Status
		//code := response.StatusCode
		//fmt.Println("响应状态：", status)
		//fmt.Println("响应状态码：", code)
		cost := time.Since(startTime)
		fmt.Println("耗时：", cost)
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	}
}

func Test() {
	engine := gin.Default()
	engine.Use(RequestInfos())
	engine.Handle("GET", "/hello", func(context *gin.Context) {

		name, _ := context.GetQuery("name")

		fmt.Println("Hello")

		_, _ = context.Writer.Write([]byte("Hello " + name))
	})
	err := engine.Run()

	if err != nil {
		log.Fatal(err.Error())
	}
}
