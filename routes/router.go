package routes

import (
	"gobdd/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	httpClient := &http.Client{}

	r.GET("/hello", handlers.HelloHandler())
	r.GET("/hello/:name", handlers.HelloNameHandler())
	r.GET("/hello/calculator", handlers.HelloCalculatorHandler())
	r.GET("hello/third-party", handlers.HelloThirdPartyApiHandler(httpClient))

	return r
}
