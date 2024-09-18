package utils

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		starttime := time.Now()

		c.Next()

		duration := time.Since(starttime)
		log.Printf("%s %s %s %s", c.Request.Method, c.Request.URL.Path, c.Request.Proto, duration)

	}

}
