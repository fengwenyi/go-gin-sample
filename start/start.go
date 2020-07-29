package start

import (
	"github.com/gin-gonic/gin"
	"log"
)

func StartApp() {
	engine := gin.Default()
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		_, _ = context.Writer.Write([]byte("Hello"))
	})
	err := engine.Run() // 默认 8080端口
	// err := engine.Run("9090") // 定义端口
	if err != nil {
		log.Fatal(err.Error())
	}
}

// http://localhost:8080/hello
// Hello
