package app

import (
	"github.com/gin-gonic/gin"
	"mcenter-callback-demo/app/controller/test"
	"mcenter-callback-demo/app/loger"
	"mcenter-callback-demo/app/middleware"
	"mcenter-callback-demo/app/protocol"
	"net/http"
)

type App struct {
	engine *gin.Engine
}

func Default() *App {
	app := &App{}
	app.engine = gin.Default()
	return app
}

func (this *App) Run() {
	loger.Default()
	this.SetCors()
	this.setRouter()
	this.set404()
}

func (this *App) GetEngin() *gin.Engine {
	return this.engine
}
func (this *App) SetCors() {
	this.engine.Use(middleware.Cors())
}


func (this *App) setRouter() {
	t := test.Test{}
	authorized := this.engine.Group("/callback")
	authorized.Use(middleware.CheckAuth())
	{
		authorized.POST("/test", t.Test)
	}
}

func (this *App) set404() {
	this.engine.NoRoute(func(context *gin.Context) {
		resp := protocol.Resp{Code: 404, Msg: "page not exists!", Data: ""}
		//返回404状态码
		context.JSON(http.StatusNotFound, resp)
	})
}
