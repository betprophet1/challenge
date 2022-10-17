package simplebet

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"project/common/failure"
	v1 "project/project/services/simplebet/v1"
)

func InitRouter(router *gin.Engine) {
	router.Use(func(ctx *gin.Context) {
		ctx.Next()
		for _, err := range ctx.Errors {
			if apperr, ok := err.Err.(failure.Failure); ok {
				ctx.JSON(apperr.Code(), gin.H{
					"error_code": apperr.ErrorCode(),
					"msg":        apperr.Err().Error(),
				})
				return
			}
			ctx.JSON(http.StatusExpectationFailed, err)
			return
		}
	})
	router.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "healthy")
	})
	rootGroup := router.Group("/api")
	v1.InitGroup(rootGroup.Group("/v1"))
}
