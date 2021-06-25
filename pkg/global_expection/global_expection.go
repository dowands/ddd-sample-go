package global_expection

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

func Recovery(f func(c *gin.Context, err interface{})) gin.HandlerFunc {
	return RecoveryWithWriter(f, gin.DefaultErrorWriter)
}

func RecoveryWithWriter(f func(c *gin.Context, err interface{}), out io.Writer) gin.HandlerFunc {
	var logger *log.Logger
	logger = log.New(out, "\n\n\x1b[31m", log.LstdFlags)

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Printf("%v\n", err)
				f(c, err)
			}
		}()
		c.Next()
	}
}