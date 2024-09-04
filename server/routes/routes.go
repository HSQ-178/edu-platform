package routes

import (
	"edu-profit/app/controllers"
	"edu-profit/app/middlewares/cors"
	"edu-profit/app/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./static")

	r.Use(cors.Cors())

	root := r.Group("/api")
	{
		user := root.Group("/user")
		{
			user.POST("/register", controllers.User.Register)        // 用户注册
			user.POST("/login", controllers.User.Login)              // 用户登录
			user.POST("/list", controllers.User.List)                // 用户列表
			user.POST("/update", jwt.JWT(), controllers.User.Update) // 用户更新
		}
	}

	return r
}
