package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func Cors() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	//corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie", "X-Token", "*"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowOrigins = []string{
		"http://localhost:8080",
		"http://127.0.0.1:8080",
		"http://127.0.0.1:8080",
		"http://localhost:9528",
		"http://127.0.0.1:9528",
		"http://localhost:9529",
		"http://127.0.0.1:9529",
		"http://127.0.0.1:8080",
		"https://www.tinymce.com",
		"https://codepen.io",
		"http://127.0.0.1:8081",
	}
	//corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	return cors.New(corsConfig)
}


func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "http://127.0.0.1:8080",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}