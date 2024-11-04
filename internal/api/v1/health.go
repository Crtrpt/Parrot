package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Health(c *gin.Context) {
	logrus.Info("health")
	c.JSON(200, gin.H{
		"code":   "0",
		"mesage": "ok",
	})
}
