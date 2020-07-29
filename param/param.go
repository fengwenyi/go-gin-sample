package param

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func QueryParam() {
	engine := gin.Default()
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		fmt.Println("请求路径：" + context.FullPath())

		//name := context.DefaultQuery("name", "佚名")

		name, _ := context.GetQuery("name")

		_, _ = context.Writer.Write([]byte("Hello " + name))
	})
	err := engine.Run() // 默认 8080端口
	// err := engine.Run("9090") // 定义端口
	if err != nil {
		log.Fatal(err.Error())
	}
}

//

//context.DefaultPostForm()
//context.GetQueryArray()
//context.GetQueryMap()

// http://localhost:8080/hello?name=%E5%BC%A0%E4%B8%89

// Hello 张三

func BindParam() {
	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("请求路径：" + context.FullPath())

		var loginUser LoginUser
		//err := context.BindQuery(&loginUser)
		//err := context.ShouldBind(&loginUser)
		err := context.ShouldBindQuery(&loginUser)
		if err != nil {
			log.Fatal(err.Error())
		}

		_, _ = context.Writer.Write([]byte(loginUser.Name + "-" + loginUser.Password))
	})
	err := engine.Run() // 默认 8080端口
	// err := engine.Run("9090") // 定义端口
	if err != nil {
		log.Fatal(err.Error())
	}
}

type LoginUser struct {
	Name     string `form:"name" json:"name" xml:"name"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}
