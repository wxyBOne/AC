package main

import (
	"ac/controller"
	"ac/middlewares"
	"ac/utils"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()        // 返回一个*Engine实例，创建路由引擎
	router.Use(middlewares.Cors()) // 启用跨域中间件
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Println(".env 文件未找到，尝试使用系统环境变量")
	}
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("环境变量 DB_DSN 未设置")
	}
	err = utils.Init(dsn) // 初始化数据库连接
	if err != nil {
		panic(err) // 如果连接失败，抛出错误
	}
	registerRouter(router) // 注册路由
	router.Run(":8001")    // 启动服务器
}

var resetSum = 0
var resetUserInfo = false // 全局变量，指示是否需要重置用户信息

// 路由设置
func registerRouter(router *gin.Engine) {
	artCopyProController := new(controller.ArtCopyProController) // 创建艺术品控制器实例
	artCopyProController.ArtRouterController(router)             // 注册艺术品相关路由

	adminController := new(controller.AdminController)
	adminController.AdminRouterController(router)

	userController := new(controller.UserController) // 创建用户控制器实例
	userController.UserRouterController(router)      // 注册用户相关路由

	router.GET("/status", func(c *gin.Context) {
		//重启项目刷新用户信息
		if resetSum == 0 {
			resetUserInfo = true
			resetSum++
		} else {
			resetUserInfo = false
		}
		c.JSON(http.StatusOK, gin.H{"resetUserInfo": resetUserInfo}) // 返回当前状态
	})
}
