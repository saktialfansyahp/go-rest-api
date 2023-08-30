package router

import (
	"net/http"

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
	r.POST("role", func(ctx *gin.Context) {
		authcontroller.Role(ctx.Writer, ctx.Request)
	})
	api := r.Group("api")
	{
		admin := api.Group("admin")
		admin.Use(middleware.GINMiddleware("admin"))
		admin.GET("product", func(ctx *gin.Context) {
			productcontroller.Admin(ctx.Writer, ctx.Request)
		})
		user := api.Group("user")
		user.Use(middleware.GINMiddleware("admin", "user"))
		user.GET("productUser", func(ctx *gin.Context) {
			productcontroller.User(ctx.Writer, ctx.Request)
		})
	}

	r.Run()
}