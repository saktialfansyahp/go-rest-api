package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/saktialfansyahp/go-rest-api/controllers/productcontroller"
	"github.com/saktialfansyahp/go-rest-api/models"
)

var (
	app *gin.Engine
)

func router(r *gin.RouterGroup) {
	models.ConnectDatabase()

	r.GET("api/product", productcontroller.Index)
}

func init(){
	app = gin.New()

	app.NoRoute(func(c *gin.Context) {
		sb := &strings.Builder{}
		sb.WriteString("routing err: no route, try this:\n")
		for _, v := range app.Routes() {
			sb.WriteString(fmt.Sprintf("%s %s\n", v.Method, v.Path))
		}
		c.String(http.StatusBadRequest, sb.String())
	})

	app.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
			c.JSON(http.StatusOK, gin.H{"message": "Preflight request successful"})
			c.Abort()
			return
		}

		if c.Request.Method == "POST" {
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
		} else if c.Request.Method == "GET" {
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		} else if c.Request.Method == "PUT" {
			c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT")
		} else if c.Request.Method == "DELETE" {
			c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE")
		}

		c.Next()
	})

	r := app.Group("/")

	router(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}