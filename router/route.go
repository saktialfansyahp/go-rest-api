package router

import (
	"github.com/gin-gonic/gin"
	"github.com/saktialfansyahp/go-rest-api/controllers/authcontroller"
	"github.com/saktialfansyahp/go-rest-api/controllers/productcontroller"
	"github.com/saktialfansyahp/go-rest-api/middleware"
	"github.com/saktialfansyahp/go-rest-api/models"
)

// 	api := r.PathPrefix("api").Subrouter()
// 	api.HandleFunc("/product", productcontroller.Index).Methods("GET")
// 	api.Use(middleware.JWTMiddleware)

func DefineRoutes() {
	r := gin.Default()
	models.ConnectDatabase()

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
	// r.GET("api/products", productcontrollers.Index)

	// r.Run()
}