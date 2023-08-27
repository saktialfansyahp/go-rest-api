package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saktialfansyahp/go-rest-api/router"
)

var(
	app *gin.Engine
)

func myRoute(r *gin.RouterGroup) {
	router.DefineRoutes()
}

func init(){
	app := gin.New()
	r := app.Group("/go")
	myRoute(r)
}

func Handler(w http.ResponseWriter, r *http.Request){
	app.ServeHTTP(w,r)
}