package chapter07

import (
	"github.com/gin-gonic/gin"
	"ginproject/logs_source"
)

func LogTest(ctx *gin.Context) {
	logs_source.Log.Info("this is info")
	logs_source.Log.Debug("this is debug")
	logs_source.Log.Warn("this is warn")

	logs_source.Log.WithField("name", "xx").Info("info")

	fileds_mapping := map[string]interface{}{
		"id": 18,
		"name": "haha",
	}
	logs_source.Log.WithFields(fileds_mapping).Info("info")
}
