package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/saktialfansyahp/go-rest-api/controllers/authcontroller"
	"github.com/saktialfansyahp/go-rest-api/controllers/cartcontroller"
	"github.com/saktialfansyahp/go-rest-api/controllers/categorycontroller"
	"github.com/saktialfansyahp/go-rest-api/controllers/colorcontroller"
	"github.com/saktialfansyahp/go-rest-api/controllers/productcontroller"
	"github.com/saktialfansyahp/go-rest-api/controllers/subcategorycontroller"
	"github.com/saktialfansyahp/go-rest-api/handler"
	"github.com/saktialfansyahp/go-rest-api/middleware"
	"github.com/saktialfansyahp/go-rest-api/models"
)

var (
	app *gin.Engine
)

func registerRouter(r *gin.RouterGroup) {

	models.ConnectDatabase()

	r.GET("/api/ping", handler.Ping)
	r.GET("home", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Home")
	})
	r.POST("/api/login", authcontroller.Login)
	r.POST("/api/register", authcontroller.Register)
	r.POST("/api/role", authcontroller.Role)
	r.POST("/api/cart", cartcontroller.Index)
	r.GET("logout", func(ctx *gin.Context) {
		authcontroller.Logout(ctx.Writer, ctx.Request)
	})
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
}

// init gin app
func init() {
	app = gin.New()

	// Handling routing errors
	app.NoRoute(func(c *gin.Context) {
		sb := &strings.Builder{}
		sb.WriteString("routing err: no route, try this:\n")
		for _, v := range app.Routes() {
			sb.WriteString(fmt.Sprintf("%s %s\n", v.Method, v.Path))
		}
		c.String(http.StatusBadRequest, sb.String())
	})

	r := app.Group("/")

	// register route
	registerRouter(r)
}

// entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
