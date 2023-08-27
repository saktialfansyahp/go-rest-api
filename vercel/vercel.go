package vercel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saktialfansyahp/go-rest-api/router"
)

var (
	app *gin.Engine
)

func init(){
	app = gin.New()
	r := app.Group("go")
	router.DefineRoutes(r)
}

func Handler(w http.ResponseWriter, r *http.Request){
	app.ServeHTTP(w, r)
}