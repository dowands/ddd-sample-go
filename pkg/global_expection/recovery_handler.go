package global_expection

import "github.com/gin-gonic/gin"

func RecoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(500, "internal error")
}
