package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Byte() {
	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("请求路径：" + context.FullPath())

		_, _ = context.Writer.Write([]byte("响应字节"))
	})
	err := engine.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Json() {
	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("请求路径：" + context.FullPath())

		context.JSON(http.StatusOK, gin.H{
			"code": 1000,
			"msg":  "Success",
		})

		var responseEntity ResponseEntity
		responseEntity.Code = 10000
		responseEntity.Success = true
		responseEntity.Message = "Success"
		responseEntity.Data = nil
		context.JSON(http.StatusOK, &responseEntity)
	})
	err := engine.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Html() {
	engine := gin.Default()
	//engine.SetHTMLTemplate()
	engine.LoadHTMLGlob("./html/*")
	//engine.LoadHTMLFiles()
	engine.Static("/static", "./static")
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("请求路径：" + context.FullPath())

		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello",
			"name":  "小冯少爷",
		})
	})
	err := engine.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}

type ResponseEntity struct {
	Code    int         `json:"code" xml:"code"`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data" xml:"data"`
	Success bool        `json:"success" xml:"success"`
}
