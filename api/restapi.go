package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saktialfansyahp/go-rest-api/controllers/authcontroller"
	"github.com/saktialfansyahp/go-rest-api/controllers/productcontroller"
	"github.com/saktialfansyahp/go-rest-api/middleware"
)

var(
	app *gin.Engine
)

func myRoute(r *gin.RouterGroup) {
	r.GET("home", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Home")
	})
	r.POST("login", func(ctx *gin.Context) {
		authcontroller.Login(ctx.Writer, ctx.Request)
	})
	r.POST("register", func(ctx *gin.Context) {
		authcontroller.Register(ctx.Writer, ctx.Request)
	})
	r.GET("logout", func(ctx *gin.Context) {
		authcontroller.Logout(ctx.Writer, ctx.Request)
	})
	api := r.Group("api")
	{
		api.Use(middleware.GINMiddleware())
		api.GET("product", func(ctx *gin.Context) {
			productcontroller.Index(ctx.Writer, ctx.Request)
		})
	}
}

func init(){
	app := gin.New()
	r := app.Group("/go")
	myRoute(r)
}

func Handler(w http.ResponseWriter, r *http.Request){
	app.ServeHTTP(w,r)
}