package routers

import (
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	routerV1(engine.Group("/v1"))
}
