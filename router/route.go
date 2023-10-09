package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saktialfansyahp/go-rest-api/controllers/authcontroller"
	"github.com/saktialfansyahp/go-rest-api/controllers/categorycontroller"
	"github.com/saktialfansyahp/go-rest-api/controllers/colorcontroller"
	"github.com/saktialfansyahp/go-rest-api/controllers/productcontroller"
	"github.com/saktialfansyahp/go-rest-api/controllers/subcategorycontroller"
	"github.com/saktialfansyahp/go-rest-api/middleware"
	"github.com/saktialfansyahp/go-rest-api/models"
)

// 	api := r.PathPrefix("api").Subrouter()
// 	api.HandleFunc("/product", productcontroller.Index).Methods("GET")
// 	api.Use(middleware.JWTMiddleware)

func DefineRoutes() {
	r := gin.Default()
	models.ConnectDatabase()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS"{
			c.JSON(http.StatusOK, gin.H{"message": "Preflight request successful"})
			c.Abort()
			return
		}
		c.Next()
	})

	r.GET("home", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Home")
	})
	r.POST("/login", authcontroller.Login)
	r.POST("/register", authcontroller.Register)
	r.POST("/role", authcontroller.Role)
	r.GET("/logout", authcontroller.Logout)
	api := r.Group("api")
	api.GET("product", func(ctx *gin.Context) {
		productcontroller.Index(ctx.Writer, ctx.Request)
	})
	{
		admin := api.Group("admin")
		admin.Use(middleware.GINMiddleware("admin"))
		admin.GET("product/:id", func(ctx *gin.Context) {
			productcontroller.ById(ctx.Writer, ctx.Request, ctx.Param("id"))
		})
		admin.POST("product", func(ctx *gin.Context) {
			productcontroller.Create(ctx.Writer, ctx.Request)
		})
		admin.PUT("product/:id", func(ctx *gin.Context) {
			productcontroller.Edit(ctx.Writer, ctx.Request, ctx.Param("id"))
		})
		admin.DELETE("product/:id", func(ctx *gin.Context) {
			productcontroller.Delete(ctx.Writer, ctx.Request, ctx.Param("id"))
		})
		admin.GET("category", func(ctx *gin.Context) {
			categorycontroller.Index(ctx.Writer, ctx.Request)
		})
		admin.POST("category", func(ctx *gin.Context) {
			categorycontroller.Create(ctx.Writer, ctx.Request)
		})
		admin.GET("subcategory", func(ctx *gin.Context) {
			subcategorycontroller.Index(ctx.Writer, ctx.Request)
		})
		admin.POST("subcategory", func(ctx *gin.Context) {
			subcategorycontroller.Create(ctx.Writer, ctx.Request)
		})
		admin.GET("color", func(ctx *gin.Context) {
			colorcontroller.Index(ctx.Writer, ctx.Request)
		})
		admin.POST("color", func(ctx *gin.Context) {
			colorcontroller.Create(ctx.Writer, ctx.Request)
		})
		user := api.Group("user")
		user.Use(middleware.GINMiddleware("admin", "user"))
		user.GET("productUser", func(ctx *gin.Context) {
			productcontroller.User(ctx.Writer, ctx.Request)
		})
	}

	r.Run()
}